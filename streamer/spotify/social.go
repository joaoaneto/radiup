package spotify

import (
	"github.com/joaoaneto/radiup/cycle"
	"github.com/zmb3/spotify"
)

type SocialSpotify struct{}

func NewSocialSpotify() *SocialSpotify {
	return &SocialSpotify{}
}

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
