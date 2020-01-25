package routes

import (
	"TPFinal/pkg/controllers"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func MapRoutes() {
	router.GET("/ping", controllers.Ping)
	router.GET("/ip", controllers.IP)
	router.GET("/location", controllers.Location)
	router.GET("/locations", controllers.GetLocations)
	router.GET("/location/:id", controllers.GetLocationId)
	router.POST("/location", controllers.PostLocation)
	router.DELETE("/location/:id", controllers.DeleteLocation)
	router.PUT("/location", controllers.UpdateLocation)
}

func Run() {
	router.Run()
}
