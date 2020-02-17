package controllers

import (
	"WeatherAPI/pkg/domain"
	"WeatherAPI/pkg/services"
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
		apiErr := ParseError(err)
		c.JSON(apiErr.Status, domain.Response{Mensaje: apiErr.Message})
		return
	} else if len(locations) == 0 {
		c.JSON(http.StatusNotFound, domain.Response{Mensaje: "No hay ubicaciones cargadas"})
		return
	}
	c.JSON(http.StatusOK, locations)
}

//GetLocationID trae una ubicación según su ID
func GetLocationID(c *gin.Context) {
	location, err := services.GetLocationID(c.Param("id"))
	if err != nil {
		apiErr := ParseError(err)
		c.JSON(apiErr.Status, domain.Response{Mensaje: apiErr.Message})
		return
	}
	c.JSON(http.StatusOK, location)
}

//PostLocation busca la localización del lugar pasado por el body
func PostLocation(c *gin.Context) {
	var body Body
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{Mensaje: "JSON inválido"})
		return
	}
	_, err := services.CreateLocation(body.City, body.State, body.Country)
	if err != nil {
		apiErr := ParseError(err)
		c.JSON(apiErr.Status, domain.Response{Mensaje: apiErr.Message})
		return
	}
	c.JSON(http.StatusCreated, domain.Response{Mensaje: "Ubicación agregada con éxito"})
}

func DeleteLocation(c *gin.Context) {
	err := services.DeleteLocation(c.Param("id"))
	if err != nil {
		apiErr := ParseError(err)
		c.JSON(apiErr.Status, domain.Response{Mensaje: apiErr.Message})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
func UpdateLocation(c *gin.Context) {
	var body domain.Locations
	if err := c.ShouldBindJSON(&body); err != nil {
		response := domain.Response{Mensaje: err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	_, err := services.UpdateLocation(body.Id, body.Name, body.Lat, body.Lon)
	if err != nil {
		apiErr := ParseError(err)
		response := domain.Response{Mensaje: apiErr.Message}
		c.JSON(apiErr.Status, response)
		return
	}
	c.JSON(http.StatusOK, domain.Response{Mensaje: "Ubicación actualizada con éxito"})
}

//GetWeather trae el clima de la ubicación actual
func GetWeather(c *gin.Context) {
	weather, err := services.GetWeather("", "")
	if err != nil {
		response := domain.Response{Mensaje: err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	c.JSON(http.StatusOK, weather)
}

func GetWeatherID(c *gin.Context) {
	weather, err := services.GetWeatherID(c.Param("id"))
	if err != nil {
		response := domain.Response{Mensaje: err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	c.JSON(http.StatusOK, weather)
}

func GetAllWeathers(c *gin.Context) {
	weathers, err := services.GetAllWeathers()
	if err != nil {
		response := domain.Response{Mensaje: err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	c.JSON(http.StatusOK, weathers)
}
