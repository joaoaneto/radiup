package repository 

import (
	"gopkg.in/mgo.v2"
)

// get_session return the connection session with MongoDB
// according configuration parameters (host, authentication, others)
func get_session() *mgo.Session {

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	return session

}