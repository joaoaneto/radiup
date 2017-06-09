package Mongo

import (
	"testing"

	"github.com/joaoaneto/radiup/streamer"
)

func TestRegisterOAuthInfo(t *testing.T) {
	i := NewPersistorOAuthInfo()
	newOAuth := streamer.OAuthInfo{ClientID: "testauth", SecretKey: "123456"}
	err := i.Register(newOAuth)

	if err != nil {
		t.Errorf("Register (OAuthInfo) fail.")
	}
}

func TestUpdateOAuthInfo(t *testing.T) {
	i := NewPersistorOAuthInfo()
	err := i.Update("testauth", "1234567")

	if err != nil {
		t.Errorf("Update (OAuthInfo) fail.")
	}
}

func TestRemoveOAuthInfo(t *testing.T) {
	i := NewPersistorOAuthInfo()
	err := i.Remove("testauth")

	if err != nil {
		t.Errorf("Remove (OAuthInfo) fail.")
	}
}

func TestSearchOAuthInfo(t *testing.T) {
	i := NewPersistorOAuthInfo()
	_, err := i.Search("testauth")

	if err != nil {
		t.Errorf("Search (OAuthInfo) fail.")
	}
}
