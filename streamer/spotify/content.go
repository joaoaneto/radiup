package spotify

<<<<<<< HEAD
import(
	"github.com/joaoaneto/spotify"
=======
import (
>>>>>>> 8931922901b81676e27fe1405ff30ed1b569bc44
	"github.com/joaoaneto/radiup/cycle"
	"github.com/joaoaneto/spotify"
)

type ContentSpotify struct {
	Playlist spotify.SimplePlaylistPage
}

func NewContentSpotify() *ContentSpotify {
	content := &ContentSpotify{}
	return content
}

/*Return all the user's playlists*/
func (cs *ContentSpotify) GetPlaylistData(client *spotify.Client) (*spotify.SimplePlaylistPage, error) {
	user, error := client.CurrentUser()
	user_Id := user.User.ID
	playlist, _ := client.GetPlaylistsForUser(user_Id)
	return playlist, error
}

/*Return tracks according to the name*/
func (cs *ContentSpotify) GetMusicData(client *spotify.Client, musicName string) ([]cycle.Music, error) {

  var musicsList []cycle.Music
  var music cycle.Music
  result, error := client.Search(musicName, spotify.SearchTypeTrack)

	for _, r := range result.Tracks.Tracks {
		var artistsList []string

		music.Name = r.SimpleTrack.Name
		music.ID = r.SimpleTrack.ID.String()

		for _, a := range r.SimpleTrack.Artists {
			artistsList = append(artistsList, a.Name)
		}

		music.Artist = artistsList

		musicsList = append(musicsList, music)
	}

	return musicsList, error
}
