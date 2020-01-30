package db

import (
	"TPFinal/pkg/domain"
)

type deleteLocationJob struct {
	toDelete string
	exitChan chan error
}

func newDeleteLocationJob(id string) *deleteLocationJob {
	return &deleteLocationJob{
		toDelete: id,
		exitChan: make(chan error, 1),
	}
}
func (j deleteLocationJob) ExitChan() chan error {
	return j.exitChan
}
func (j deleteLocationJob) Run(locations map[string]domain.Locations) (map[string]domain.Locations, error) {
	delete(locations, j.toDelete)
	return locations, nil
}
