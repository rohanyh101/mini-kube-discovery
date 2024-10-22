package cache

import (
	"context"
	"net"
)

// get server data...
func Get(name string) (net.TCPAddr, error) {
	var addr net.TCPAddr

	// keys
	keys, _, err := client.Scan(context.Background(), 0, name+"*", 10).Result()
	if err != nil {
		return addr, nil
	}

	// Check if any keys are found
	if len(keys) == 0 {
		// Return nil indicating no address found
		return addr, nil
	}

	for _, key := range keys {
		val, err := client.HGetAll(context.Background(), key).Result()
		if err != nil {
			return addr, nil
		}

		serverAddr, err := net.ResolveTCPAddr("tcp", val["addr"])
		if err != nil {
			return addr, nil
		}

		// If we find a valid address, return it
		return *serverAddr, nil
	}

	return addr, nil
}
