package services

import (
	"TPFinal/domain"
	"github.com/polds/MyIP"
	"errors"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

var poll domain.Poll
var count = make(map[string]int)
var total int
var pollResults domain.Poll

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