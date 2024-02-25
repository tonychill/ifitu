package service

import (
	"context"
)

func (s *ServiceImpl) Shutdown(ctx context.Context) error {
	s.shutdown = true
	s.repo.Shutdown(ctx)
	return nil
}
