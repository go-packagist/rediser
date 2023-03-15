# Redis'er(Redis Manager)

[![Go Version](https://badgen.net/github/release/go-packagist/rediser/stable)](https://github.com/go-packagist/rediser/releases)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-packagist/rediser)](https://pkg.go.dev/github.com/go-packagist/rediser)
[![codecov](https://codecov.io/gh/go-packagist/rediser/branch/master/graph/badge.svg?token=5TWGQ9DIRU)](https://codecov.io/gh/go-packagist/rediser)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-packagist/rediser)](https://goreportcard.com/report/github.com/go-packagist/rediser)
[![tests](https://github.com/go-packagist/rediser/actions/workflows/go.yml/badge.svg)](https://github.com/go-packagist/rediser/actions/workflows/go.yml)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

## Installation

```bash
go get github.com/go-packagist/rediser
```

## Usage

```go
package main

import (
	"context"
	"github.com/go-packagist/rediser"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	m := rediser.New(&rediser.Config{
		ClientConfig: &rediser.ClientConfig{
			Default: "default",
			Connections: map[string]rediser.ConnectionClientFunc{
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
		ClusterConfig: &rediser.ClusterConfig{
			Default: "default",
			Connections: map[string]rediser.ConnectionClusterFunc{
				"default": func() *redis.ClusterClient {
					return redis.NewClusterClient(&redis.ClusterOptions{
						Addrs: []string{"localhost:6379"},
					})
				},
			},
		},
		RingConfig: &rediser.RingConfig{
			Default: "default",
			Connections: map[string]rediser.ConnectionRingFunc{
				"default": func() *redis.Ring {
					return redis.NewRing(&redis.RingOptions{
						Addrs: map[string]string{
							"shard1": "localhost:6379",
						},
					})
				},
			},
		},
	})

	// Example for Client
	m.Client().Get(ctx, "aaa").Val()       // use client default config
	m.Client("test").Get(ctx, "bbb").Val() // use client test config

	// Example for Cluster
	m.Cluster().Get(ctx, "aaa").Val() // use cluster default config

	// Example for Ring
	m.Ring().Get(ctx, "aaa").Val() // use ring default config
}

```

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.