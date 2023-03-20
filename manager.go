package rediser

import (
	"github.com/redis/go-redis/v9"
	"sync"
)

// Config is the config of redis.
var alias = map[string]string{
	"client":  "client",
	"cluster": "cluster",
	"ring":    "ring",
}

type Manager struct {
	config  *Config
	reloved sync.Map
}

// New a redis manager.
func New(config *Config) *Manager {
	return &Manager{
		config: config,
	}
}

// Client get client.
func (m *Manager) Client(name ...string) *redis.Client {
	if len(name) > 0 {
		return m.resolveClient(name[0])
	}

	return m.resolveClient(m.config.ClientConfig.Default)
}

// resolveClient resolve client.
func (m *Manager) resolveClient(name string) *redis.Client {
	if db, ok := m.getReloved(alias["client"], name); ok {
		return db.(*redis.Client)
	}

	if _, ok := m.config.ClientConfig.Connections[name]; !ok {
		panic("client connection " + name + " is not defined")
	}

	reloved := m.config.ClientConfig.Connections[name]()
	m.setReloved(alias["client"], name, reloved)

	return reloved
}

// Cluster get cluster.
func (m *Manager) Cluster(name ...string) *redis.ClusterClient {
	if len(name) > 0 {
		return m.resolveCluster(name[0])
	}

	return m.resolveCluster(m.config.ClusterConfig.Default)
}

// resolveCluster resolve cluster.
func (m *Manager) resolveCluster(name string) *redis.ClusterClient {
	if db, ok := m.getReloved(alias["cluster"], name); ok {
		return db.(*redis.ClusterClient)
	}

	if _, ok := m.config.ClusterConfig.Connections[name]; !ok {
		panic("cluster connection " + name + " is not defined")
	}

	reloved := m.config.ClusterConfig.Connections[name]()
	m.setReloved(alias["cluster"], name, reloved)

	return reloved
}

// Ring get ring.
func (m *Manager) Ring(name ...string) *redis.Ring {
	if len(name) > 0 {
		return m.resolveRing(name[0])
	}

	return m.resolveRing(m.config.RingConfig.Default)
}

// resolveRing resolve ring.
func (m *Manager) resolveRing(name string) *redis.Ring {
	if db, ok := m.getReloved(alias["ring"], name); ok {
		return db.(*redis.Ring)
	}

	if _, ok := m.config.RingConfig.Connections[name]; !ok {
		panic("ring connection " + name + " is not defined")
	}

	reloved := m.config.RingConfig.Connections[name]()
	m.setReloved(alias["ring"], name, reloved)

	return reloved
}

// getReloved get reloved.
func (m *Manager) getReloved(prefix, name string) (interface{}, bool) {
	return m.reloved.Load(prefix + ":" + name)

}

// setReloved set reloved.
func (m *Manager) setReloved(prefix, name string, v interface{}) *Manager {
	m.reloved.Store(prefix+":"+name, v)

	return m
}
