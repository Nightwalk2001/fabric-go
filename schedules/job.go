package schedules

import (
	"context"

	"fabric/redis"
)

func ResetId() {
	redis.Redis.Set(context.Background(), "idx", 0, 0)
}
