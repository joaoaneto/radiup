package playlist

import (
	"fmt"

	"github.com/joaoaneto/radiup/cycle"
	"github.com/joaoaneto/radiup/cycle/business"
	"github.com/joaoaneto/radiup/cycle/repository/mongo"
	"github.com/joaoaneto/spotify"
)

type Playlist struct {
	PlaylistID int
	Musics     []cycle.Music
	Cycles     cycle.Cycle
	//	PlaylistRPC PlaylistRPC		??????
}

type PlaylistGenerator struct {
	Auth spotify.Authenticator
}

func NewPlaylistGenerator(auth spotify.Authenticator) business.CycleListener {
	return &PlaylistGenerator{Auth: auth}
}

func (pg *PlaylistGenerator) Notified(c *cycle.Cycle) {
	pg.GeneratePlaylist(c)
}

func (pg *PlaylistGenerator) GeneratePlaylist(c *cycle.Cycle) {

	persistorAdmin := mongo.NewPersistorAdminUser()
	admin, _ := persistorAdmin.Search("radiupapp")
	client := pg.Auth.NewClient(admin.AuthSpotify)
	playlist, _ := client.CreatePlaylistForUser(admin.AdminUser.Username, "Nova Playlist", true)
	playlistID := playlist.SimplePlaylist.ID
	var playlistSlice []cycle.Music
	vs := c.CycleVoluntarySuggestion
	ss := c.CycleStreamerSuggestion

	fmt.Println(vs)
	fmt.Println(ss)

	if len(vs) == 0 && len(ss.Musics) == 0 {
		return
	} else if len(vs) == 0 && len(ss.Musics) != 0 {
		fmt.Println("ROLA")
		fmt.Println(len(ss.Musics))
		playlistSlice = ss.Musics
	} else if len(vs) != 0 && len(ss.Musics) == 0 {
		for _, vsMusic := range vs {
			playlistSlice = append(playlistSlice, vsMusic.Track)
		}
	} else {
		for _, p := range vs {
			playlistSlice = append(playlistSlice, p.Track)
		}

		for _, strSugg := range ss.Musics {

			for j, pS := range playlistSlice {
				if strSugg.Name == pS.Name {
					break
				} else {
					if j == len(playlistSlice) {
						playlistSlice = append(playlistSlice, strSugg)
					}
				}
			}

		}
	}

	var playlistIDS []spotify.ID
	fmt.Println(playlistIDS)
	for _, play := range playlistSlice {
		teste := spotify.ID(play.ID)
		playlistIDS = append(playlistIDS, teste)
	}
	fmt.Println(len(playlistIDS))

	for _, add := range playlistIDS {
		client.AddTracksToPlaylist(admin.AdminUser.Username, playlistID,
			add)
	}
}
