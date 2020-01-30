package db

import (
	"TPFinal/pkg/domain"
)

type saveLocationJob struct {
	toSave   domain.Locations
	saved    chan domain.Locations
	exitChan chan error
}

func newSaveLocationJob(locations domain.Locations) *saveLocationJob {
	return &saveLocationJob{
		toSave:   locations,
		saved:    make(chan domain.Locations, 1),
		exitChan: make(chan error, 1),
	}
}
func (j saveLocationJob) ExitChan() chan error {
	return j.exitChan
}
func (j saveLocationJob) Run(locations map[string]domain.Locations) (map[string]domain.Locations, error) {
	var loc domain.Locations
	loc = j.toSave
	locations[loc.Id] = loc
	j.saved <- loc
	return locations, nil
}
