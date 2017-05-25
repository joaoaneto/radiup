package cycle

import (	
	"time"
)

//for storage the user information authentication
type AuthInfo struct{
	clientId string
	secretKey string
}

type User struct {
	name string
	username string
	password string //temp... we should search for a real alternative in web scenario
	birth_day time.Time
	email string
	sex byte
}

type SimpleUser struct{
	user User
	connectivity int
	authSpotify AuthInfo
	authDeezer AuthInfo
}

type AdminUser struct {
	user User
}