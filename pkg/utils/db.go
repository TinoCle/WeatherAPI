package utils

import (
	"TPFinal/pkg/domain"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ReadDB() map[string]domain.Locations {
	var locations map[string]domain.Locations
	file, _ := ioutil.ReadFile("db.json")
	_ = json.Unmarshal([]byte(file), &locations)
	return locations
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
