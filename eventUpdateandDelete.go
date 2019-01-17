package events

import (
	"log"

	events "github.com/angadsharma1016/omega_dbconfig"
)

// delete event with given query
func DeleteEvent(q Query, c chan error) {
	result, err := events.Session.Run(`
		MATCH(n:EVENT)-[r]->(a)
		WHERE n.`+q.Key+`=$val
		DETACH DELETE n, a
	`, map[string]interface{}{
		"val": q.Value,
	})
	if err != nil {
		c <- err
	}

	result.Next()
	log.Println(result.Record())

	if err = result.Err(); err != nil {
		c <- err
		return
	}
	log.Println("Event deleted")
	c <- nil
	return
}

// update event with given query and new value
func UpdateEvent(q Query, c chan error) {
	result, err := events.Session.Run(`
		MATCH(n:EVENT)
		WHERE n.`+q.Key+`=$val
		SET n.`+q.ChangeKey+`=$val1
		RETURN n.`+q.ChangeKey+`
	`, map[string]interface{}{
		"val":  q.Value,
		"val1": q.ChangeValue,
	})

	if err != nil {
		c <- err
		return
	}
	c <- nil

	result.Next()
	log.Printf("Updated %s to %s", q.Key, result.Record().GetByIndex(0).(string))

	if err = result.Err(); err != nil {
		c <- err
		return
	}
}
