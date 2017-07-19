package business

import (
	"fmt"

	"time"

	"github.com/joaoaneto/radiup/streamer"
	"github.com/joaoaneto/radiup/cycle"
	"github.com/joaoaneto/radiup/cycle/repository/mongo"
	streamerSpotify "github.com/joaoaneto/radiup/streamer/spotify"
	"github.com/joaoaneto/spotify"
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
		CheckTokenExpiry(a)
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

func CreateCycle (id int , startTime time.Time, endTime time.Time, cycleType string,
				  description string) {

	
	cycle := cycle.Cycle{
		ID: id, 
		Start: startTime,
		End: endTime,
		CycleType: cycleType,
		Description: description}
	
    /*Call cycle persistor*/
	cyclePersistor := mongo.NewPersistorCycle()

	/*Saving cycle*/
	cyclePersistor.Create(cycle)
}

type Dispatcher struct {
	listeners []CycleListener 
}

func (d* Dispatcher) AddListener(cl CycleListener) {
	d.listeners = append(d.listeners, cl)
}

func (d* Dispatcher) NotifyAll() {
	for _, m := range d.listeners {
		m.Notified()
	}
}

type CycleManager struct {
	Dis Dispatcher
}

func (cm* CycleManager) NewListener(cl CycleListener) {
	cm.Dis.AddListener(cl)
}

func (cm* CycleManager) Notify(){
	cm.Dis.NotifyAll()
}

func (cm* CycleManager)ManageCycle(c* cycle.Cycle) {
	
	for {
		now := time.Now()
		if now.Equal(c.End) || now.After(c.End) {
			cm.Notify()
			break
		}
	}
}

func CheckTokenExpiry(user cycle.SimpleUser) {

	spotifyStreamer := streamer.GetStreamerManager().Get("SPOTIFY")
	expiry := user.AuthSpotify.Expiry
	duration := time.Duration.Seconds(time.Since(expiry))
	
	if duration > 3600 {
		ntkn, _ := spotifyStreamer.AuthRPC.GetAuthenticator().RefreshToken(user.AuthSpotify.RefreshToken)
		userPersistor := mongo.NewPersistorSimpleUser()
		userPersistor.Update(user.SimpleUser.Username, user.SimpleUser.Name, user.SimpleUser.Password, user.SimpleUser.BirthDay, user.SimpleUser.Email, user.SimpleUser.Sex, ntkn)
	}

}