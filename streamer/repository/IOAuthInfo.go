package repository

import (
	"radiup/streamer"
)

type IOAuthInfo interface {
	RegisterOAuthInfo()
	UpdateOAuthInfo(/*id int OU secretKey string*/)
	RemoveOAuthInfo(/*id int OU secretKey string*/)
	SearchOAuthInfo(/*id int OU secretKey string*/)
}