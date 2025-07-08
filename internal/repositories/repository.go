package repositories

import "sync"

type StatusRepo struct {
	mu      sync.RWMutex
	status  map[string]string
}

func NewStatusRepo() *StatusRepo {
	return &StatusRepo{status: make(map[string]string)}
}

func (r *StatusRepo) Get(name string) (string, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	s, ok := r.status[name]
	return s, ok
}

func (r *StatusRepo) Set(name, status string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.status[name] = status
}

func (r *StatusRepo) All() map[string]string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	copy := make(map[string]string)
	for k, v := range r.status {
		copy[k] = v
	}
	return copy
}
