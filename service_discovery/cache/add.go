package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/roh4nyh/service_discovery/model"
)

func Add(data model.Server) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return client.HSet(ctx,
		fmt.Sprintf("%s:%s", data.Name, data.Uuid),
		"uuid", data.Uuid,
		"addr", data.Addr.String(),
		"health", data.HealthCheck,
	).Err()
}
