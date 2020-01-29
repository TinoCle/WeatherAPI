package utils

import (
	"TPFinal/pkg/domain"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ReadDB() (map[string]domain.Locations, error) {
	var locations map[string]domain.Locations
	file, err := ioutil.ReadFile("db.json")
	json.Unmarshal([]byte(file), &locations)
	return locations, err
}

func SaveDB(locations map[string]domain.Locations) error {
	jsonString, err := json.Marshal(locations)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_ = ioutil.WriteFile("db.json", jsonString, 0644)
	return nil
}