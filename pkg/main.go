package main

import (
	"TPFinal/pkg/routes"
	"TPFinal/pkg/db"
)

func main() {
	db.InitDb()
	routes.MapRoutes()
	routes.Run()
}
