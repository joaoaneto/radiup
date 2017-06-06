package repository

import (
	"github.com/joaoaneto/radiup"
	"github.com/joaoaneto/radiup/streamer"
	"gopkg.in/mgo.v2/bson"
	"log"
	"fmt"
)

type OAuthInfoPersistor struct {
}

func NewPersistorOAuthInfo() OAuthInfoManager {
	return OAuthInfoPersistor{}
}

func (p OAuthInfoPersistor) Register(oAuth streamer.OAuthInfo){

	r := STREAMER.GetCollection()
	err := r.Insert(&oAuth)

	if err != nil {
		log.Fatal(err)
	}

}

func (p OAuthInfoPersistor) Search(clientId string) streamer.OAuthInfo {

	r := STREAMER.GetCollection()
	result := streamer.OAuthInfo{}
	err := r.Find(bson.M{"clientid":clientId}).One(&result)
	if err != nil {
		panic(err)
	}

	return result
}

func (p OAuthInfoPersistor) Update(clientId string, secretKey string) {

	r := STREAMER.GetCollection()

	selectOld := bson.M{"clientid":clientId}
	change := bson.M{"$set":bson.M{"secretKey":secretKey}}
	err := r.Update(selectOld, change)
	
	if err != nil{
		panic(err)
	}

}

func (p OAuthInfoPersistor) Remove(clientId string) {

	session:= get_session()

	r := STREAMER.GetCollection()

	err := r.Remove(bson.M{"clientid": clientId})
	if err != nil {
		panic(err)
	}
}	