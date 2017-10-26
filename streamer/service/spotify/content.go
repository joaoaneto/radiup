package spotify

import (
	"github.com/joaoaneto/radiup/streamer/model"
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

	user, err := client.CurrentUser()
	if err != nil {
		return nil, err
	}

	userId := user.User.ID

	playlist, err := client.GetPlaylistsForUser(userId)
	if err != nil {
		return nil, err
	}

	return playlist, nil

}

/*Return tracks according to the name*/
func (cs *ContentSpotify) GetMusicData(client *spotify.Client, musicName string) ([]model.Music, error) {

	var musicsList []model.Music
	var music model.Music
	result, err := client.Search(musicName, spotify.SearchTypeTrack)
	if err != nil {
		return nil, err
	}

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

	return musicsList, nil
}
