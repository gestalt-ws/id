package demesne

import (
	"context"

	"github.com/google/uuid"
)

// Repository interface for Demesne entity.
type repository interface {
	GetByID(ctx context.Context, id uuid.UUID) (Demesne, error)
	DeleteByID(ctx context.Context, id uuid.UUID) error
	Rename(ctx context.Context, id uuid.UUID, newName string) error
}
