package rediser

import "github.com/redis/go-redis/v9"

type ConnectionClientFunc func() *redis.Client
type ConnectionClusterFunc func() *redis.ClusterClient
type ConnectionRingFunc func() *redis.Ring

type ClientConfig struct {
	Default     string
	Connections map[string]ConnectionClientFunc
}

type ClusterConfig struct {
	Default     string
	Connections map[string]ConnectionClusterFunc
}

type RingConfig struct {
	Default     string
	Connections map[string]ConnectionRingFunc
}

type Config struct {
	*ClientConfig
	*ClusterConfig
	*RingConfig
}
