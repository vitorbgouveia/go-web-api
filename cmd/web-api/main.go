package main

import (
	"github.com/vitorbgouveia/go-web-api/configs/env"

	"github.com/vitorbgouveia/go-web-api/configs/database"

	"github.com/vitorbgouveia/go-web-api/internal/app/server"
)

func main() {
	env.Load()
	database.StartDB()
	server.Init()
}
