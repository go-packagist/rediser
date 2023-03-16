package rediser

import (
	"github.com/redis/go-redis/v9"
	"sync"
)

type Manager struct {
	config         *Config
	relovedClient  map[string]*redis.Client
	relovedCluster map[string]*redis.ClusterClient
	relovedRing    map[string]*redis.Ring
	rw             sync.RWMutex
}

func New(config *Config) *Manager {
	return &Manager{
		config:         config,
		relovedClient:  make(map[string]*redis.Client),
		relovedCluster: make(map[string]*redis.ClusterClient),
		relovedRing:    make(map[string]*redis.Ring),
		rw:             sync.RWMutex{},
	}
}

func (m *Manager) Client(name ...string) *redis.Client {
	if len(name) > 0 {
		return m.resolveClient(name[0])
	}

	return m.resolveClient(m.config.ClientConfig.Default)
}

func (m *Manager) resolveClient(name string) *redis.Client {
	if db, ok := m.relovedClient[name]; ok {
		return db
	}

	if _, ok := m.config.ClientConfig.Connections[name]; !ok {
		panic("client connection " + name + " is not defined")
	}

	m.rw.Lock()
	defer m.rw.Unlock()

	m.relovedClient[name] = m.config.ClientConfig.Connections[name]()

	return m.relovedClient[name]
}

func (m *Manager) Cluster(name ...string) *redis.ClusterClient {
	if len(name) > 0 {
		return m.resolveCluster(name[0])
	}

	return m.resolveCluster(m.config.ClusterConfig.Default)
}

func (m *Manager) resolveCluster(name string) *redis.ClusterClient {
	if db, ok := m.relovedCluster[name]; ok {
		return db
	}

	if _, ok := m.config.ClusterConfig.Connections[name]; !ok {
		panic("cluster connection " + name + " is not defined")
	}

	m.rw.Lock()
	defer m.rw.Unlock()

	m.relovedCluster[name] = m.config.ClusterConfig.Connections[name]()

	return m.relovedCluster[name]
}

func (m *Manager) Ring(name ...string) *redis.Ring {
	if len(name) > 0 {
		return m.resolveRing(name[0])
	}

	return m.resolveRing(m.config.RingConfig.Default)
}

func (m *Manager) resolveRing(name string) *redis.Ring {
	if db, ok := m.relovedRing[name]; ok {
		return db
	}

	if _, ok := m.config.RingConfig.Connections[name]; !ok {
		panic("ring connection " + name + " is not defined")
	}

	m.rw.Lock()
	defer m.rw.Unlock()

	m.relovedRing[name] = m.config.RingConfig.Connections[name]()

	return m.relovedRing[name]
}
