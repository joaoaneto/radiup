package repository

import (
	"time"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
    "github.com/joaoaneto/radiup/cycle"
)

type CycleMGO struct {
	id int	
	start time.Time
	end time.Time
	_type string
	description string
	voluntarySuggestion VoluntarySuggestion
	streamerSuggestion streamerSuggestion
	contentSuggestion ContentSuggestion
}

type CyclePersistence struct {

}

func (cp *CyclePersistence) CreateCycle(c cycle.Cycle) {
	
	//Connection
	session, err := mgo.Dial("localhost")
    if err != nil {
    	panic(err)
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)

    c := session.DB("radiup").C("cycle")
   
    err = c.Insert(&c)

    if err != nil {
    	log.Fatal(err)
    }

}

func (cp *CyclePersistence) UpdateCycle(registered_id int, start time.Time,
														   end time.Time,
														   _type string,
														   description string,
														   voluntarySuggestion cycle.VoluntarySuggestion,
														   streamerSuggestion cycle.streamerSuggestion,
														   contentSuggestion cycle.ContentSuggestion) {
	//Connection
	session, err := mgo.Dial("localhost")
    if err != nil {
    	panic(err)
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)

    c := session.DB("radiup").C("cycle")
    
    wantedCycle := bson.M{"id" : registered_id}

    changes := bson.M{"$set" : bson.M{"start" : start,
									  "end" : end,
									  "_type" : _type,
									  "description" : description,
									  "voluntarySuggestion" : voluntarySuggestion,
									  "streamerSuggestion" : streamerSuggestion
								  	  "contentSuggestion" : contentSuggestion}}

	err = c.Update(wantedCycle, changes)

	if err != nil {
    	log.Fatal(err)
    }

}

func (cp *CyclePersistence) RemoveCycle(id int) {
	
	//Connection
	session, err := mgo.Dial("localhost")
    if err != nil {
    	panic(err)
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)

    c := session.DB("radiup").C("cycle")

    err = c.Remove(bson.M{"id" : id})

    if err != nil {
    	log.Fatal(err)
    }

}

func (cp *CyclePersistence) SearchCycle(id int) cycle.Cycle {

	result := cycle.Cycle{}

	//Connection
	session, err := mgo.Dial("localhost")
    if err != nil {
    	panic(err)
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)

    c := session.DB("radiup").C("cycle")

    err = c.Find(bson.M{"id" : id}).One(&result)

	if err != nil {
    	log.Fatal(err)
    }

    return result
}