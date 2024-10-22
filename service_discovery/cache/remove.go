package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/roh4nyh/service_discovery/model"
)

// Remove Server from cache
func Remove(data model.Server) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	return client.Del(ctx,
		fmt.Sprintf("%s:%s", data.Name, data.Uuid),
	).Err()
}
