package business

import (
	"fmt"

	"time"

	"github.com/joaoaneto/radiup/cycle"
	"github.com/joaoaneto/radiup/cycle/repository/mongo"
	streamerSpotify "github.com/joaoaneto/radiup/streamer/spotify"
	"github.com/zmb3/spotify"
)

type StreamerSuggestionDealer struct{}

func NewStreamerSuggestionDealer() StreamerSuggestionOperator {
	return &StreamerSuggestionDealer{}
}

func (dealer *StreamerSuggestionDealer) GetUpdatedMusicList(auth spotify.Authenticator) (cycle.StreamerSuggestion, error) {
	streamerSuggestionPersistor := mongo.NewPersistorStreamerSuggestion()
	sugg, err := streamerSuggestionPersistor.SearchAll()
	suggList := sugg.Musics

	if err != nil {
		return sugg, err
	}

	userPersistor := mongo.NewPersistorSimpleUser()
	userList, err := userPersistor.SearchAll()

	for _, a := range userList {
		client := auth.NewClient(a.AuthSpotify)
		socialSpotify := streamerSpotify.NewSocialSpotify()
		musics, err := socialSpotify.GetLastPlayedMusics(&client)

		if err != nil {
			continue
		}

		fmt.Println(a.SimpleUser.Name)

		for _, m := range musics {
			state := true

			for _, ms := range suggList {
				if m.ID == ms.ID {
					state = false
					break
				}
			}

			if state == true {

				suggList = append(suggList, m)

			}

		}
	}

	if len(sugg.Musics) == 0 {

		sugg = cycle.StreamerSuggestion{
			Musics:           suggList,
			ModificationDate: time.Now(),
		}

		streamerSuggestionPersistor.Register(sugg)

		sugg, err = streamerSuggestionPersistor.SearchAll()

		return sugg, err
	}

	streamerSuggestionPersistor.Update(sugg.ModificationDate, suggList)

	fmt.Println("PASSEI AQUI! 1")

	sugg, err = streamerSuggestionPersistor.SearchAll()

	fmt.Println("PASSEI AQUI! 1")

	return sugg, err
}
