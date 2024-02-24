package repo

import (
	"context"
)

func (r *repoImpl) Shutdown(ctx context.Context) error {
	r.shutdown = true
	return nil
}
