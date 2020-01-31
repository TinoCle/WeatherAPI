package routes

import (
	"WeatherAPI/pkg/controllers"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func MapRoutes() {
	router.GET("/ping", controllers.Ping)
	router.GET("/ip", controllers.IP)
	router.GET("/location", controllers.GetLocation)
	router.GET("/locations", controllers.GetLocations)
	router.GET("/location/:id", controllers.GetLocationID)
	router.POST("/location", controllers.PostLocation)
	router.DELETE("/location/:id", controllers.DeleteLocation)
	router.PUT("/location", controllers.UpdateLocation)
	router.GET("/weather", controllers.GetWeather)
	router.GET("/allWeathers", controllers.GetAllWeathers)
}

func Run() {
	router.Run()
}
