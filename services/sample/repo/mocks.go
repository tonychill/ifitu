package repo

import (
	"context"
	"sync"

	"github.com/tonychill/ifitu/apis/pb/go/global"
)

type mockRepoImpl struct {
	sync.RWMutex
	index map[string]*global.Rate
}

func newMockRepository() (Repository, error) {
	return &mockRepoImpl{
		index: make(map[string]*global.Rate),
	}, nil
}

func (r *mockRepoImpl) Shutdown(ctx context.Context) error {
	return nil
}
