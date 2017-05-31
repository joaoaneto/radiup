package repositoty

import(
	"fmt"
	"time"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/joaoaneto/radiup/cycle"
)

type UserMGO struct {
	name string
	username string
	password string //temp... we should search for a real alternative in web scenario
	birth_day time.Time
	email string
	sex byte
}

type UserPersistence struct {
}

func (up *UserPersistence) CreateUser(u cycle.User) {
	
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

func (up *UserPersistence) UpdateUser(registered_user string,
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

func (up *UserPersistence) RemoveUser(username string) {
	
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

func (up *UserPersistence) SearchUser(username string) cycle.User {

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