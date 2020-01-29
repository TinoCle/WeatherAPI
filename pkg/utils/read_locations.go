package utils

import "TPFinal/pkg/domain"

type ReadLocationsJob struct {
	locations chan map[string]domain.Locations
	exitChan  chan error
}

func NewReadLocationsJob() *ReadLocationsJob {
	return &ReadLocationsJob{
		locations: make(chan map[string]domain.Locations, 1),
		exitChan:  make(chan error, 1),
	}
}
func (j ReadLocationsJob) ExitChan() chan error {
	return j.exitChan
}
func (j ReadLocationsJob) Run(locations map[string]domain.Locations) (map[string]domain.Locations, error) {
	j.locations <- locations
	return nil, nil
}
