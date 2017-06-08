package repository

import (
	"github.com/joaoaneto/radiup/streamer"
)

type OAuthInfoManager interface {
	Register(oAuth streamer.OAuthInfo) error
	Update(clientID string, secretKey string) error
	Remove(clientID string) error
	Search(clientID string) (streamer.OAuthInfo, error)
}
