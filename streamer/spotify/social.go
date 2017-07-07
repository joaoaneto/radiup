package spotify

import (
	"github.com/joaoaneto/radiup/cycle"
	"github.com/zmb3/spotify"
)

// SocialSpotify is a type for access to the implemented features of Spotify.
type SocialSpotify struct{}

// NewSocialSpotify is a constructor of SocialSpotify type
func NewSocialSpotify() *SocialSpotify {
	return &SocialSpotify{}
}

// GetInstant picks up the music the user is currently listening to.
func (s *SocialSpotify) GetInstant(client *spotify.Client) (cycle.Music, error) {
	current, err := client.PlayerCurrentlyPlaying()

	if current == nil {
		return cycle.Music{}, err
	}

	artistSpotify := current.Item.Artists
	var artistName []string

	for _, a := range artistSpotify {
		artistName = append(artistName, a.Name)
	}

	return cycle.Music{
		Name:     current.Item.Name,
		Artist:   artistName,
		ID:       current.Item.ID.String(),
		SourceID: 0,
	}, err

}

// GetLastPlayedMusics picks up a list of songs that the user has heard recently.
func (s *SocialSpotify) GetLastPlayedMusics(client *spotify.Client) ([]cycle.Music, error) {
	list, err := client.PlayerRecentlyPlayed()
	var musicList []cycle.Music

	if list == nil {
		return musicList, err
	}

	for _, a := range list {

		var artistName []string
		for _, b := range a.Track.Artists {
			artistName = append(artistName, b.Name)
		}

		newMusic := cycle.Music{
			Name:     a.Track.Name,
			Artist:   artistName,
			ID:       a.Track.ID.String(),
			SourceID: 0,
			PlayedAt: a.PlayedAt,
		}

		musicList = append(musicList, newMusic)

	}

	return musicList, err
}
