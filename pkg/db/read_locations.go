package db

import "TPFinal/pkg/domain"

type readLocationsJob struct {
	locations chan map[string]domain.Locations
	exitChan  chan error
}

func newReadLocationsJob() *readLocationsJob {
	return &readLocationsJob{
		locations: make(chan map[string]domain.Locations, 1),
		exitChan:  make(chan error, 1),
	}
}
func (j readLocationsJob) ExitChan() chan error {
	return j.exitChan
}
func (j readLocationsJob) Run(locations map[string]domain.Locations) (map[string]domain.Locations, error) {
	j.locations <- locations
	return nil, nil
}
