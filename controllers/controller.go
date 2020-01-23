package controllers

import (
	"TPFinal/domain"
	"TPFinal/services"
	"net/http"
	"github.com/gin-gonic/gin"
)

// Setear esta variable en true si la API está hosteada online
var online = false

//Ping es para comprobar la disponibilidad de la API
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, domain.Response{Mensaje: "pong!"})
}

//IP es para obtener la IP pública de la máquina en que corre la API
func IP(c *gin.Context) {
	// Si la API corre localmente, busca su IP, sino busca la del cliente que hace la request
	if online {
		response := domain.Response{Mensaje: c.ClientIP()}
		c.JSON(http.StatusInternalServerError, response)
	} else {
		ip, err := services.GetIP()
		if err != nil {
			response := domain.Response{Mensaje: err.Error()}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		c.JSON(http.StatusOK, domain.Response{Mensaje: ip})
	}
}

//Location trae datos geográficos según la IP
func Location(c *gin.Context) {
	location, err := services.GetLocation()
	if err != nil {
		response := domain.Response{Mensaje: err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	c.JSON(http.StatusOK, location)
}

func Weather(c *gin.Context) {
	weather, err := services.GetWeather()
	if err != nil {
		response := domain.Response{Mensaje: err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	c.JSON(http.StatusOK, weather)
}