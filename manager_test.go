package rediser

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var ctx = context.Background()

func TestRediser(t *testing.T) {
	m := New(&Config{
		ClientConfig: &ClientConfig{
			Default: "default",
			Connections: map[string]ConnectionClientFunc{
				"default": func() *redis.Client {
					return redis.NewClient(&redis.Options{
						Addr:     "localhost:6379",
						Password: "", // no password set
						DB:       0,  // use default DB
					})
				},
				"test": func() *redis.Client {
					return redis.NewClient(&redis.Options{
						Addr:     "localhost:6379",
						Password: "", // no password set
						DB:       1,  // use default DB
					})
				},
			},
		},
	})

	// use db 0
	assert.Nil(t, m.Client().Set(ctx, "aaa", "1", time.Second*2).Err())
	assert.Equal(t, "1", m.Client().Get(ctx, "aaa").Val())

	// use db 1
	assert.Nil(t, m.Client("test").Set(ctx, "bbb", "2", time.Second*5).Err())
	assert.Equal(t, "2", m.Client("test").Get(ctx, "bbb").Val())

	// use db 0 to get db 1
	assert.Error(t, m.Client().Get(ctx, "bbb").Err())
	assert.Equal(t, "", m.Client().Get(ctx, "bbb").Val())

	// test sleep
	time.Sleep(time.Second * 3)
	assert.Equal(t, "", m.Client().Get(ctx, "aaa").Val())
}
