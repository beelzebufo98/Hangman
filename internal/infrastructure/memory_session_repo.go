package infrastructure

import (
	"fmt"
	"sync"

	"github.com/beelzebufo98/Hangman/internal/domain"
)

type MemorySessionRepo struct {
	data map[string]domain.Session
	mu   sync.RWMutex
}

func NewMemorySessionRepo() *MemorySessionRepo {
	return &MemorySessionRepo{data: make(map[string]domain.Session)}
}

func (r *MemorySessionRepo) Save(id string, session domain.Session) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[id] = session
	return nil
}

func (r *MemorySessionRepo) Load(id string) (domain.Session, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	sess, ok := r.data[id]
	if !ok {
		return domain.Session{}, fmt.Errorf("session %q not found", id)
	}
	return sess, nil
}
