# Redis'er(Redis Manager)

![Go](https://badgen.net/badge/Go/%3E=1.18/green)
[![Go Version](https://badgen.net/github/release/go-packagist/rediser/stable)](https://github.com/go-packagist/rediser/releases)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-packagist/rediser/v2)](https://pkg.go.dev/github.com/go-packagist/rediser/v2)
[![codecov](https://codecov.io/gh/go-packagist/rediser/branch/master/graph/badge.svg?token=5TWGQ9DIRU)](https://codecov.io/gh/go-packagist/rediser)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-packagist/rediser)](https://goreportcard.com/report/github.com/go-packagist/rediser)
[![tests](https://github.com/go-packagist/rediser/actions/workflows/go.yml/badge.svg)](https://github.com/go-packagist/rediser/actions/workflows/go.yml)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

## Installation

```bash
go get github.com/go-packagist/rediser/v2
```

## Usage

```go
package main

import (
	"context"
	"github.com/go-packagist/rediser/v2"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	m := rediser.New(&rediser.Config{
		Default: "db1",
		Connections: map[string]rediser.Configable{
			"db1": &redis.Options{
				Addr:     "localhost:6379",
				Password: "", // no password set
				DB:       0,  // use default DB
			},
			"db2": &redis.Options{
				Addr:     "localhost:6379",
				Password: "", // no password set
				DB:       1,  // use default DB
			},
			"db3": &redis.ClusterOptions{
				Addrs: []string{"localhost:6379", "localhost:6379"},
			},
			"db4": &redis.RingOptions{
				Addrs: map[string]string{
					"server1": "localhost:6379",
					"server2": "localhost:6379",
				},
			},
		},
	}, rediser.WithInstance)

	m.Connect().Set(ctx, "aaa", "1", 0).Err()                  // use default(db1)
	m.Connect("db1").Set(ctx, "aaa", "1", 0).Err()             // db1
	m.Connect("db2").Set(ctx, "bbb", "1", 0).Err()             // db2
	rediser.Instance().Connect().Set(ctx, "ccc", "1", 0).Err() // use instance
	rediser.Connect().Set(ctx, "ddd", "1", 0).Err()            // use instance connect
}

```

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.