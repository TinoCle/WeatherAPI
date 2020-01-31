package services

import (
	"WeatherAPI/pkg/db"
	"WeatherAPI/pkg/domain"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	myip "github.com/polds/MyIP"
)

var weatherAPIKey string = "b64966af79891ad1f90c85de924bbe10"
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

func GetLocations() ([]domain.Locations, error) {
	/* client := db.GetClient() */
	locations, err := db.GetLocations()
	if err != nil {
		return locations, err
	}
	return locations, nil
}

func GetLocationID(id string) (domain.Locations, error) {
	location, err := db.GetLocationID(id)
	if err != nil {
		err = errors.New("Error al buscar la ubicación")
	}
	return location, err
}

func cleanQuery(params []string) string {
	var query string
	for _, param := range params {
		query += url.QueryEscape(param) + ","
	}
	return query
}

func CreateLocation(city, state, country string) (domain.Search, error) {
	var search []domain.Search
	cleaned := []string{city, state, country}
	query := cleanQuery(cleaned)
	url := "https://us1.locationiq.com/v1/search.php?key=" + key + "&q=" + query + "&format=json"
	resp, err := http.Get(url)
	if err != nil {
		err = errors.New("Error al buscar la ubicación")
		return search[0], err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			err = errors.New("Error al crear la ubicación")
			return search[0], err
		}
		json.Unmarshal(data, &search)
		_, err = db.GetLocationID(search[0].Id)
		if err == nil {
			return search[0], errors.New("La ubicación ya se encuentra en la lista")
		}
		aux := domain.Locations{
			Id:   search[0].Id,
			Name: search[0].Name,
			Lat:  search[0].Lat,
			Lon:  search[0].Lon,
		}
		locations[search[0].Id] = aux
		db.SaveLocation(aux)
		return search[0], nil
	}
	return search[0], errors.New("Error al buscar la ubicación")
}

func DeleteLocation(id string) error {
	_, err := db.GetLocationID(id)
	if err != nil {
		return errors.New("Error al borrar la ubicación")
	}
	err = db.DeleteLocation(id)
	if err != nil {
		return errors.New("Error al borrar la ubicación")
	}
	return nil
}

func UpdateLocation(id, name, lat, lon string) (domain.Locations, error) {
	location, err := db.GetLocationID(id)
	if err != nil {
		var locAux domain.Locations
		return locAux, errors.New("Error al actualizar la ubicación")
	}
	new := domain.Locations{
		Id:   location.Id,
		Name: name,
		Lat:  lat,
		Lon:  lon,
	}
	newLocation, err := db.SaveLocation(new)
	if err != nil {
		return newLocation, errors.New("Error al actualizar la ubicación")
	}
	return newLocation, nil
}

func GetWeather(lat string, lon string) (domain.Weather, error) {
	var err error
	var location domain.Location
	var weather domain.Weather
	if lat == "" || lon == "" {
		location, err = GetLocation()
		lat = location.Data.Latitud
		lon = location.Data.Longitud
	}
	if err != nil {
		err = errors.New("Error al obtener su ubicación")
		return weather, err
	}
	var url string = "http://api.openweathermap.org/data/2.5/weather?"
	url += "lat=" + lat
	url += "&lon=" + lon
	url += "&units=metric"
	url += "&appid=" + weatherAPIKey
	resp, err2 := http.Get(url)
	data, err3 := ioutil.ReadAll(resp.Body)
	if err2 != nil || err3 != nil {
		err = errors.New("Error al obtener el clima")
		return weather, err
	}
	json.Unmarshal(data, &weather)
	return weather, nil
}

func worker(wg *sync.WaitGroup, location domain.Locations, list *[]domain.Weather, index int) {
	weather, _ := GetWeather(location.Lat, location.Lon) // Obtiene el clima
	*list = append(*list, weather) 						 // Lo agrega al array
	defer wg.Done()
}

func GetAllWeathers() ([]domain.Weather, error){
	locations, err := GetLocations()
	var weather []domain.Weather
	if err != nil {
		return weather, err
	}
	if len(locations) == 0 {
		return weather, errors.New("No hay ubicaciones cargadas")
	}
	var wg sync.WaitGroup
	for i:=0; i<len(locations); i++ {
		wg.Add(1)
		go worker(&wg, locations[i], &weather, i) // Lanzo una goroutine por cada ubicación en la lista
	}
	wg.Wait() // Espero a que terminen todas las goroutines
	return weather, nil
}
