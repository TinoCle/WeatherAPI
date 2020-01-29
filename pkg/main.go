package main

import (
	"TPFinal/pkg/routes"
	"TPFinal/pkg/utils"
)

func main() {
	utils.InitDb()
	routes.MapRoutes()
	routes.Run()
}
