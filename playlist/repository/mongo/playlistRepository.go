package mongo

import (
	"fmt"
	"log"
	"gopkg.in/mgo.v2/bson"
	"github.com/joaoaneto/radiup/dbconf"
	"github.com/joaoaneto/radiup/playlist"
	playlistRep "github.com/joaoaneto/radiup/playlist/repository"
)

/*Playlist Mongo implementations*/

type PlaylistPersistor struct {
	db *dbconf.DbConfig
}

func NewPersistorPlaylist() playlistRep.PlaylistManager {
	return &PlaylistPersistor{dbconf.NewDbConfig()}
}

func (pp PlaylistPersistor) Create(p playlist.Playlist) error {
	
    collection := pp.db.GetCollection(dbconf.PLAYLIST)
   
    err := collection.Insert(&p)

    if err != nil {
    	log.Fatal(err)
    }

	return err
}

func (pp PlaylistPersistor) Update(playlistID int,
								   musics []cycle.Music,
								   cycles cycle.Cycle) error {

    collection := pp.db.GetCollection(dbconf.PLAYLIST)
    
    wantedPlaylist := bson.M{"playlistid" : playlistID}

    changes := bson.M{"$set" : bson.M{"musics" : musics,
									  "cycles" : cycles}}

	err := c.Update(wantedPlaylist, changes)

	if err != nil {
    	log.Fatal(err)
    }

	return err
}

func (pp PlaylistPersistor) Remove(playlistID int) error {
	
    collection := pp.db.GetCollection(dbconf.PLAYLIST)

    err := c.Remove(bson.M{"playlistid" : playlistID})

    if err != nil {
    	log.Fatal(err)
    }

	return err
}

func (pp PlaylistPersistor) Search(playlistID int) (playlist.Playlist, error) {

	result := playlist.Playlist{}

    collection := pp.db.GetCollection(dbconf.PLAYLIST)

    err := c.Find(bson.M{"playlistid" : playlistID}).One(&result)

	if err != nil {
    	log.Fatal(err)
    }

    return result, err
}