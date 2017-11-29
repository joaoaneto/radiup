package model

import (
	"time"
	//"github.com/jinzhu/gorm"
	//"golang.org/x/oauth2"
)

type SpotifyToken struct {
	SpotifyID    uint      `gorm:"primary_key"`
	AccessToken  string    `json:"access_token"`
	TokenType    string    `json:"token_type"`
	RefreshToken string    `json:"refresh_token"`
	Expiry       time.Time `json:"expiry"`
	FKUserID     uint
}

type User struct {
	UserID   uint `gorm:"primary_key"`
	Name     string
	Username string
	Password string
	Email    string

	//Relationship
	SpotifyToken SpotifyToken `gorm:"ForeignKey:FKUserID"`
}

type SimpleUser struct {
	User //composition
	BirthDay   time.Time
	Sex        string
}

type AdminUser struct {
	User //composition
	RadioName string
}