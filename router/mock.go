package router

import "context"

type mockRouterImpl struct {
}

func NewMockRouter() (Router, error) {
	return &mockRouterImpl{}, nil
}

func (r *mockRouterImpl) Shutdown(ctx context.Context) error {
	return nil
}
