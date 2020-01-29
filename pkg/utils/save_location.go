package utils

import (
	"TPFinal/pkg/domain"
)

type SaveLocationJob struct {
	toSave   domain.Locations
	saved    chan domain.Locations
	exitChan chan error
}

func NewSaveLocationJob(locations domain.Locations) *SaveLocationJob {
	return &SaveLocationJob{
		toSave:   locations,
		saved:    make(chan domain.Locations, 1),
		exitChan: make(chan error, 1),
	}
}
func (j SaveLocationJob) ExitChan() chan error {
	return j.exitChan
}
func (j SaveLocationJob) Run(locations map[string]domain.Locations) (map[string]domain.Locations, error) {
	var loc domain.Locations
	loc = j.toSave
	locations[loc.Id] = loc
	j.saved <- loc
	return locations, nil
}
