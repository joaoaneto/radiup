package repository

import (
	"github.com/joaoaneto/radiup/streamer"
	"gopkg.in/mgo.v2/bson"
	"log"
	"fmt"
)

type OAuthInfoRep struct {
	clientID string //anybody gotta update in OAuthInfo entitie
	SecretKey string 
}

type persistor struct {
}

func NewPersistor() OAuthInfoManager {
	return persistor{}
}

func (p persistor) Register(oAuth streamer.OAuthInfo){

	session := get_session()

	//oAuthDef := OAuthInfoRep(oAuth)

	r := session.DB("radiup").C("oAuthInfo")
	err := r.Insert(&oAuth)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inseriu com sucesso!")

}

func (p persistor) Search(client_id string) OAuthInfoRep {

	session := get_session()

	r := session.DB("radiup").C("oAuthInfo")
	result := OAuthInfoRep{}
	err := r.Find(bson.M{"clientid":client_id}).One(&result)
	if err != nil {
		panic(err)
	}

	return result
}

func (p persistor) Update(client_id string, secret_key string) {

	session := get_session()

	r := session.DB("radiup").C("oAuthInfo")

	selectOld := bson.M{"clientid":client_id}
	change := bson.M{"$set":bson.M{"secretKey":secret_key}}
	err := r.Update(selectOld, change)
	
	if err != nil{
		panic(err)
	}

}

func (p persistor) Remove(client_id string) {

	session:= get_session()

	r := session.DB("radiup").C("oAuthInfo")

	err := r.Remove(bson.M{"clientid": client_id})
	if err != nil {
		panic(err)
	}
}	