package main

import (
	"log"

	"github.com/syndicatedb/vodka"
	"github.com/syndicatedb/vodka/adapters"
	"github.com/syndicatedb/vodka/example/controllers"
	"github.com/syndicatedb/vodka/example/modules/users"
)

type infrastructure struct {
	Postgres adapters.Adapter
}

var repos infrastructure
var config Config
var err error

var userModule *users.API

func init() {
	if config, err = NewConfig("./config.json"); err != nil {
		log.Fatalln("Error reading config: ", err)
	}
	repos.Postgres = adapters.NewPostgres(config.Postgres)

	userModule = users.New(repos.Postgres)
}

func main() {
	engine := vodka.New()
	engine.Server(config.HTTPServer)
	engine.Validation("./validation.json")

	userCtrl := controllers.NewUsers(userModule)

	engine.Router.GET("/users", userCtrl.Find)
	engine.Router.GET("/users/:id", userCtrl.FindByID)
	engine.Router.POST("/users", userCtrl.Create)
	engine.Router.PUT("/users/:id", userCtrl.UpdateByID)
	engine.Router.PUT("/users", userCtrl.Update)
	engine.Router.DELETE("/users/:id", userCtrl.DeleteByID)

	for {
		engine.Start()
	}
}
