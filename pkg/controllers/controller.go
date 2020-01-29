package controllers

import (
	"TPFinal/pkg/domain"
	"TPFinal/pkg/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Setear esta variable en true si la API está hosteada online
var online = false

type Body struct {
	City    string `json:"city" binding:"required"`
	State   string `json:"state"`
	Country string `json:"country"`
}

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

//GetLocation trae datos geográficos según la IP
func GetLocation(c *gin.Context) {
	location, err := services.GetLocation()
	if err != nil {
		response := domain.Response{Mensaje: err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	c.JSON(http.StatusOK, location)
}

//GetLocations trae las ubicaciones guardadas
func GetLocations(c *gin.Context) {
	locations, err := services.GetLocations()
	if err != nil {
		response := domain.Response{Mensaje: err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	} else if len(locations) == 0 {
		c.JSON(http.StatusOK, domain.Response{Mensaje: "No hay ubicaciones cargadas."})
		return
	}
	c.JSON(http.StatusOK, locations)
}

//GetLocationID trae una ubicación según su ID
func GetLocationID(c *gin.Context) {
	location, err := services.GetLocationID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, domain.Response{Mensaje: "No se encontró la ubicación"})
		return
	}
	c.JSON(http.StatusOK, location)
}

//PostLocation busca la localización del lugar pasado por el body
func PostLocation(c *gin.Context) {
	var body Body
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{Mensaje: "Json inválido."})
		return
	}
	res, err := services.CreateLocation(body.City, body.State, body.Country)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{Mensaje: "La Localización ya se encuentra registrada"})
		return
	}
	c.JSON(http.StatusOK, res)
}

func DeleteLocation(c *gin.Context) {
	err := services.DeleteLocation(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, domain.Response{Mensaje: "Ubicación no encontrada."})
		return
	}
	c.JSON(http.StatusOK, domain.Response{Mensaje: "Ubicación eliminada."})
}
func UpdateLocation(c *gin.Context) {
	var body domain.Locations

	if err := c.ShouldBindJSON(&body); err != nil {
		response := domain.Response{Mensaje: err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	fmt.Println(body)
	res, err := services.UpdateLocation(body.Id, body.Name, body.Lat, body.Lon)
	if err != nil {
		response := domain.Response{Mensaje: err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	c.JSON(http.StatusOK, res)
}

//GetWeather trae el clima de la ubicación actual
func GetWeather(c *gin.Context) {
	weather, err := services.GetWeather()
	if err != nil {
		response := domain.Response{Mensaje: err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	c.JSON(http.StatusOK, weather)
}
