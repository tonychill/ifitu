package repo

import (
	"context"
)

func (r *repoImpl) Shutdown(ctx context.Context) error {
	r.shutdown = true
	// TODO: flush redis and other db connections.
	return nil
}
