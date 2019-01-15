package events

import (
	"log"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

var (
	Session neo4j.Session
)

func SetDB(s neo4j.Session) {
	Session = s
}

func handleErr(err error) bool {
	if err != nil {
		log.Println(err)
		return true
	}
	return false
}
