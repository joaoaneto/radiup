package repository 

import (
	"testing"
	"github.com/joaoaneto/radiup/playlist"
	"github.com/joaoaneto/radiup/cycle"
)

/*Tests in the Playlist interface's methods*/

TestCreatePlaylist(t *testing.T) {
	/*Creating variables Cycle and []Musics for playlist*/

	c := Cycle{ID: 1}
	var m []Music
	m = []Music {Music{Name: "teste"} , Music{Name: "teste2"} }

	playlistTest := Playlist{PlaylistID:1, Musics: m, Cycle:c} /*Creating a new Playlist for the test*/ 
	play := PlaylistPersistor{} /*New type that implement the interfaces*/
	err := play.Create(playlistTest) //Call Create method and get the err return

	/*Testing if the creation occured well*/
	if err != nil{
		t.Errorf("The Create method that must insert the playlist object on the DB, don't work well")
	}
}

TestUpdatePlaylist(t *testing.T) {
	play := PlaylistPersistor{}
	err := play.Update(2) /*Using the Update method with a id 2 in the parameter*/

	/*Testing if the update occured well*/
	if err != nil {
		t.Errorf("The playlist has not been updated with the Update method")	
	}
}

TestRemovePlaylist(t *testing.T) {
	play := PlaylistPersistor{}
	err := play.Remove(2)/*Trying to remove a playlist with id 2*/

	/*Testing if the object was removed*/
	if err != nil {
		t.Errorf("The playlist has not been removed with the Remove method")	
	}
}

TestSearchPlaylist(t *testing.T) {
	
	play := PlaylistPersistor{}
	playlist, err := play.Search(2)/*Trying to search a playlist with id 2*/

	/*Testing if the object was removed*/
	if err != nil {
		t.Errorf("The playlist has not been removed with the Remove method")	
	}
}
