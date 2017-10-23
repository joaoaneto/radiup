package model

import "golang.org/x/oauth2"

type TokenNotification struct {
	Token *oauth2.Token `json:"token"`
	User  string        `json:"user"`
}
