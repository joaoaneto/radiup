package repository

import (
	"github.com/joaoaneto/radiup/streamer"
)

type OAuthInfoManager interface {
	Register(oAuth streamer.OAuthInfo)
	Update(client_id string, secret_key string)
	Remove(client_id string)
	Search(client_id string) OAuthInfoRep
}