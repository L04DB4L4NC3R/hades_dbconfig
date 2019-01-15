package events

import "github.com/neo4j/neo4j-go-driver/neo4j"

func ConnectToDB() (neo4j.Session, neo4j.Driver, error) {

	// define driver, session and result vars
	var (
		driver  neo4j.Driver
		session neo4j.Session
		err     error
	)

	// initialize driver to connect to localhost with ID and password
	if driver, err = neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("angad", "angad", "")); err != nil {
		return nil, nil, err
	}

	// Open a new session with write access
	if session, err = driver.Session(neo4j.AccessModeWrite); err != nil {
		return nil, nil, err
	}
	return session, driver, nil
}
