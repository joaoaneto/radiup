package mongo

import (
	"log"
	"time"
	//"gopkg.in/mgo.v2"

	cycle "github.com/joaoaneto/radiup/cycle/model"
	cycleRep "github.com/joaoaneto/radiup/cycle/repository"
	"github.com/joaoaneto/radiup/dbconf"
	"golang.org/x/oauth2"
	"gopkg.in/mgo.v2/bson"
)

/*Implementation of Cycle's repository interfaces*/

//StreamerSuggestion Mongo implementations
//StreamerSuggestionPersistor .
type StreamerSuggestionPersistor struct {
	//db *dbconf.DbConfig
}

//NewPersistorStreamerSuggestion .
func NewPersistorStreamerSuggestion() cycleRep.StreamerSuggestionManager {
	return &StreamerSuggestionPersistor{}
}

// Register Streamer Suggestion
func (p StreamerSuggestionPersistor) Register(cycleID int, cs cycle.StreamerSuggestion) error {
	cyclePers := NewPersistorCycle()
	cycleActual, err := cyclePers.Search(cycleID)

	if err != nil {
		log.Fatal(err)
		return err
	}

	cycleActual.CycleStreamerSuggestion = cs

	err = cyclePers.Update(cycleID, cycleActual)

	return err
}

// SearchAll StreamerSuggestion
func (p StreamerSuggestionPersistor) SearchAll(cycleID int) (cycle.StreamerSuggestion, error) {
	cyclePers := NewPersistorCycle()
	actualCycle, err := cyclePers.Search(cycleID)
	var actualStreamerSuggestion cycle.StreamerSuggestion

	if err != nil {
		log.Fatal(err)
		return actualStreamerSuggestion, err
	}

	actualStreamerSuggestion = actualCycle.CycleStreamerSuggestion

	return actualStreamerSuggestion, err
}

// Update StreamerSuggestion
func (p StreamerSuggestionPersistor) Update(cycleID int, listMusic []cycle.Music) error {
	cyclePers := NewPersistorCycle()
	actualCycle, err := cyclePers.Search(cycleID)

	if err != nil {
		log.Fatal(err)
		return err
	}

	modif := time.Now()

	actualCycle.CycleStreamerSuggestion = cycle.StreamerSuggestion{Musics: listMusic, ModificationDate: modif}

	err = cyclePers.Update(cycleID, actualCycle)

	return err

}

/*ContentSuggestion Mongo implementations*/
type ContentSuggestionPersistor struct {
	//db *dbconf.DbConfig
}

func NewPersistorContentSuggestion() cycleRep.ContentSuggestionManager {
	return &ContentSuggestionPersistor{ /*dbconf.NewDbConfig()*/ }
}

func (p ContentSuggestionPersistor) Register(cycleID int, cs cycle.ContentSuggestion) error {

	//c := p.db.GetCollection(dbconf.CYCLE)
	c := NewPersistorCycle()

	cy, err := c.Search(cycleID)

	cy.CycleContentSuggestion = append(cy.CycleContentSuggestion, cs)

	c.Update(cycleID, cy)

	//err := c.Insert(&cs)

	if err != nil {
		log.Fatal(err)
	}

	return err

}

/*
func (p ContentSuggestionPersistor) Search(nameUser interface{}) ([]cycle.ContentSuggestion, error) {

	c := p.db.GetCollection(dbconf.CYCLE)

	result := []cycle.ContentSuggestion{}

	err := c.Find(bson.M{"name": nameUser}).One(&result)

	if err != nil {
		log.Fatal(err)
	}

	return result, err

}
*/

func (p ContentSuggestionPersistor) SearchAll(cycleID int) ([]cycle.ContentSuggestion, error) {

	//c := p.db.GetCollection(dbconf.CYCLE)

	c := NewPersistorCycle()

	cy, err := c.Search(cycleID)

	//result := []cycle.ContentSuggestion{}

	result := cy.CycleContentSuggestion
	//iter := c.Find(nil).Iter()
	//err := iter.All(&result)

	if err != nil {
		return nil, err
	}

	return result, err

}

/*Cycle Mongo implementations*/

type CyclePersistor struct {
	db *dbconf.DbConfig
}

func NewPersistorCycle() cycleRep.CycleManager {
	return &CyclePersistor{dbconf.NewDbConfig()}
}

