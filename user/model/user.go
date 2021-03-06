package model

import (
	"time"
	//"github.com/jinzhu/gorm"
	//"golang.org/x/oauth2"
)

type Teste struct {
	Username string
}

type SpotifyToken struct {
	//gorm.Model
	ID           uint      `gorm:"primary_key"`
	AccessToken  string    `json:"access_token"`
	TokenType    string    `json:"token_type"`
	RefreshToken string    `json:"refresh_token"`
	Expiry       time.Time `json:"expiry"`
	UserID       uint
}

type User struct {
	//gorm.Model
	UserID   uint `gorm:"primary_key"`
	Name     string
	Username string
	Password string
	BirthDay time.Time
	Email    string
	Sex      string

	//Relationship
	SpotifyToken SpotifyToken `gorm:"ForeignKey:UserID"`
}
