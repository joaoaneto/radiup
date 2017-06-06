package repository

import (
	"github.com/joaoaneto/radiup/streamer"
)

type OAuthInfoManager interface {
	Register(oAuth streamer.OAuthInfo)
	Update(clientId string, secretKey string)
	Remove(clientId string)
	Search(clientId string) streamer.OAuthInfo
}