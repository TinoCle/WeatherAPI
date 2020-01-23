package services

import (
	"TPFinal/domain"
	"github.com/polds/MyIP"
	"errors"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"math"
)

var apiKey string = "b64966af79891ad1f90c85de924bbe10"

func GetIP() (string, error) {
	ip, err := myip.GetMyIP()
	if err != nil {
		err = errors.New("Error al obtener su IP")
	}
	return ip[:len(ip)-2], err
}

func GetLocation() (domain.Location, error) {
	ip, err := GetIP()
	var location domain.Location
	if err != nil {
		err = errors.New("Error al obtener su IP")
		return location, err;
	}
	resp, err2 := http.Get("http://ipvigilante.com/" + ip)
	data, err3 := ioutil.ReadAll(resp.Body)
	if err2 != nil || err3 != nil {
		err = errors.New("Error al obtener su ubicación")
		return location, err;
	}
	json.Unmarshal(data, &location);
	if location.Status != "success" {
		err = errors.New("Error al obtener su ubicación")
		return location, err;
	}
	return location, nil
}

func toCelsius(a float64) float64 {
	a = (a-32)/1.8 // to celsius
	return math.Round(a*10)/10 // rounded to 1 decimal
}

func GetWeather() (domain.Weather, error) {
	location, err := GetLocation()
	var weather domain.Weather
	if err != nil {
		err = errors.New("Error al obtener su ubicación")
		return weather, err;
	}
	var url string = "http://api.openweathermap.org/data/2.5/weather?"
	url += "lat=" + location.Data.Latitud
	url += "&lon=" + location.Data.Longitud
	url += "&units=metric"
	url += "&appid=" + apiKey
	resp, err2 := http.Get(url)
	data, err3 := ioutil.ReadAll(resp.Body)
	if err2 != nil || err3 != nil {
		err = errors.New("Error al obtener el clima")
		return weather, err;
	}
	json.Unmarshal(data, &weather);
	// weather.Detalle.Temp = toCelsius(weather.Detalle.Temp)
	// weather.Detalle.SensacionTermica = toCelsius(weather.Detalle.SensacionTermica)
	return weather, nil
}