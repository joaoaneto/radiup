package repository

import(
	"fmt"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/joaoaneto/radiup/cycle"
)

type MusicMGO struct {
	name string
	artist []string
	id string
	source_id int //or string? - come from DataSource
}

type musicPersistence struct{
}

func (mp *musicPersistence) RegisterMusic(m cycle.Music){
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

func (mp *musicPersistence) RemoveMusic(id string){
	
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

func (mp *musicPersistence) SearchMusic(id string){

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

