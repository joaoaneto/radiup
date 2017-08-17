package model

import (
	"time"
	//"github.com/jinzhu/gorm"
	//"golang.org/x/oauth2"
)

type SpotifyToken struct {
	//gorm.Model
	ID				uint `gorm:"primary_key"`
	AccessToken 	string
	TokenType 		string
	RefreshToken	string
	Expiry			time.Time
	UserID			uint
}

type User struct {
	//gorm.Model
	UserID		 uint `gorm:"primary_key"`
	Name     string
	Username string
	Password string
	BirthDay time.Time
	Email    string
	Sex      string

	//Relationship
	SpotifyToken SpotifyToken `gorm:"ForeignKey:UserID"`

}