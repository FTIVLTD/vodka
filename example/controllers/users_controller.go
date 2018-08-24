package controllers

import (
	"github.com/FTIVLTD/vodka/base"
	"github.com/FTIVLTD/vodka/example/modules/users"
)

// Users - users controller struct
type Users struct {
	base.Controller
}

// NewUsers - users constructors
func NewUsers(m *users.API) *Users {
	return &Users{
		Controller: base.NewController(m),
	}
}
