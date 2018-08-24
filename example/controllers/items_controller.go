package controllers

import (
	"github.com/FTIVLTD/vodka/base"
	"github.com/FTIVLTD/vodka/example/modules/items"
)

// Items - Items controller struct
type Items struct {
	base.Controller
}

// NewItems - Items constructors
func NewItems(m *items.API) *Items {
	return &Items{
		Controller: base.NewController(m),
	}
}
