package mongo

import(
	"fmt"
	"time"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/joaoaneto/radiup"
	"github.com/joaoaneto/radiup/cycle"
)

/*Implementation of Cycle's repository interfaces*/

/*ContentSuggestion Mongo implementations*/
type ContentSuggestionPersistor struct {
}

func NewPersistorContentSuggestion() ContentSuggestionPersistor {
	return ContentSuggestionPersistor{}
}

func (p ContentSuggestionPersistor) Register(cs cycle.ContentSuggestion) {
	
	//defer session.Close()

	c := CYCLE.GetCollection()

	err = c.Insert(&cs)

	if err != nil {
		log.Fatal(err)
	}

}

func (p ContentSuggestionPersistor) Search(nameUser interface{}) []ContentSuggestion{

	//defer session.Close()

	c := CYCLE.GetCollection()

	result := []ContentSuggestion{}

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

/*Cycle Mongo implementations*/

type CyclePersistor struct {
}

func NewPersistorCycle() CyclePersistor {
	return CyclePersistor{}
}

func (cp CyclePersistor) Create(c cycle.Cycle) {
	
    //defer session.Close()

    c := CYCLE.GetCollection()
   
    err = c.Insert(&c)

    if err != nil {
    	log.Fatal(err)
    }

}

func (cp CyclePersistor) Update(registered_id int, start time.Time,
														   end time.Time,
														   _type string,
														   description string,
														   voluntarySuggestion cycle.VoluntarySuggestion,
														   streamerSuggestion cycle.StreamerSuggestion,
														   contentSuggestion cycle.ContentSuggestion) {
	
    //defer session.Close()

    c := CYCLE.GetCollection()
    
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

func (cp CyclePersistor) Remove(id int) {
	
	//defer session.Close()
   
    c := CYCLE.GetCollection()

    err = c.Remove(bson.M{"id" : id})

    if err != nil {
    	log.Fatal(err)
    }

}

func (cp CyclePersistor) Search(id int) cycle.Cycle {

	result := cycle.Cycle{}
	
	//defer session.Close()

    c := CYCLE.GetCollection()

    err = c.Find(bson.M{"id" : id}).One(&result)

	if err != nil {
    	log.Fatal(err)
    }

    return result
}


/*Music Mongo implementations*/

type MusicPersistor struct{
}

func NewPersistorMusic() MusicPersistor {
	return MusicPersistor{}
}

func (mp MusicPersistor) Register(m cycle.Music){
	
	//defer session.Close()

	c := CYCLE.GetCollection()

	/*Insert the music object on DataBase*/
	err = c.Insert(&m)

	if err != nil{
		log.Fatal(err)
	}

}

func (mp MusicPersistor) Remove(id string){

	//defer session.Close()

	c := CYCLE.GetCollection()

	/*Insert the music object on DataBase*/
	err = c.Remove(bson.M{"id" : id})

	if err != nil{
		log.Fatal(err)
	}	

}

func (mp MusicPersistor) Search(id string){

	result := cycle.Music{}

	//defer session.Close()

	c := CYCLE.GetCollection()

	/*Insert the music object on DataBase*/
	err = c.Find(bson.M{"id" : id }).One(&result)

	if err != nil{
		log.Fatal(err)
	}

	return result

}

/*VoluntarySuggestion Mongo implementations*/

type VoluntarySuggestionPersistor struct {
}

func NewPersistorVoluntarySuggestion() VoluntarySuggestionPersistor {
	return VoluntarySuggestionPersistor{}
}

func (p VoluntarySuggestionPersistor) Register(v cycle.VoluntarySuggestion) {

	//defer session.Close()

	c := CYCLE.GetCollection()

	err = c.Insert(&v)

	if err != nil {
		log.Fatal(err)
	}

}

func (p VoluntarySuggestionPersistor) Search(nameUser string) []VoluntarySuggestion {

	//defer session.Close()

	c := CYCLE.GetCollection()

	result := []VoluntarySuggestion{}

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

/*User mongo implementations*/

type UserPersistor struct {
}

func NewPersistorUser() UserPersistor {
	return UserPersistor{}
}

func (up UserPersistor) Create(u cycle.User) {
	
	//defer session.Close()
    
    c := CYCLE.GetCollection()
   
    err = c.Insert(&u)

    if err != nil {
    	log.Fatal(err)
    }

}

func (up UserPersistor) Update(registered_user string,
									  name string,
									  password string,
									  birth_day time.Time,
									  email string,
									  sex byte) {
	
    //defer session.Close()
 
    c := CYCLE.GetCollection()
    
    wantedUser := bson.M{"username" : registered_user}

    changes := bson.M{"$set" : bson.M{"name" : name, 
    								  "password" : password,
							  		  "birth_day" : birth_day,
							   		  "email" : email,
							   		  "sex" : sex}}

	err = c.Update(wantedUser, changes)

	if err != nil {
    	log.Fatal(err)
    }

}

func (up UserPersistor) Remove(username string) {
	
	//defer session.Close()

    c := CYCLE.GetCollection()

    err = c.Remove(bson.M{"username" : username})

    if err != nil {
    	log.Fatal(err)
    }

}

func (up UserPersistor) Search(username string) cycle.User {

	result := cycle.User{}

    //defer session.Close()

    c := CYCLE.GetCollection()

    err = c.Find(bson.M{"username" : username}).One(&result)

	if err != nil {
    	log.Fatal(err)
    }

    return result
}