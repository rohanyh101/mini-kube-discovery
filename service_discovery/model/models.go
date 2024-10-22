package model

import "net"

type Server struct {
	Uuid        string
	Name        string
	Addr        net.TCPAddr
	HealthCheck string
}
