package repository

import (
	"github.com/joaoaneto/radiup/streamer"
)

type IOAuthInfo interface {
	RegisterOAuthInfo(oAuth streamer.OAuthInfo)
	UpdateOAuthInfo(client_id string, secret_key string)
	RemoveOAuthInfo(client_id string)
	SearchOAuthInfo(client_id string) OAuthInfoRep
}