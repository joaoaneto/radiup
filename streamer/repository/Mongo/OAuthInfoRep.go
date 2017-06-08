package Mongo

import (
	"log"

	"github.com/joaoaneto/radiup/dbconf"
	"github.com/joaoaneto/radiup/streamer"
	streamerRepository "github.com/joaoaneto/radiup/streamer/repository"
	"gopkg.in/mgo.v2/bson"
)

type OAuthInfoPersistor struct {
	db *dbconf.DbConfig
}

func NewPersistorOAuthInfo() streamerRepository.OAuthInfoManager {
	return &OAuthInfoPersistor{dbconf.NewDbConfig()}
}

func (p OAuthInfoPersistor) Register(oAuth streamer.OAuthInfo) error {

	r := p.db.GetCollection(dbconf.STREAMER)
	err := r.Insert(&oAuth)

	if err != nil {
		log.Fatal(err)
	}

	return err

}

func (p OAuthInfoPersistor) Search(clientID string) (streamer.OAuthInfo, error) {

	r := p.db.GetCollection(dbconf.STREAMER)
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
