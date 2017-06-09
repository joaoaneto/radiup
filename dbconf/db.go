package dbconf

import (
	"os"
	//"path/filepath"
	"io/ioutil"
	"encoding/json"
	"gopkg.in/mgo.v2"
)

type DbConfig struct {
	session     *mgo.Session
	connectInfo	*mgo.DialInfo
	collections []string
}

func NewDbConfig() *DbConfig {
	//data := GetConnectionData()
	dbcfg := &DbConfig{collections: []string{"CYCLE", "STREAMER", "PLAYLIST"}, connectInfo: GetConnectionData()}
	var err error
	dbcfg.session, err = mgo.DialWithInfo(dbcfg.connectInfo)
	
	if err != nil {
		panic(err)
	}
	dbcfg.session.SetMode(mgo.Monotonic, true)

	return dbcfg
}

//Enum interface used for abstract the ConnectionSetup inputs
type Enum interface {
	GetCollection() *mgo.Collection
	//GetConnection() ConnectionData
}

type ConnectionSetup uint

//This const block is equivalent to traditional Enum (Java, C, ...)
const (
	CYCLE ConnectionSetup = iota //it works as an autoincrement
	STREAMER
	PLAYLIST
)

//When it magic happens
//return mgo.Collection according to subsystem types
func (db *DbConfig) GetCollection(cs ConnectionSetup) *mgo.Collection {
	c := db.session.DB(db.connectInfo.Database).C(db.collections[cs])
	return c
}

//future func for get data of connect setup file
func GetConnectionData() *mgo.DialInfo {
	gopath := os.Getenv("GOPATH")
	src := gopath + "/src/github.com/joaoaneto/radiup/dbconf/db.config"
	var mongoDialInfo *mgo.DialInfo
	data, err := ioutil.ReadFile(src)
	if err != nil {
		panic(err)
	}

	err2 := json.Unmarshal(data, &mongoDialInfo)
	if err2 != nil {
		panic(err)
	}

	return mongoDialInfo

}

//close the session above declared and initialized
func (db *DbConfig) CloseSession() {
	db.session.Close()
}