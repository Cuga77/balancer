package core

import (
	"balancer/pool"
	"balancer/types"
	"errors"
	"sync"
)

type RoundRobinBalancer struct {
	pool *pool.ServerPool
	mu   sync.Mutex
	idx  int
}

func NewRoundRobinBalancer(pool *pool.ServerPool) *RoundRobinBalancer {
	return &RoundRobinBalancer{
		pool: pool,
		idx:  -1,
	}
}

func (rb *RoundRobinBalancer) GetNextServer() (*types.Server, error) {
	servers := rb.pool.GetAllServers()

	if len(servers) == 0 {
		return nil, errors.New("No servers  found")
	}

	rb.mu.Lock()
	defer rb.mu.Unlock()

	rb.idx = (rb.idx + 1) % len(servers)
	selectedServer := servers[rb.idx]

	return selectedServer, nil
}
