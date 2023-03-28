package rediser

import "github.com/redis/go-redis/v9"

var instance *Manager

func SetInstance(i *Manager) {
	instance = i
}

func Instance() *Manager {
	return instance
}

func Connect(name ...string) redis.Cmdable {
	return Instance().Connect(name...)
}
