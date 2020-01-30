package main

import (
	"WeatherAPI/pkg/routes"
	"WeatherAPI/pkg/db"
)

func main() {
	db.InitDb()
	routes.MapRoutes()
	routes.Run()
}