func (cp CyclePersistor) Create(c cycle.Cycle) error {

	//defer session.Close()

	coll := cp.db.GetCollection(dbconf.CYCLE)

	err := coll.Insert(&c)

	if err != nil {
		log.Fatal(err)
	}

	return err

}

func (cp CyclePersistor) Update(registeredID int, updatedCycle cycle.Cycle) error {

	//defer session.Close()

	c := cp.db.GetCollection(dbconf.CYCLE)

	wantedCycle := bson.M{"id": updatedCycle.ID}

	changes := bson.M{"$set": bson.M{"start": updatedCycle.Start,
		"end":                      updatedCycle.End,
		"cycletype":                updatedCycle.CycleType,
		"description":              updatedCycle.Description,
		"cyclevoluntarysuggestion": updatedCycle.CycleVoluntarySuggestion,
		"cyclestreamersuggestion":  updatedCycle.CycleStreamerSuggestion,
		"cyclecontentsuggestion":   updatedCycle.CycleContentSuggestion}}

	err := c.Update(wantedCycle, changes)

	if err != nil {
		log.Fatal(err)
	}

	return err

}

func (cp CyclePersistor) Remove(id int) error {

	//defer session.Close()

	c := cp.db.GetCollection(dbconf.CYCLE)
	time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	err := c.Remove(bson.M{"id": id})

	if err != nil {
		log.Fatal(err)
	}

	return err

}

func (cp CyclePersistor) Search(id int) (cycle.Cycle, error) {

	result := cycle.Cycle{}

	//defer session.Close()

	c := cp.db.GetCollection(dbconf.CYCLE)

	err := c.Find(bson.M{"id": id}).One(&result)

	if err != nil {
		log.Fatal(err)
	}

	return result, err
}

/*Music Mongo implementations*/

type MusicPersistor struct {
	db *dbconf.DbConfig
}

func NewPersistorMusic() cycleRep.MusicManager {
	return &MusicPersistor{dbconf.NewDbConfig()}
}

func (mp MusicPersistor) Register(m cycle.Music) error {

	c := mp.db.GetCollection(dbconf.CYCLE)

	/*Insert the music object on DataBase*/
	err := c.Insert(&m)

	if err != nil {
		log.Fatal(err)
	}

	return err

}

func (mp MusicPersistor) Remove(id string) error {

	c := mp.db.GetCollection(dbconf.CYCLE)

	/*Insert the music object on DataBase*/
	err := c.Remove(bson.M{"id": id})

	if err != nil {
		log.Fatal(err)
	}

	return err

}

func (mp MusicPersistor) Search(id string) (cycle.Music, error) {

	result := cycle.Music{}

	c := mp.db.GetCollection(dbconf.CYCLE)

	/*Insert the music object on DataBase*/
	err := c.Find(bson.M{"id": id}).One(&result)

	if err != nil {
		log.Fatal(err)
	}

	return result, err

}

/*VoluntarySuggestion Mongo implementations*/

type VoluntarySuggestionPersistor struct {
	db *dbconf.DbConfig
}

func NewPersistorVoluntarySuggestion() cycleRep.VoluntarySuggestionManager {
	return &VoluntarySuggestionPersistor{dbconf.NewDbConfig()}
}

