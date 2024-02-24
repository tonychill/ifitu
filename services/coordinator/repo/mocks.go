package repo

import (
	"context"
	"sync"
)

type mockRepoImpl struct {
	sync.RWMutex
	flows map[string]any
}

func newMockRepository() (Repository, error) {
	return &mockRepoImpl{}, nil
}

func (r *mockRepoImpl) Shutdown(ctx context.Context) error {
	return nil
}
