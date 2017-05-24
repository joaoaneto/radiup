package repository

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"teste_go/cycle"
	"time"
)

type VoluntarySuggestionRep struct {
	//track Music
	user      cycle.User
	votes     int
	Timestamp time.Time
}

type persistor struct {
}

func NewPersistor() IVoluntarySuggestion {
	return persistor{}
}

func (p persistor) RegisterVSuggestion(v cycle.VoluntarySuggestion) {
	session, err := mgo.Dial("localhost")

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	defer session.Close()

	c := session.DB("radiup").C("cycle")

	err = c.Insert(&v)

	if err != nil {
		log.Fatal(err)
	}

}

func (p persistor) SearchVSuggestion(nameUser string) []VoluntarySuggestionRep {

	session, err := mgo.Dial("localhost")

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	defer session.Close()

	c := session.DB("radiup").C("cycle")

	result := []VoluntarySuggestionRep{}

	err = c.Find(bson.M{"name": nameUser}).One(&result)

	if err != nil {
		log.Fatal(err)
	}

	return result

}

/*func (v VoluntarySuggestion) RemoveVSuggestion(nameUser string) {

	session, err := mgo.Dial("localhost")

	if err != nil {
		panic(err)
	}

	defer session.Close()

	c := session.DB("teste").C("Cycle")

	result := voluntarySuggestion{}

	err = c.Remove(bson.M{"name" : nameUser})

	if err != nil {
      log.Fatal(err)
    }

}*/