func (p VoluntarySuggestionPersistor) Register(cycleID int, vs cycle.VoluntarySuggestion) error {

	c := p.db.GetCollection(dbconf.CYCLE)
	err := c.Insert(&vs)
	/*cp := NewPersistorCycle()
	c, err := cp.Search(cycleID)*/

	//c.CycleVoluntarySuggestion = append(c.CycleVoluntarySuggestion, vs)
	//cp.Update(cycleID, c)

	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (p VoluntarySuggestionPersistor) SearchAll(cycleID int) ([]cycle.VoluntarySuggestion, error) {

	c := p.db.GetCollection(dbconf.CYCLE)

	result := []cycle.VoluntarySuggestion{}

	err := c.Find(nil).All(&result)

	if err != nil {
		log.Println("[SearchAll] - VoluntarySuggetions: ", err)
	}

	return result, nil

}

func (p VoluntarySuggestionPersistor) Search(cycleID int, musicID string) (cycle.VoluntarySuggestion, error) {

	c := p.db.GetCollection(dbconf.CYCLE)

	result := cycle.VoluntarySuggestion{}

	err := c.Find(bson.M{"track.id": musicID}).One(&result)

	return result, err
}

func (p VoluntarySuggestionPersistor) Update(vs cycle.VoluntarySuggestion) error {

	c := p.db.GetCollection(dbconf.CYCLE)

	wantedVs := bson.M{"track.id": vs.Track.ID}
	changes := bson.M{"$set": bson.M{"users": vs.Users, "votes": vs.Votes}}

	err := c.Update(wantedVs, changes)
	if err != nil {
		return err
	}

	return nil

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

/*
//User mongo implementations
type SimpleUserPersistor struct {
	db *dbconf.DbConfig
}

func NewPersistorSimpleUser() cycleRep.SimpleUserManager {
	return &SimpleUserPersistor{dbconf.NewDbConfig()}
}

func (up SimpleUserPersistor) Create(u cycle.SimpleUser) error {

	c := up.db.GetCollection(dbconf.CYCLE)

	err := c.Insert(&u)

	if err != nil {
		log.Fatal(err)
	}

	return err

}

func (up SimpleUserPersistor) Update(registered_user string,
	name string,
	password []byte,
	birthDay time.Time,
	email string,
	sex string,
	authSpotify *oauth2.Token) error {

	c := up.db.GetCollection(dbconf.CYCLE)

	wantedUser := bson.M{"simpleuser.username": registered_user}
	user := cycle.User{name, registered_user, password, birthDay, email, sex}
	changes := bson.M{"$set": bson.M{"simpleuser": user,
		"authspotify": authSpotify}}

	err := c.Update(wantedUser, changes)

	if err != nil {
		log.Fatal(err)
	}

	return err

}

func (up SimpleUserPersistor) Remove(username string) error {

	c := up.db.GetCollection(dbconf.CYCLE)

	err := c.Remove(bson.M{"username": username})

	if err != nil {
		log.Fatal(err)
	}

	return err

}

func (up SimpleUserPersistor) Search(username string) (cycle.SimpleUser, error) {

	result := cycle.SimpleUser{}

	c := up.db.GetCollection(dbconf.CYCLE)

	err := c.Find(bson.M{"simpleuser.username": username}).One(&result)

	if err != nil {
		log.Fatal(err)
	}

	return result, err
}

func (up SimpleUserPersistor) SearchAll() ([]cycle.SimpleUser, error) {

	result := []cycle.SimpleUser{}

	c := up.db.GetCollection(dbconf.CYCLE)

	iter := c.Find(bson.M{"simpleuser": bson.M{"$exists": true}}).Iter()
	err := iter.All(&result)

	if err != nil {
		return nil, err
	}

	return result, err
}

type AdminUserPersistor struct {
	db *dbconf.DbConfig
}

func NewPersistorAdminUser() cycleRep.AdminUserManager {
	return &AdminUserPersistor{dbconf.NewDbConfig()}
}

func (ap AdminUserPersistor) Create(u cycle.AdminUser) error {

	c := ap.db.GetCollection(dbconf.CYCLE)

	err := c.Insert(&u)

	if err != nil {
		log.Fatal(err)
	}

	return err

}

func (ap AdminUserPersistor) Update(user cycle.User,
	authSpotify *oauth2.Token) error {

	c := ap.db.GetCollection(dbconf.CYCLE)

	wantedUser := bson.M{"adminuser.username": user.Name}
	changes := bson.M{"$set": bson.M{"adminuser": user,
		"authspotify": authSpotify}}

	err := c.Update(wantedUser, changes)

	if err != nil {
		log.Fatal(err)
	}

	return err

}

func (ap AdminUserPersistor) Remove(username string) error {

	c := ap.db.GetCollection(dbconf.CYCLE)

	err := c.Remove(bson.M{"adminuser.username": username})

	if err != nil {
		log.Fatal(err)
	}

	return err

}

func (ap AdminUserPersistor) Search(username string) (cycle.AdminUser, error) {

	result := cycle.AdminUser{}

	c := ap.db.GetCollection(dbconf.CYCLE)

	err := c.Find(bson.M{"adminuser.username": username}).One(&result)

	if err != nil {
		log.Fatal(err)
	}

	return result, err
}

func (ap AdminUserPersistor) SearchAll() ([]cycle.AdminUser, error) {

	result := []cycle.AdminUser{}

	c := ap.db.GetCollection(dbconf.CYCLE)

	iter := c.Find(bson.M{"adminuser": bson.M{"$exists": true}}).Iter()
	err := iter.All(&result)

	if err != nil {
		return nil, err
	}

	return result, err
}
*/
