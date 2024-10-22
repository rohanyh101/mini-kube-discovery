package main

import (
	"log"
	"time"

	"github.com/roh4nyh/service_discovery/cache"
	"github.com/roh4nyh/service_discovery/model"
)

func NewCronJob(data model.Server, cycle time.Duration) {
	ticker := time.NewTicker(cycle)
	defer ticker.Stop()

	for {
		<-ticker.C

		if ok, err := helthCheck(data.HealthCheck); !ok {
			log.Printf("service %s is down with error: %s\n", data.Name, err)

			if err := cache.Remove(data); err != nil {
				log.Printf("Failed to remove service %s...", data.Name)
			}

			log.Printf("deleting... %s => %s\n", data.Name, data.HealthCheck)
			return
		}

		log.Printf("%s service is up and running...", data.Name)
	}
}
