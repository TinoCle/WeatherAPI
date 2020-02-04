package services

import (
	"WeatherAPI/pkg/db"
	"WeatherAPI/pkg/domain"
	"testing"
)

func init() {
	db.InitDb()
}

func TestGetIP(t *testing.T) {
	ip := "0.0.0.0"
	myIP, err := GetIP()
	if err != nil {
		t.Fail()
	}
	if myIP == ip {
		t.Fail()
	}
}

func TestGetLocations(t *testing.T) {
	_, err := GetLocations()
	if err != nil {
		t.Fail()
	}
}

func TestGetLocationsIDNotFound(t *testing.T) {
	location, err := GetLocationID("idUntraceable")
	if err != ErrorLocationNotFound {
		t.Fail()
	}
	if (domain.Locations{}) != location {
		t.Fail()
	}
}

func TestGetLocationsIDFound(t *testing.T) {
	aux := domain.Locations{
		Id:   "333086313822",
		Name: "Córdoba, Capital (Córdoba), Cordoba, Argentina",
		Lat:  "-31.4135",
		Lon:  "-64.18105",
	}
	db.SaveLocation(aux)
	found, err := GetLocationID(aux.Id)
	if (domain.Locations{}) == found {
		t.Fail()
	}
	if err == ErrorLocationNotFound {
		t.Fail()
	}
	if found != aux {
		t.Fail()
	}
}

func TestDeleteExistentLocation(t *testing.T) {
	aux := domain.Locations{
		Id:   "333086313822",
		Name: "Córdoba, Capital (Córdoba), Cordoba, Argentina",
		Lat:  "-31.4135",
		Lon:  "-64.18105",
	}
	db.SaveLocation(aux)
	err := DeleteLocation(aux.Id)
	if err != nil {
		t.Fail()
	}
}

func TestDeleteNonExistentLocation(t *testing.T) {
	id := "idUntraceable"
	err := DeleteLocation(id)
	if err == nil {
		t.Fail()
	}
}

func TestCreateLocationNew(t *testing.T) {
	city, state, country := "Chicago", "Illinois", "USA"
	location, err := CreateLocation(city, state, country)
	if err == ErrorLocationAlreadyExists {
		DeleteLocation(location.Id)
		_, err = CreateLocation(city, state, country)
		if err != nil {
			t.Fail()
		}
	}
	if err != nil && err != ErrorLocationAlreadyExists {
		t.Fail()
	}
}

func TestCreateExistentLocation(t *testing.T) {
	city, state, country := "Chicago", "Illinois", "USA"
	_, err := CreateLocation(city, state, country)
	if err != nil && err != ErrorLocationAlreadyExists {
		t.Fail()
	}
	_, err = CreateLocation(city, state, country)
	if err != ErrorLocationAlreadyExists {
		t.Fail()
	}
}
