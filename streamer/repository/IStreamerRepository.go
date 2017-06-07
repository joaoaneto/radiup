package repository

import (
	"github.com/joaoaneto/radiup/streamer"
)

type OAuthInfoManager interface {
	Register(oAuth streamer.OAuthInfo) string
	Update(clientID string, secretKey string) string
	Remove(clientID string) string
	Search(clientID string) (streamer.OAuthInfo, string)
}
