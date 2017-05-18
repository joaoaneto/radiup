package repository

import (
	"radiup/cycle"
)

type UserManager interface {
	CreateUser(u User)
	UpdateUser(username string)
	RemoveUser(username string)
	SearchUser(username string)
}