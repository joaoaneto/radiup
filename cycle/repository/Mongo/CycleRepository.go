package repository

import(
	"fmt"
	"time"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/joaoaneto/radiup/cycle"
)
/*Implementation of Cycle's repository interfaces*/

/*ContentSuggestion Mongo implementations*/
type ContentSuggestionPersistor struct {
	Session *mgo.Session
}

func NewPersistorContentSuggestion(pSession *mgo.Session) ContentSuggestionPersistor {
	return ContentSuggestionPersistor{pSession}
}

func (p ContentSuggestionPersistor) Register(cs cycle.ContentSuggestion) {
	
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

func (p ContentSuggestionPersistor) Search(nameUser interface{}) []ContentSuggestion{

	session = p.Session

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	defer session.Close()

	c := session.DB("radiup").C("cycle")

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
	Session mgo*session
}

func NewPersistorCyclePersistor(pSession *mgo.Session) CyclePersistor {
	return CyclePersistor{pSession}
}

func (cp *CyclePersistor) Create(c cycle.Cycle) {
	
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

func (cp *CyclePersistor) Update(registered_id int, start time.Time,
														   end time.Time,
														   _type string,
														   description string,
														   voluntarySuggestion cycle.VoluntarySuggestion,
														   streamerSuggestion cycle.StreamerSuggestion,
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

func (cp *CyclePersistor) Remove(id int) {
	
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

func (cp *CyclePersistor) Search(id int) cycle.Cycle {

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


/*Music Mongo implementations*/

type MusicPersistor struct{
	Session *mgo.Session
}

func NewPersistorMusicPersistor(pSession *mgo.Session) MusicPersistor {
	return MusicPersistor{pSession}
}

func (mp *MusicPersistor) Register(m cycle.Music){
	/*This will come from the "arquivo de conexão"*/

	session, err := mgo.Dial("localhost")

	if err != nil {
		panic(err)
	}	
	
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("radiup").C("cycle")

	/*Insert the music object on DataBase*/
	err = c.Insert(&m)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("Objeto música inserido com sucesso")
}

func (mp *MusicPersistor) Remove(id string){
	
	session, err := mgo.Dial("localhost")

	if err != nil {
		panic(err)
	}	

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("radiup").C("cycle")

	/*Insert the music object on DataBase*/
	err = c.Remove(bson.M{"id" : id})

	if err != nil{
		log.Fatal(err)
	}	

	fmt.Println("Objeto removido")

}

func (mp *MusicPersistor) Search(id string){

	result := cycle.Music{}

	session, err := mgo.Dial("localhost")

	if err != nil {
		panic(err)
	}	

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("radiup").C("cycle")

	/*Insert the music object on DataBase*/
	err = c.Find(bson.M{"id" : id }).One(&result)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(result)
}

/*VoluntarySuggestion Mongo implementations*/

type VoluntarySuggestionPersistor struct {
	Session *mgo.Session
}

func NewPersistor(pSession *mgo.Session) VoluntarySuggestionPersistor {
	return VoluntarySuggestionPersistor{pSession}
}

func (p VoluntarySuggestionPersistor) Register(v cycle.VoluntarySuggestion) {
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

func (p VoluntarySuggestionPersistor) Search(nameUser string) []VoluntarySuggestionRep {

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

/*User mongo implementations*/

type UserPersistor struct {
	Session *mgo.Session
}

func NewPersistorUserPersistor(pSession *mgo.Session) UserPersistor {
	return UserPersistor{pSession}
}

func (up *UserPersistor) Create(u cycle.User) {
	
	//Connection
	session, err := mgo.Dial("localhost")
    if err != nil {
    	panic(err)
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)

    c := session.DB("radiup").C("cycle")
   
    err = c.Insert(&u)

    if err != nil {
    	log.Fatal(err)
    }

    fmt.Printf("Usuário criado com sucesso!")
}

func (up *UserPersistor) Update(registered_user string,
									  name string,
									  password string,
									  birth_day time.Time,
									  email string,
									  sex byte) {
	//Connection
	session, err := mgo.Dial("localhost")
    if err != nil {
    	panic(err)
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)

    c := session.DB("radiup").C("cycle")
    
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

    fmt.Println("Usuário atualizado com sucesso!")
}

func (up *UserPersistor) Remove(username string) {
	
	//Connection
	session, err := mgo.Dial("localhost")
    if err != nil {
    	panic(err)
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)

    c := session.DB("radiup").C("cycle")

    err = c.Remove(bson.M{"username" : username})

    if err != nil {
    	log.Fatal(err)
    }

    fmt.Println("Usuário removido com sucesso!")
}

func (up *UserPersistor) Search(username string) cycle.User {

	result := cycle.User{}

	//Connection
	session, err := mgo.Dial("localhost")
    if err != nil {
    	panic(err)
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)

    c := session.DB("radiup").C("cycle")

    err = c.Find(bson.M{"username" : username}).One(&result)

	if err != nil {
    	log.Fatal(err)
    }

    return result
}