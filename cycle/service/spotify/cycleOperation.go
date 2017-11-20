package spotify

import (
	"log"

	"github.com/joaoaneto/radiup/cycle/repository/mongo"
)

type VoluntarySuggestionDealer struct{}

func NewVoluntarySuggestionDealer() VoluntarySuggestionOperator {
	return &VoluntarySuggestionDealer{}
}

func (daeler *VoluntarySuggestionDealer) HasVoluntarySuggestion(cycleID int, musicID string) bool {

	voluntarySuggestionPersistor := mongo.NewPersistorVoluntarySuggestion()

	_, err := voluntarySuggestionPersistor.Search(0, musicID)
	if err != nil {
		return false
	}

	return true

}

func (dealer *VoluntarySuggestionDealer) VerifyUserVote(cycleID int, musicID string, user string) (bool, error) {

	voluntarySuggestionPersistor := mongo.NewPersistorVoluntarySuggestion()

	sugg, err := voluntarySuggestionPersistor.Search(cycleID, musicID)
	if err != nil {
		log.Println(err)
		return false, err
	}

	for _, suggUser := range sugg.Users {
		if suggUser == user {
			return true, nil
		}
	}

	return false, nil

}

//type StreamerSuggestionDealer struct{}

/*func NewStreamerSuggestionDealer() StreamerSuggestionOperator {
	return &StreamerSuggestionDealer{}
}

func (dealer *StreamerSuggestionDealer) GetUpdatedMusicList(cycleID int, auth spotify.Authenticator) (cycle.StreamerSuggestion, error) {
	streamerSuggestionPersistor := mongo.NewPersistorStreamerSuggestion()
	sugg, err := streamerSuggestionPersistor.SearchAll(cycleID)
	suggList := sugg.Musics

	if err != nil {
		return sugg, err
	}

	userPersistor := mongo.NewPersistorSimpleUser()
	userList, err := userPersistor.SearchAll()

	fmt.Println(userList)

	for _, a := range userList {
		CheckTokenExpiry(a)
		client := auth.NewClient(a.AuthSpotify)
		musics, _ := streamerSpotify.NewSocialSpotify().GetLastPlayedMusics(&client)

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

	fmt.Println(len(suggList))

	if len(sugg.Musics) == 0 {

		sugg = cycle.StreamerSuggestion{
			Musics:           suggList,
			ModificationDate: time.Now(),
		}

		streamerSuggestionPersistor.Register(cycleID, sugg)

		sugg, err = streamerSuggestionPersistor.SearchAll(cycleID)

		return sugg, err
	}

	streamerSuggestionPersistor.Update(cycleID, suggList)

	sugg, err = streamerSuggestionPersistor.SearchAll(cycleID)

	return sugg, err
}*/

/*func CreateCycle(id int, startTime time.Time, endTime time.Time, cycleType string,
	description string) {

	cycle := cycle.Cycle{
		ID:          id,
		Start:       startTime,
		End:         endTime,
		CycleType:   cycleType,
		Description: description}

	cyclePersistor := mongo.NewPersistorCycle()

	cyclePersistor.Create(cycle)
}

type Dispatcher struct {
	listeners []CycleListener
}

func (d *Dispatcher) AddListener(cl CycleListener) {
	d.listeners = append(d.listeners, cl)
}

func (d *Dispatcher) NotifyAll(c *cycle.Cycle) {
	for _, m := range d.listeners {
		m.Notified(c)
	}
}

type CycleManager struct {
	Dis Dispatcher
}

func NewCycleManager() *CycleManager {
	return &CycleManager{}
}

func (cm *CycleManager) NewListener(cl CycleListener) {
	cm.Dis.AddListener(cl)
}

func (cm *CycleManager) Notify(c *cycle.Cycle) {
	cm.Dis.NotifyAll(c)
}

func (cm *CycleManager) ManageCycle() {
	c, _ := mongo.NewPersistorCycle().Search(0)
	for {
		now := time.Now()
		if now.Equal(c.End) || now.After(c.End) {
			c, _ = mongo.NewPersistorCycle().Search(0)
			cm.Notify(&c)
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

}*/
