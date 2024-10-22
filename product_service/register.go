package main

import (
	"log"
	"net"
	"net/rpc"
	"strconv"
)

type ServiceDetail struct {
	Uuid        string
	Name        string
	Addr        net.TCPAddr
	HealthCheck string
}

var service ServiceDetail

func RegisterService() {
	var err error

	service.Name = "product"
	service.HealthCheck = "http://" + ip + ":" + port + "/health"
	service.Addr.IP = net.ParseIP(ip)
	service.Addr.Port, err = strconv.Atoi(port)
	if err != nil {
		panic(err)
	}

	client, err := rpc.Dial("tcp", "discovery:5555")
	if err != nil {
		panic(err)
	}

	// log.Printf("service, %+v\n", service)

	// register
	var uuid string
	if err := client.Call("Discovery.Register", service, &uuid); err != nil {
		panic(err)
	}

	log.Println("register success, uuid:", uuid)
}
