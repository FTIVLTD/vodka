package controllers

import (
	"vodka/base"
	"vodka/example/modules/orders"
)

// Order - users controller struct
type Order struct {
	base.Controller
}

// NewOrders - users constructors
func NewOrders(m *orders.API) *Order {
	return &Order{
		Controller: base.NewController(m),
	}
}
