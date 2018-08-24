package orders

import (
	"github.com/FTIVLTD/vodka/adapters"
	"github.com/FTIVLTD/vodka/base"
	"github.com/FTIVLTD/vodka/repositories"
)

const source = "orders"

// API — external API for module
type API struct {
	base.Service
}

// Order — struct that describes User
type Order struct {
	ID     int64  `db:"id" key:"true" json:"id"`
	Name   string `db:"name" json:"name"`
	Status string `db:"status" json:"status"`
}

// New - module constructor
func New(adapter adapters.Adapter) *API {
	var u Order
	repo := repositories.NewMySQL(adapter, source, &u)
	return &API{
		Service: base.NewService(repo),
	}
}
