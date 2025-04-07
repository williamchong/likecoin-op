package testutil

import (
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
)

func GetRedis(t *testing.T) (*redis.Client, redismock.ClientMock) {
	return redismock.NewClientMock()
}
