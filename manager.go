package rediser

import (
	"github.com/redis/go-redis/v9"
	"sync"
)

var (
	aliasClient  = "client"
	aliasCluster = "cluster"
	aliasRing    = "ring"
)

type Manager struct {
	config   *Config
	resolved sync.Map
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
	if db, ok := m.getResolved(aliasClient, name); ok {
		return db.(*redis.Client)
	}

	if _, ok := m.config.ClientConfig.Connections[name]; !ok {
		panic("client connection " + name + " is not defined")
	}

	reloved := m.config.ClientConfig.Connections[name]()
	m.setResolved(aliasClient, name, reloved)

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
	if db, ok := m.getResolved(aliasCluster, name); ok {
		return db.(*redis.ClusterClient)
	}

	if _, ok := m.config.ClusterConfig.Connections[name]; !ok {
		panic("cluster connection " + name + " is not defined")
	}

	reloved := m.config.ClusterConfig.Connections[name]()
	m.setResolved(aliasCluster, name, reloved)

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
	if db, ok := m.getResolved(aliasRing, name); ok {
		return db.(*redis.Ring)
	}

	if _, ok := m.config.RingConfig.Connections[name]; !ok {
		panic("ring connection " + name + " is not defined")
	}

	reloved := m.config.RingConfig.Connections[name]()
	m.setResolved(aliasRing, name, reloved)

	return reloved
}

// getResolved get resolved.
func (m *Manager) getResolved(prefix, name string) (interface{}, bool) {
	return m.resolved.Load(prefix + ":" + name)

}

// setResolved set resolved.
func (m *Manager) setResolved(prefix, name string, v interface{}) *Manager {
	m.resolved.Store(prefix+":"+name, v)

	return m
}
