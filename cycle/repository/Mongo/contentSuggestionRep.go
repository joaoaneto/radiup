package repository

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"github.com/joaoaneto/radiup/cycle"
)

type ContentSuggestionRep struct {
	title       string
	description string
	user        cycle.User
	votes       int
	validated   bool
	done        bool
}

type persistor struct {
	Session *mgo.Session
}

func NewPersistor(pSession *mgo.Session) ContentSuggestionManager {
	return persistor{pSession}
}

func (p persistor) RegisterCSuggestion(cs cycle.ContentSuggestion) {
	
	session = p.Session

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	defer session.Close()

	c := session.DB("radiup").C("cycle")

	err = c.Insert(&cs)

	if err != nil {
		log.Fatal(err)
	}

}

func (p persistor) SearchCSuggestion(nameUser interface{}) []ContentSuggestionRep {

	session = p.Session

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	defer session.Close()

	c := session.DB("radiup").C("cycle")

	result := []ContentSuggestionRep{}

	err = c.Find(bson.M{"name": nameUser}).One(&result)

	if err != nil {
		log.Fatal(err)
	}

	return result

}

/*func (cs contentSuggestion) RemoveCSuggestion(nameUser string) {

	session, err := mgo.Dial("localhost")

	if err != nil {
		panic(err)
	}

	defer session.Close()

	c := session.DB("teste").C("Cycle")

	result := contentSuggestion{}

	err = c.Remove(bson.M{"name" : nameUser})

	if err != nil {
      log.Fatal(err)
    }

}*/
