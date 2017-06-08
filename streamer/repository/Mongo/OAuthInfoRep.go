package repository

import (
	"log"

	db "github.com/joaoaneto/radiup/dbconf"
	"github.com/joaoaneto/radiup/streamer"
	"gopkg.in/mgo.v2/bson"
)

type OAuthInfoPersistor struct {
}

func NewPersistorOAuthInfo() OAuthInfoPersistor {
	return OAuthInfoPersistor{}
}

func (p OAuthInfoPersistor) Register(oAuth streamer.OAuthInfo) error {

	r := db.STREAMER.GetCollection()
	err := r.Insert(&oAuth)

	if err != nil {
		log.Fatal(err)
	}

	return err

}

func (p OAuthInfoPersistor) Search(clientID string) (streamer.OAuthInfo, error) {

	r := db.STREAMER.GetCollection()
	result := streamer.OAuthInfo{}
	err := r.Find(bson.M{"clientid": clientID}).One(&result)
	if err != nil {
		panic(err)
	}

	return result, err
}

func (p OAuthInfoPersistor) Update(clientID string, secretKey string) error {

	r := db.STREAMER.GetCollection()

	selectOld := bson.M{"clientid": clientID}
	change := bson.M{"$set": bson.M{"secretKey": secretKey}}
	err := r.Update(selectOld, change)

	if err != nil {
		panic(err)
	}

	return err

}

func (p OAuthInfoPersistor) Remove(clientID string) error {

	//session := get_session()

	r := db.STREAMER.GetCollection()

	err := r.Remove(bson.M{"clientid": clientID})
	if err != nil {
		panic(err)
	}

	return err
}
