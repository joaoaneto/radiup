package spotify

import(
	"github.com/zmb3/spotify"
)

type ContentSpotify struct {
	Playlist spotify.SimplePlaylistPage
}

func NewContentSpotify() *ContentSpotify {
	content := &ContentSpotify{}
	return content
}

/*Return all the user's playlists*/
func (cs *ContentSpotify) GetPlaylistData(client *spotify.Client)  (*spotify.SimplePlaylistPage, error) {
	user, error := client.CurrentUser()
	user_Id := user.User.ID
	playlist, _ := client.GetPlaylistsForUser(user_Id)
	return playlist, error
}


/*Return tracks according to the name*/
func (cs *ContentSpotify) GetMusicData(client *spotify.Client, musicName string) (*spotify.SearchResult, error){
	result, error := client.Search(musicName, spotify.SearchTypeTrack)
	return result, error
}