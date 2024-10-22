package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/google/uuid"
	"github.com/roh4nyh/service_discovery/cache"
	"github.com/roh4nyh/service_discovery/model"
)

type Discovery int

func (d *Discovery) Register(data model.Server, reply *string) error {
	// validate
	if err := validateData(data); err != nil {
		return err
	}

	// insert
	data.Uuid = uuid.New().String()[0:8]
	// log.Printf("registering... %s => %s\n, ", data.Name, data.Uuid)

	if err := cache.Add(data); err != nil {
		return errors.New("failed to register service")
	}

	// triger cron job to check health...
	go NewCronJob(data, time.Second*5)

	*reply = data.Uuid
	return nil
}

func (d *Discovery) Get(name string, ret *net.TCPAddr) error {
	// validate
	if name == "" {
		return errors.New("missing service name")
	}

	// get the IP
	var err error
	*ret, err = cache.Get(name)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("service %s not found", name)
	}

	return nil
}
