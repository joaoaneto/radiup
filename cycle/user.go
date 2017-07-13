package cycle

import (
	"time"

	"golang.org/x/oauth2"
)

type User struct {
	Name     string
	Username string
	Password []byte
	BirthDay time.Time
	Email    string
	Sex      string
}

type SimpleUser struct {
	SimpleUser  User
	AuthSpotify *oauth2.Token
	//AuthDeezer   *oauth2.Token
}

type AdminUser struct {
	AdminUser User
}
