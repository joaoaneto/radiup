package Mongo

import (
	"testing"

	"github.com/joaoaneto/radiup/streamer"
)

func TestRegisterOAuthInfo(t *testing.T) {
	var i OAuthInfoManager
	i = NewPersistorOAuthInfo()
	newOAuth := streamer.OAuthInfo{ClientID: "testauth", SecretKey: "123456"}
	err := i.Register(newOAuth)

	if err {
		t.Errorf("Register (OAuthInfo) fail.")
	}
}

func TestUpdateOAuthInfo(t *testing.T) {
	var i OAuthInfoManager
	i := NewPersistorOAuthInfo()
	err := i.Update("testauth", "1234567")

	if err {
		t.Errorf("Update (OAuthInfo) fail.")
	}
}

func TestRemoveOAuthInfo(t *testing.T) {
	var i OAuthInfoManager
	i := NewPersistorOAuthInfo()
	err := i.Remove("testauth")

	if err {
		t.Errorf("Remove (OAuthInfo) fail.")
	}
}

//func TestSearchOAuthInfo(t *testing.T){}
func TestSearchOAuthInfo(t *testing.T) {
	var i OAuthInfoManager
	i := NewPersistorOAuthInfo()
	oAuthInfos, err := i.Search("testauth")

	if err {
		t.Errorf("Search (OAuthInfo) fail.")
	}
}
