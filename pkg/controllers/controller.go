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

//PostLocation busca la localización del lugar pasado por el body
func PostLocation(c *gin.Context) {
	var body Body
	if err := c.ShouldBindJSON(&body); err != nil {
		response := domain.Response{Mensaje: err.Error()}
		c.JSON(http.StatusBadRequest, response)
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
		c.JSON(http.StatusNotFound, domain.Response{Mensaje: "Location not found"})
		return
	}
	c.JSON(http.StatusOK, domain.Response{Mensaje: "Eliminado"})
}

func GetLocations(c *gin.Context) {
	locations, err := services.GetLocations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, locations)
}

func GetLocationId(c *gin.Context) {
	location, err := services.GetLocationId(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, domain.Response{Mensaje: "Location not found"})
		return
	}
	c.JSON(http.StatusOK, location)
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
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, res)
}
