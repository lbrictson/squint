package main

import (
	"github.com/lbrictson/squint/pkg/server"

	"github.com/lbrictson/squint/pkg/conf"
	"github.com/lbrictson/squint/pkg/db"
	"github.com/lbrictson/squint/pkg/store"
)

func main() {
	config, err := conf.New()
	if err != nil {
		panic(err)
	}
	p, err := db.New(config)
	if err != nil {
		panic(err)
	}
	s, err := store.New(p)
	if err != nil {
		panic(err)
	}
	app := server.New(server.NewServerConfig{Storage: s})
	app.Run()
}
