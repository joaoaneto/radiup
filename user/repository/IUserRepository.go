package repository

import (
	"github.com/joaoaneto/radiup/user/model"
)

type SimpleUserManager interface {
	Create(u *model.User) error
	Update(u *model.User) error
	Remove(username string) error
	Search(username string) (model.User, error)
	SearchAll() ([]model.User, error)
}
