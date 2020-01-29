package utils

import (
	"TPFinal/pkg/domain"
	"errors"
)

type Client struct {
	Jobs chan Job
}

func (c *Client) SaveLocation(location domain.Locations) (domain.Locations, error) {
	job := NewSaveLocationJob(location)
	c.Jobs <- job

	if err := <-job.ExitChan(); err != nil {
		return domain.Locations{}, err
	}
	return <-job.saved, nil
}

func (c *Client) GetLocations() ([]domain.Locations, error) {
	arr := make([]domain.Locations, 0)
	locations, err := c.getLocationHash()
	if err != nil {
		return arr, err
	}

	for _, value := range locations {
		arr = append(arr, value)
	}
	return arr, nil
}

func (c *Client) GetLocationId(id string) (domain.Locations, error) {
	location, _ := c.getLocationHash()
	_, ok := location[id]
	if ok {
		return location[id], nil
	} else {
		return domain.Locations{}, errors.New("Not Found")
	}
}

func (c *Client) DeleteLocation(id string) error {
	location, _ := c.getLocationHash()
	_, ok := location[id]
	if ok {
		job := NewDeleteLocationJob(id)
		c.Jobs <- job
		if err := <-job.ExitChan(); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) getLocationHash() (map[string]domain.Locations, error) {
	job := NewReadLocationsJob()
	c.Jobs <- job
	if err := <-job.ExitChan(); err != nil {
		return make(map[string]domain.Locations, 0), err
	}
	return <-job.locations, nil
}
