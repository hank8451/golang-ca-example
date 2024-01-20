package main

import (
	"github.com/hank8451/cockroach_ca/config"
	"github.com/hank8451/cockroach_ca/database"
	"github.com/hank8451/cockroach_ca/server"
)

func main() {
	cfg := config.GetConfig()

	db := database.NewPostgresDatabase(&cfg)

	server.NewEchoServer(&cfg, db.GetDb()).Start()
}
