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

var (
	//API keys
	weatherAPIKey  string = "b64966af79891ad1f90c85de924bbe10"
	locationAPIkey string = "440d88bc9073b1"
	//errors
	ErrorLocationNotFound      = errors.New("Ubicación no encontrada")
	ErrorLocationAlreadyExists = errors.New("La Localización ya se encuentra registrada")
	ErrorCreateLocation        = errors.New("Error al crear la ubicación")
	ErrorDeleteLocation        = errors.New("Error al borrar la ubicación")
	ErrorUpdateLocation        = errors.New("Error al actualizar la ubicación")
)

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
	locations, err := db.GetLocations()
	if err != nil {
		return locations, err
	}
	return locations, nil
}

func GetLocationID(id string) (domain.Locations, error) {
	location, err := db.GetLocationID(id)
	if err != nil {
		return domain.Locations{}, err
	}
	if (domain.Locations{}) == location {
		return domain.Locations{}, ErrorLocationNotFound
	}
	return location, nil
}

func getUrl(params []string) string {
	var query string
	for _, param := range params {
		query += url.QueryEscape(param) + ","
	}
	url := "https://us1.locationiq.com/v1/search.php?key=" + locationAPIkey + "&q=" + query + "&format=json"
	return url
}

func CreateLocation(city, state, country string) (domain.Search, error) {
	url := getUrl([]string{city, state, country})
	resp, err := http.Get(url)
	if err != nil {
		return domain.Search{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return domain.Search{}, ErrorLocationNotFound
	}
	var search []domain.Search
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return domain.Search{}, ErrorCreateLocation
	}
	json.Unmarshal(data, &search)
	_, err = db.GetLocationID(search[0].Id)
	if err == nil {
		return domain.Search{}, ErrorLocationAlreadyExists
	}
	aux := domain.Locations{
		Id:   search[0].Id,
		Name: search[0].Name,
		Lat:  search[0].Lat,
		Lon:  search[0].Lon,
	}
	_, err = db.SaveLocation(aux)
	if err != nil {
		return domain.Search{}, ErrorCreateLocation
	}
	return search[0], nil
}

func DeleteLocation(id string) error {
	_, err := GetLocationID(id)
	if err != nil {
		return ErrorLocationNotFound
	}
	err = db.DeleteLocation(id)
	if err != nil {
		return ErrorDeleteLocation
	}
	return nil
}

func UpdateLocation(id, name, lat, lon string) (domain.Locations, error) {
	location, err := GetLocationID(id)
	if err != nil {
		return domain.Locations{}, ErrorLocationNotFound
	}
	new := domain.Locations{
		Id:   location.Id,
		Name: name,
		Lat:  lat,
		Lon:  lon,
	}
	newLocation, err := db.SaveLocation(new)
	if err != nil {
		return domain.Locations{}, ErrorUpdateLocation
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

func GetWeatherID(id string) (domain.Weather, error) {
	location, err := GetLocationID(id)
	if err != nil {
		return domain.Weather{}, ErrorLocationNotFound
	}
	weather, err := GetWeather(location.Lat, location.Lon)
	if err != nil {
		return weather, err
	}
	return weather, nil
}

func worker(wg *sync.WaitGroup, location domain.Locations, list *[]domain.Weather, index int) {
	weather, _ := GetWeather(location.Lat, location.Lon) // Obtiene el clima
	*list = append(*list, weather)                       // Lo agrega al array
	defer wg.Done()
}

func GetAllWeathers() ([]domain.Weather, error) {
	locations, err := GetLocations()
	var weather []domain.Weather
	if err != nil {
		return weather, err
	}
	if len(locations) == 0 {
		return weather, errors.New("No hay ubicaciones cargadas")
	}
	var wg sync.WaitGroup
	for i := 0; i < len(locations); i++ {
		wg.Add(1)
		go worker(&wg, locations[i], &weather, i) // Lanzo una goroutine por cada ubicación en la lista
	}
	wg.Wait() // Espero a que terminen todas las goroutines
	return weather, nil
}
