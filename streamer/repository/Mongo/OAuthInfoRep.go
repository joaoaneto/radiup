package repository

import (
	"log"

	"github.com/joaoaneto/radiup/streamer"
	"gopkg.in/mgo.v2/bson"
)

type OAuthInfoPersistor struct {
}

func NewPersistorOAuthInfo() OAuthInfoManager {
	return OAuthInfoPersistor{}
}

func (p OAuthInfoPersistor) Register(oAuth streamer.OAuthInfo) string {

	r := STREAMER.GetCollection()
	err := r.Insert(&oAuth)

	if err != nil {
		log.Fatal(err)
	}

	return err

}

func (p OAuthInfoPersistor) Search(clientID string) (streamer.OAuthInfo, string) {

	r := STREAMER.GetCollection()
	result := streamer.OAuthInfo{}
	err := r.Find(bson.M{"clientid": clientID}).One(&result)
	if err != nil {
		panic(err)
	}

	return result, err
}

func (p OAuthInfoPersistor) Update(clientID string, secretKey string) string {

	r := STREAMER.GetCollection()

	selectOld := bson.M{"clientid": clientID}
	change := bson.M{"$set": bson.M{"secretKey": secretKey}}
	err := r.Update(selectOld, change)

	if err != nil {
		panic(err)
	}

	return err

}

func (p OAuthInfoPersistor) Remove(clientID string) string {

	session := get_session()

	r := STREAMER.GetCollection()

	err := r.Remove(bson.M{"clientid": clientID})
	if err != nil {
		panic(err)
	}

	return err
}
