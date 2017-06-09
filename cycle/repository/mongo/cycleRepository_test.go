package mongo

import (
	"testing"

	"github.com/joaoaneto/radiup/cycle"
)

// ContentSuggestion
func TestRegisterContentSuggestion(t *testing.T) {
	i := NewPersistorContentSuggestion()
	newContentSuggestion := cycle.ContentSuggestion{Title: "Teste", Description: "description"}
	err := i.Register(newContentSuggestion)

	if err != nil {
		t.Error("Register (ContentSuggestion) fail.")
	}
}

func TestSearchContentSuggestion(t *testing.T) {
	i := NewPersistorContentSuggestion()
	suggestions, err := i.Search("Teste")

	if err != nil {
		t.Errorf("Search (ContentSuggestion) fail.")
	}
}

// Cycle

func TestCreateCycle(t *testing.T) {
	i := NewPersistorCycle()
	newCycle := cycle.Cycle{ID: 1, Description: "description"}
	err := i.Create(newCycle)

	if err != nil {
		t.Errorf("Create (Cycle) fail.")
	}
}

func TestUpdateCycle(t *testing.T) {
	i := NewPersistorCycle()
	err := i.Update(1, _, _, _, "newdescription", _, _, _)

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
	cycles, err := i.Search(1)

	if err != nil {
		t.Errorf("Search (Cycle) fail.")
	}
}

// Music

func TestCreateMusic(t *testing.T) {
	i := NewPersistorMusic()
	newMusic := cycle.Music{Name: "Teste", Id: 1}
	err := i.Create(newMusic)

	if err != nil {
		t.Errorf("Create (Music) fail.")
	}
}

func TestRemoveMusic(t *testing.T) {
	i := NewPersistorMusic()
	err := i.Remove("1")

	if err != nil {
		t.Errorf("Remove (Music) fail.")
	}
}

func TestSearchMusic(t *testing.T) {
	i := NewPersistorMusic()
	musics, err := i.Search(1)

	if err != nil {
		t.Errorf("Search (Music) fail.")
	}
}

// User

func TestCreateUser(t *testing.T) {
	i := NewPersistorUser()
	newUser := cycle.User{Name: "Teste"}
	err := i.Create(newUser)

	if err != nil {
		t.Errorf("Create (User) fail.")
	}
}

func TestRemoveUser(t *testing.T) {
	i := NewPersistorUser()
	err := i.Remove("Teste")

	if err != nil {
		t.Errorf("Remove (User) fail.")
	}
}

func TestUpdateUser(t *testing.T) {
	i := NewPersistorUser()
	err := i.Update("Teste", "Teste2", _, _, _, _)

	if err != nil {
		t.Errorf("Update (User) fail.")
	}
}
func TestSearchUser(t *testing.T) {
	i := NewPersistorUser()
	users := i.Search("Teste")

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
	suggestions, err := i.Search("Teste")

	if err != nil {
		t.Errorf("Search (VoluntarySuggestion) fail.")
	}
}
