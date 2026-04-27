package todo

import "sync"

type Todo struct {
	ID   int
	Text string
}

type Repository interface {
	Save(t Todo) error
	FindAll() ([]Todo, error)
}

type MemoryRepo struct {
	mu   sync.RWMutex
	data []Todo
}

func (r *MemoryRepo) Save(t Todo) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data = append(r.data, t)
	return nil
}

func (r *MemoryRepo) FindAll() ([]Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]Todo, len(r.data))
	copy(out, r.data)
	return out, nil
}
