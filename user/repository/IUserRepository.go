package repository

import (
	"github.com/joaoaneto/radiup/user/model"
)

type SimpleUserManager interface {
	Create(su *model.SimpleUser) error
	Update(su *model.SimpleUser) error
	Remove(username string) error
	Search(username string) (model.SimpleUser, error)
	SearchAll() ([]model.SimpleUser, error)
}

type AdminUserManager interface {
	Create(au *model.AdminUser) error
	Update(au *model.AdminUser) error
	Remove(username string) error
	Search(username string) (model.AdminUser, error)
	SearchAll() ([]model.AdminUser, error)
}