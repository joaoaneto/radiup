package mongo

import (
	"github.com/joaoaneto/radiup/cycle"
	"testing"
	"time"
)

// ContentSuggestion
func TestRegisterContentSuggestion(t *testing.T) {
	i := NewPersistorContentSuggestion()

	newUser := cycle.User{Name: "Alfredo Lucas", Username: "alf-lucas", Password: "seila",
		BirthDay: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), Email: "alf@upe.br", Sex: 'M'}

	newContentSuggestion := cycle.ContentSuggestion{Title: "Teste", Description: "description", ContentSuggestionUser: newUser,
		Votes: 0, Validated: false, Done: false}

	err := i.Register(newContentSuggestion)

	if err != nil {
		t.Error("Register (ContentSuggestion) fail.")
	}
}

func TestSearchContentSuggestion(t *testing.T) {
	i := NewPersistorContentSuggestion()
	_, err := i.Search("alf-lucas")

	if err != nil {
		t.Errorf("Search (ContentSuggestion) fail.")
	}
}

// Cycle

func TestCreateCycle(t *testing.T) {
	i := NewPersistorCycle()

	newUser := cycle.User{
		Name:     "Alfredo Lucas",
		Username: "alf-lucas",
		Password: "seila",
		BirthDay: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		Email:    "alf@upe.br",
		Sex:      'M',
	}

	newContentSuggestion := cycle.ContentSuggestion{
		Title:                 "Teste",
		Description:           "description",
		ContentSuggestionUser: newUser,
		Votes:     231,
		Validated: false,
		Done:      false,
	}

	newMusic := cycle.Music{Name: "Otherwise", Artist: []string{"Numsei"}, ID: "dasdas", SourceID: 2}

	newStreamerSuggestion := cycle.StreamerSuggestion{Musics: []cycle.Music{newMusic}}

	newVoluntarySuggestion := cycle.VoluntarySuggestion{VoluntarySuggestionUser: newUser, Votes: 123, Timestamp: time.Date(2015, 4, 2, 0, 15, 30, 918273645, time.UTC)}

	newCycle := cycle.Cycle{
		ID: 1, Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		End:                      time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		CycleType:                "Semanal",
		Description:              "description",
		CycleVoluntarySuggestion: newVoluntarySuggestion,
		CycleStreamerSuggestion:  newStreamerSuggestion,
		CycleContentSuggestion:   newContentSuggestion,
	}

	err := i.Create(newCycle)

	if err != nil {
		t.Errorf("Create (Cycle) fail.")
	}
}

func TestUpdateCycle(t *testing.T) {
	i := NewPersistorCycle()

	newUser := cycle.User{
		Name:     "Alfredo Lucas",
		Username: "alf-lucas",
		Password: "seila",
		BirthDay: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		Email:    "alf@upe.br",
		Sex:      'M',
	}

	newContentSuggestion := cycle.ContentSuggestion{
		Title:                 "Teste",
		Description:           "description",
		ContentSuggestionUser: newUser,
		Votes:     231,
		Validated: false,
		Done:      false,
	}

	newMusic := cycle.Music{Name: "Otherwise", Artist: []string{"Numsei"}, ID: "dasdas", SourceID: 2}

	newStreamerSuggestion := cycle.StreamerSuggestion{Musics: []cycle.Music{newMusic}}

	newVoluntarySuggestion := cycle.VoluntarySuggestion{VoluntarySuggestionUser: newUser, Votes: 123, Timestamp: time.Date(2015, 4, 2, 0, 15, 30, 918273645, time.UTC)}

	err := i.Update(1, time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		"Semanal", "description", newVoluntarySuggestion, newStreamerSuggestion, newContentSuggestion)

	if err != nil {
		t.Errorf("Update (Cycle) fail.")
	}
}

func TestRemoveCycle(t *testing.T) {
	i := NewPersistorCycle()
	err := i.Remove(1)

	if err != nil {
		t.Errorf("Remove (Cycle) fail.")
	}
}

func TestSearchCycle(t *testing.T) {
	i := NewPersistorCycle()
	_, err := i.Search(1)

	if err != nil {
		t.Errorf("Search (Cycle) fail.")
	}
}

// Music

func TestCreateMusic(t *testing.T) {
	i := NewPersistorMusic()
	
	newMusic := cycle.Music{Name: "Otherwise", Artist: []string{"Numsei"}, ID: "dasdas", SourceID: 2}

	err := i.Register(newMusic)

	if err != nil {
		t.Errorf("Create (Music) fail.")
	}
}

func TestRemoveMusic(t *testing.T) {
	i := NewPersistorMusic()
	err := i.Remove("dasdas")

	if err != nil {
		t.Errorf("Remove (Music) fail.")
	}
}

func TestSearchMusic(t *testing.T) {
	i := NewPersistorMusic()
	_, err := i.Search("sdasdaa")

	if err != nil {
		t.Errorf("Search (Music) fail.")
	}
}

// User

func TestCreateUser(t *testing.T) {
	i := NewPersistorUser()

	newUser := cycle.User{
		Name:     "Alfredo Lucas",
		Username: "alf-lucas",
		Password: "seila",
		BirthDay: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		Email:    "alf@upe.br",
		Sex:      'M',
	}

	err := i.Create(newUser)

	if err != nil {
		t.Errorf("Create (User) fail.")
	}
}

func TestRemoveUser(t *testing.T) {
	i := NewPersistorUser()
	err := i.Remove("alf-lucas")

	if err != nil {
		t.Errorf("Remove (User) fail.")
	}
}

func TestUpdateUser(t *testing.T) {
	i := NewPersistorUser()

	err := i.Update("Alfredo Lucas", "alflucas", "ops0",
		time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		"n sei@gmail.com", 'M')

	if err != nil {
		t.Errorf("Update (User) fail.")
	}
}
func TestSearchUser(t *testing.T) {
	i := NewPersistorUser()
	_, err := i.Search("alf-lucas")

	if err != nil {
		t.Errorf("Search (User) fail.")
	}
}

// VoluntarySuggestion

func TestRegisterVoluntarySuggestion(t *testing.T) {
	i := NewPersistorVoluntarySuggestion()

	newUser := cycle.User{Name: "Teste", Username: "testename"}

	newVoluntarySuggestion := cycle.VoluntarySuggestion{VoluntarySuggestionUser: newUser}

	err := i.Register(newVoluntarySuggestion)

	if err != nil {
		t.Errorf("Register (VoluntarySuggestion) fail.")
	}
}

func TestSearchVoluntarySuggestion(t *testing.T) {
	i := NewPersistorVoluntarySuggestion()
	_, err := i.Search("Teste")

	if err != nil {
		t.Errorf("Search (VoluntarySuggestion) fail.")
	}
}