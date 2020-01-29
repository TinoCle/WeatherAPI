package utils

import (
	"TPFinal/pkg/domain"
)

type DeleteLocationJob struct {
	toDelete string
	exitChan chan error
}

func NewDeleteLocationJob(id string) *DeleteLocationJob {
	return &DeleteLocationJob{
		toDelete: id,
		exitChan: make(chan error, 1),
	}
}
func (j DeleteLocationJob) ExitChan() chan error {
	return j.exitChan
}
func (j DeleteLocationJob) Run(locations map[string]domain.Locations) (map[string]domain.Locations, error) {
	delete(locations, j.toDelete)
	return locations, nil
}
