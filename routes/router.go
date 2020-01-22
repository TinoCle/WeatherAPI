package routes

import (
	"TPFinal/controllers"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func MapRoutes() {
	router.GET("/ping", controllers.Ping)
	router.GET("/ip", controllers.IP)
	router.GET("/location", controllers.Location)
}

func Run() {
	router.Run()
}
