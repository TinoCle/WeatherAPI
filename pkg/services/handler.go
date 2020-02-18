package services

import (
	"WeatherAPI/pkg/db"
	"WeatherAPI/pkg/domain"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"

	myip "github.com/polds/MyIP"
)

var (
	//API keys
	weatherAPIKey  string = "b64966af79891ad1f90c85de924bbe10"
	locationAPIkey string = "440d88bc9073b1"

	myClient = &http.Client{Timeout: 10 * time.Second}
	//errors
	ErrorLocationNotFound      = errors.New("Ubicación no encontrada")
	ErrorLocationAlreadyExists = errors.New("La ubicación ya se encuentra registrada")
	ErrorCreateLocation        = errors.New("Error al crear la ubicación")
	ErrorDeleteLocation        = errors.New("Error al borrar la ubicación")
	ErrorUpdateLocation        = errors.New("Error al actualizar la ubicación")
	ErrorDB                    = errors.New("Error al obtener los datos de la base de datos")
	ErrorGetIP                 = errors.New("Error al obtener su IP")
	ErrorGetLocation           = errors.New("Error al obtener su ubicación")
	ErrorNoLocations           = errors.New("No hay ubicaciones cargadas")
)

//GetIP obtiene la dirección IP desde donde se lanzó la petición
func GetIP() (string, error) {
	ip, err := myip.GetMyIP()
	if err != nil {
		return "", ErrorGetIP
	}
	return ip[:len(ip)-2], nil
}

//GetLocation obtiene la ubicación de la IP desde donde se manda la request
func GetLocation() (domain.Location, error) {
	ip, err := GetIP()
	var location domain.Location
	if err != nil {
		return domain.Location{}, ErrorGetIP
	}
	resp, err2 := http.Get("http://ipvigilante.com/" + ip)
	data, err3 := ioutil.ReadAll(resp.Body)
	if err2 != nil || err3 != nil {
		return location, ErrorGetLocation
	}
	json.Unmarshal(data, &location)
	if location.Status != "success" {
		return location, ErrorGetLocation
	}
	return location, nil
}

//GetLocations obtiene todas las ubicaciones guardadas en la base de datos
func GetLocations() ([]domain.Locations, error) {
	locations, err := db.GetLocations()
	if err != nil {
		return locations, ErrorDB
	}
	return locations, nil
}

//GetLocationID obtiene una ubicación según su ID
func GetLocationID(id string) (domain.Locations, error) {
	location, err := db.GetLocationID(id)
	if err != nil {
		return domain.Locations{}, ErrorDB
	}
	if (domain.Locations{}) == location {
		return domain.Locations{}, ErrorLocationNotFound
	}
	return location, nil
}

func getURL(params []string) string {
	var query string
	for _, param := range params {
		query += url.QueryEscape(param) + ","
	}
	url := "https://us1.locationiq.com/v1/search.php?key=" + locationAPIkey + "&q=" + query + "&format=json"
	return url
}

func makeAPICall(url string) ([]domain.Search, error) {
	var (
		err     error
		resp    *http.Response
		retries = 5
	)
	for retries > 0 {
		resp, err = myClient.Get(url)
		if resp.StatusCode == 429 {
			time.Sleep(2 * time.Second)
			retries--
		} else if resp.StatusCode == http.StatusNotFound {
			return []domain.Search{}, ErrorLocationNotFound
		} else if err != nil {
			log.Fatalln(err)
			return []domain.Search{}, err
		} else {
			break
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	var search []domain.Search
	err = json.Unmarshal(body, &search)
	if err != nil {
		return []domain.Search{}, ErrorCreateLocation
	}
	return search, nil
}

//CreateLocation guarda una nueva ubicación en la base de datos si ya no se encuentra registrada
func CreateLocation(city, state, country string) (domain.Locations, error) {
	url := getURL([]string{city, state, country})
	search, err := makeAPICall(url)
	if err != nil {
		return domain.Locations{}, err
	}
	location, err := GetLocationID(search[0].Id)
	if err == nil {
		return location, ErrorLocationAlreadyExists
	}
	aux := domain.Locations{
		Id:   search[0].Id,
		Name: search[0].Name,
		Lat:  search[0].Lat,
		Lon:  search[0].Lon,
	}
	_, err = db.SaveLocation(aux)
	if err != nil {
		fmt.Println(err.Error())
		return domain.Locations{}, ErrorCreateLocation
	}
	return location, nil
}

//DeleteLocation borra una ubicación de la base de datos según su ID
func DeleteLocation(id string) error {
	_, err := GetLocationID(id)
	if err != nil {
		return err
	}
	err = db.DeleteLocation(id)
	if err != nil {
		return ErrorDeleteLocation
	}
	return nil
}

//UpdateLocation actualiza una ubicación de la base de datos según su ID
func UpdateLocation(id, name, lat, lon string) (domain.Locations, error) {
	location, err := GetLocationID(id)
	if err != nil {
		return domain.Locations{}, err
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
		return domain.Weather{}, err
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
		return weather, ErrorNoLocations
	}
	var wg sync.WaitGroup
	for i := 0; i < len(locations); i++ {
		wg.Add(1)
		go worker(&wg, locations[i], &weather, i) // Lanzo una goroutine por cada ubicación en la lista
	}
	wg.Wait() // Espero a que terminen todas las goroutines
	return weather, nil
}
