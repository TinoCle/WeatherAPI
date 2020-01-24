package services

import (
	"TPFinal/pkg/domain"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	myip "github.com/polds/MyIP"
)

//API KEY: 440d88bc9073b1
var poll domain.Poll
var count = make(map[string]int)
var total int
var pollResults domain.Poll
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
		return location, err
	}
	resp, err2 := http.Get("http://ipvigilante.com/" + ip)
	data, err3 := ioutil.ReadAll(resp.Body)
	if err2 != nil || err3 != nil {
		err = errors.New("Error al obtener su ubicación")
		return location, err
	}
	json.Unmarshal(data, &location)
	if location.Status != "success" {
		err = errors.New("Error al obtener su ubicación")
		return location, err
	}
	return location, nil
}

func CreateLocation(city, state, country string) (domain.Search, error) {
	var search []domain.Search
	city = url.QueryEscape(city)
	state = url.QueryEscape(state)
	country = url.QueryEscape(country)
	url := "https://us1.locationiq.com/v1/search.php?key=440d88bc9073b1&q=" + city + "," + state + "," + country + "&format=json"
	resp, err := http.Get(url)
	if err != nil {
		err = errors.New("Error al buscar la localización")
		return search[0], err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(data, &search)

		aux := domain.Locations{
			Name: search[0].Name,
			Lat:  search[0].Lat,
			Lon:  search[0].Lon,
		}
		locations[search[0].Id] = aux
	}
	return search[0], nil
}

func DeleteLocation(id string) error {
	_, ok := locations[id]
	if ok {
		delete(locations, id)
		return nil
	}
	return errors.New("Failed to delete location")
}

func GetLocations() map[string]domain.Locations {
	return locations
}

func GetLocationId(id string) (domain.Locations, error) {
	location, ok := locations[id]
	fmt.Println(location)
	if ok {
		return location, nil
	}
	return location, errors.New("Failed to delete location")

}
