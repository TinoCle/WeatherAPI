package services

import (
	"TPFinal/pkg/domain"
	"github.com/polds/MyIP"
	"TPFinal/pkg/utils"
	"errors"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"net/url"
	"fmt"
)

var weatherApiKey string = "b64966af79891ad1f90c85de924bbe10"
var key string = "440d88bc9073b1"

var locations = make(map[string]domain.Locations)

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

func GetLocations() (map[string]domain.Locations, error) {
	locations, err := utils.ReadDB()
	if err != nil {
		err = errors.New("Error al leer la base de datos.")
	}
	return locations, err
}

func GetLocationId(id string) (domain.Locations, error) {
	locations, _ := utils.ReadDB()
	location, ok := locations[id]
	if ok {
		return location, nil
	}
	return location, errors.New("Error al borrar la ubicación")
}

func CreateLocation(city, state, country string) (domain.Search, error) {
	var search []domain.Search
	city = url.QueryEscape(city)
	state = url.QueryEscape(state)
	country = url.QueryEscape(country)
	url := "https://us1.locationiq.com/v1/search.php?key=" + key + "&q=" + city + "," + state + "," + country + "&format=json"
	resp, err := http.Get(url)
	if err != nil {
		err = errors.New("Error al buscar la localización")
		return search[0], err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			err = errors.New("Error al crear la localización")
			return search[0], err
		}
		json.Unmarshal(data, &search)

		aux := domain.Locations{
			Name: search[0].Name,
			Lat:  search[0].Lat,
			Lon:  search[0].Lon,
		}
		locations[search[0].Id] = aux
		utils.SaveDB(locations)
	}
	return search[0], nil
}

func DeleteLocation(id string) error {
	locations, err := utils.ReadDB()
	if err != nil {
		return errors.New("Error al borrar la ubicación")
	}
	_, ok := locations[id]
	if ok {
		delete(locations, id)
		utils.SaveDB(locations)
		return nil
	}
	return errors.New("Error al borrar la ubicación")
}

func UpdateLocation(id, lat, lon string) (domain.Locations, error) {
	locations, err := utils.ReadDB()
	if err != nil {
		var locAux domain.Locations
		return locAux, errors.New("Error al actualizar la ubicación")
	}
	_, ok := locations[id]
	if ok {
		fmt.Println("UPDATING LOCATION")
		var x = locations[id]
		x.Lat = lat
		x.Lon = lon
		locations[id] = x
		jsonString, err := json.Marshal(locations)
		if err != nil {
			return locations[id], errors.New(err.Error())
		}
		_ = ioutil.WriteFile("db.json", jsonString, 0644)
		return locations[id], nil
	}
	return locations[id], errors.New("Error al actualizar la ubicación")
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
	url += "&appid=" + weatherApiKey
	resp, err2 := http.Get(url)
	data, err3 := ioutil.ReadAll(resp.Body)
	if err2 != nil || err3 != nil {
		err = errors.New("Error al obtener el clima")
		return weather, err;
	}
	json.Unmarshal(data, &weather);
	return weather, nil
}