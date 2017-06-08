package dbconf

import (
	"fmt"
	//	"io/ioutil"
	//	"encoding/json"
	"gopkg.in/mgo.v2"
)

type DbConfig struct {
	session     *mgo.Session
	collections []string
}

func NewDbConfig() *DbConfig {
	dbcfg := &DbConfig{collections: []string{"CYCLE", "STREAMER", "PLAYLIST"}}
	var err error
	dbcfg.session, err = mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	dbcfg.session.SetMode(mgo.Monotonic, true)
	return dbcfg
}

//ConnectionData type for future setup connect file
/*Type ConnectionData struct {
	Addrs []string
	Timeout int
	Username string
	Password string
	Database string
}*/

//Enum interface used for abstract the ConnectionSetup inputs
type Enum interface {
	SetSession()
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

//initialize the session above declared

//When it magic happens
//return mgo.Collection according to subsystem types
func (db *DbConfig) GetCollection(cs ConnectionSetup) *mgo.Collection {
	c := db.session.DB("radiup").C(db.collections[cs])
	fmt.Print("Collec")
	return c
}

//future func for get data of connect setup file
/*func (cs ConnectionSetup) GetConnection() ConnectionData {

	var message ConnectionData

	data, err := ioutil.ReadFile("db.config")
	if err != nil {
		panic(err)
	}

	err2 := json.Unmarshal(data, &message)
	if err2 != nil {
		panic(err)
	}

	fmt.Println(message.Addrs[0])
	return message

}*/

//close the session above declared and initialized
func (db *DbConfig) CloseSession() {
	db.session.Close()
}
