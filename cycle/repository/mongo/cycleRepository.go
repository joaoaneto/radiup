package mongo

import(
	"fmt"
	"time"
	"log"
	//"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/joaoaneto/radiup/dbconf"
	"github.com/joaoaneto/radiup/cycle"
	cycleRep "github.com/joaoaneto/radiup/cycle/repository"
)

/*Implementation of Cycle's repository interfaces*/

/*ContentSuggestion Mongo implementations*/
type ContentSuggestionPersistor struct {
	db *dbconf.DbConfig
}

func NewPersistorContentSuggestion() cycleRep.ContentSuggestionManager {
	return &ContentSuggestionPersistor{dbconf.NewDbConfig()}
}

func (p ContentSuggestionPersistor) Register(cs cycle.ContentSuggestion) error {
	
	//defer session.Close()

	c := p.db.GetCollection(dbconf.CYCLE)

	err := c.Insert(&cs)

	if err != nil {
		log.Fatal(err)
	}

	return err

}

func (p ContentSuggestionPersistor) Search(nameUser interface{}) ([]cycle.ContentSuggestion, error) {

	//defer session.Close()

	c := p.db.GetCollection(dbconf.CYCLE)

	result := []cycle.ContentSuggestion{}

	err := c.Find(bson.M{"name": nameUser}).One(&result)

	if err != nil {
		log.Fatal(err)
	}

	return result, err

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
	db *dbconf.DbConfig
}

func NewPersistorCycle() cycleRep.CycleManager {
	return &CyclePersistor{dbconf.NewDbConfig()}
}

func (cp CyclePersistor) Create(c cycle.Cycle) error {
	
    //defer session.Close()

    coll := cp.db.GetCollection(dbconf.CYCLE)
   
    err := coll.Insert(&c)

    if err != nil {
    	log.Fatal(err)
    }

	return err

}

func (cp CyclePersistor) Update(registeredID int, start time.Time,
														   end time.Time,
														   cycleType string,
														   description string,
														   voluntarySuggestion cycle.VoluntarySuggestion,
														   streamerSuggestion cycle.StreamerSuggestion,
														   contentSuggestion cycle.ContentSuggestion) error {
	
    //defer session.Close()

    c := cp.db.GetCollection(dbconf.CYCLE)
    
    wantedCycle := bson.M{"id" : registeredID}

    changes := bson.M{"$set" : bson.M{"start" : start,
									  "end" : end,
									  "cycleType" : cycleType,
									  "description" : description,
									  "voluntarySuggestion" : voluntarySuggestion,
									  "streamerSuggestion" : streamerSuggestion,
								  	  "contentSuggestion" : contentSuggestion}}

	err := c.Update(wantedCycle, changes)

	if err != nil {
    	log.Fatal(err)
    }

	return err

}

func (cp CyclePersistor) Remove(id int) error {
	
	//defer session.Close()
   
    c := cp.db.GetCollection(dbconf.CYCLE)

    err := c.Remove(bson.M{"id" : id})

    if err != nil {
    	log.Fatal(err)
    }

	return err

}

func (cp CyclePersistor) Search(id int) (cycle.Cycle, error) {

	result := cycle.Cycle{}
	
	//defer session.Close()

    c := cp.db.GetCollection(dbconf.CYCLE)

    err := c.Find(bson.M{"id" : id}).One(&result)

	if err != nil {
    	log.Fatal(err)
    }

    return result, err
}


/*Music Mongo implementations*/

type MusicPersistor struct{
	db *dbconf DbConfig
}

func NewPersistorMusic() cycleRep.MusicManager {
	return &MusicPersistor{dbconf.NewDbConfig()}
}

func (mp MusicPersistor) Register(m cycle.Music) error {
	
	//defer session.Close()

	c := mp.db.GetCollection(dbconf.CYCLE)

	/*Insert the music object on DataBase*/
	err := c.Insert(&m)

	if err != nil{
		log.Fatal(err)
	}

	return err

}

func (mp MusicPersistor) Remove(id string) error {

	//defer session.Close()

	c := mp.db.GetCollection(dbconf.CYCLE)

	/*Insert the music object on DataBase*/
	err := c.Remove(bson.M{"id" : id})

	if err != nil{
		log.Fatal(err)
	}	

	return err

}

func (mp MusicPersistor) Search(id string) (cycle.Music, error) {

	result := cycle.Music{}

	//defer session.Close()

	c := mp.db.GetCollection(dbconf.CYCLE)

	/*Insert the music object on DataBase*/
	err := c.Find(bson.M{"id" : id }).One(&result)

	if err != nil{
		log.Fatal(err)
	}

	return result, err

}

/*VoluntarySuggestion Mongo implementations*/

type VoluntarySuggestionPersistor struct {
	db *dbconf.DbConfig
}

func NewPersistorVoluntarySuggestion() cycleRep.VoluntarySuggestionManager {
	return &VoluntarySuggestionPersistor{dbconf.NewDbConfig()}
}

func (p VoluntarySuggestionPersistor) Register(v cycle.VoluntarySuggestion) error {

	//defer session.Close()

	c := p.db.GetCollection(dbconf.CYCLE)

	err := c.Insert(&v)

	if err != nil {
		log.Fatal(err)
	}

	return err

}

func (p VoluntarySuggestionPersistor) Search(nameUser string) ([]cycle.VoluntarySuggestion, error) {

	//defer session.Close()

	c := p.db.GetCollection(dbconf.CYCLE)

	result := []cycle.VoluntarySuggestion{}

	err := c.Find(bson.M{"name": nameUser}).One(&result)

	if err != nil {
		log.Fatal(err)
	}

	return result, err

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
	db *dbconf.DbConfig
}

func NewPersistorUser() cycleRep.UserManager {
	return &UserPersistor{dbconf.NewDbConfig()}
}

func (up UserPersistor) Create(u cycle.User) error {
	
	//defer session.Close()
    
    fmt.Print("OLa")
    c := up.db.GetCollection(dbconf.CYCLE)
   	fmt.Print("Oi")

    err := c.Insert(&u)
    fmt.Print("Ei")
    if err != nil {
    	log.Fatal(err)
    }

	return err

}

func (up UserPersistor) Update(registered_user string,
									  name string,
									  password string,
									  birth_day time.Time,
									  email string,
									  sex byte) error {
	
    //defer session.Close()
 
    c := up.db.GetCollection(dbconf.CYCLE)
    
    wantedUser := bson.M{"username" : registered_user}

    changes := bson.M{"$set" : bson.M{"name" : name, 
    								  "password" : password,
							  		  "birth_day" : birth_day,
							   		  "email" : email,
							   		  "sex" : sex}}

	err := c.Update(wantedUser, changes)

	if err != nil {
    	log.Fatal(err)
    }

	return err

}

func (up UserPersistor) Remove(username string) error {
	
	//defer session.Close()

    c := up.db.GetCollection(dbconf.CYCLE)

    err := c.Remove(bson.M{"username" : username})

    if err != nil {
    	log.Fatal(err)
    }

	return err

}

func (up UserPersistor) Search(username string) (cycle.User, error) {

	result := cycle.User{}

    //defer session.Close()

    c := up.db.GetCollection(dbconf.CYCLE)

    err := c.Find(bson.M{"username" : username}).One(&result)

	if err != nil {
    	log.Fatal(err)
    }

    return result, err
}