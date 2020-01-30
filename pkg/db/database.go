package db

import (
	"TPFinal/pkg/domain"
	"errors"
	"io/ioutil"
	"log"
	"sync"
)

type handler struct {
	Jobs chan job
}

const db = "./db.json"

var client *handler
var doOnce sync.Once

//InitDb crea el archivo json para la db y lanza una go routine para escuchar las peticiones a la base de datos
func InitDb() {
	doOnce.Do(func() {
		// initialize empty-object json file if not found
		if _, err := ioutil.ReadFile(db); err != nil {
			str := "{}"
			if err = ioutil.WriteFile(db, []byte(str), 0644); err != nil {
				log.Fatal(err)
			}
		}
		// create channel to communicate over
		var jobs = make(chan job)
		// start watching jobs channel for work
		go processJobs(jobs, db)
		// create dependencies
		client = &handler{Jobs: jobs}
	})
}

//GetLocations obtiene todas las ubiaciones guardadas en la base de datos
func GetLocations() ([]domain.Locations, error) {
	arr := make([]domain.Locations, 0)
	locations, err := getLocationHash()
	if err != nil {
		return arr, err
	}

	for _, value := range locations {
		arr = append(arr, value)
	}
	return arr, nil
}

//GetLocationID busca la ubicacion de acuerdo al ID
func GetLocationID(id string) (domain.Locations, error) {
	location, _ := getLocationHash()
	_, ok := location[id]
	if ok {
		return location[id], nil
	}
	return domain.Locations{}, errors.New("Not Found")
}

//SaveLocation guarda o actualiza una ubicación en la db
func SaveLocation(location domain.Locations) (domain.Locations, error) {
	job := newSaveLocationJob(location)
	client.Jobs <- job

	if err := <-job.ExitChan(); err != nil {
		return domain.Locations{}, err
	}
	return <-job.saved, nil
}

//DeleteLocation elimina una ubicación en la db
func DeleteLocation(id string) error {
	location, _ := getLocationHash()
	_, ok := location[id]
	if ok {
		job := newDeleteLocationJob(id)
		client.Jobs <- job
		if err := <-job.ExitChan(); err != nil {
			return err
		}
	}
	return nil
}

func getLocationHash() (map[string]domain.Locations, error) {
	job := newReadLocationsJob()
	client.Jobs <- job
	if err := <-job.ExitChan(); err != nil {
		return make(map[string]domain.Locations, 0), err
	}
	return <-job.locations, nil
}
