package cycle

import (
	"time"
)

//for storage the user information authentication
type AuthInfo struct {
	AccessToken  string
	RefreshToken string
	TokenType    string
	TokenExpiry  string
}

type User struct {
	Name     string
	Username string
	Password []byte
	BirthDay time.Time
	Email    string
	Sex      string
}

type SimpleUser struct {
	SimpleUser   User
	Connectivity int
	AuthSpotify  AuthInfo
	AuthDeezer   AuthInfo
}

type AdminUser struct {
	AdminUser User
}
