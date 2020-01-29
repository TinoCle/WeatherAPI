package utils

import (
	"io/ioutil"
	"log"
	"sync"
)

const Db = "./db.json"

var client *Client
var doOnce sync.Once

func InitDb() {
	doOnce.Do(func() {
		// initialize empty-object json file if not found
		if _, err := ioutil.ReadFile(Db); err != nil {
			str := "{}"
			if err = ioutil.WriteFile(Db, []byte(str), 0644); err != nil {
				log.Fatal(err)
			}
		}
		// create channel to communicate over
		var jobs = make(chan Job)
		// start watching jobs channel for work
		go ProcessJobs(jobs, Db)
		// create dependencies
		client = &Client{Jobs: jobs}
	})
}

func GetClient() *Client {
	return client
}
