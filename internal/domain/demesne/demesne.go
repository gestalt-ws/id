/**
Copyright Gestalt Project 2025 - 2026. All rights reserved.
*/

package demesne

import (
	"context"
	"fmt"

	"github.com/gestaltid/internal/platform"
	"github.com/google/uuid"
)

// Demesne is the highest organisational unit within the Gestalt IAM hierarchy.
// It represents the parent IDP unit. For Managed Service Providers, a customer
// organisation's IAM configuration would be represented as a Realm, for example.
// This allows a single parent organisation to manage multiple, isolated,
// independent organisations from a unified place.
type Demesne struct {
	ID   uuid.UUID
	Name string
}

func NewDemesne(name string) (Demesne, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return Demesne{}, fmt.Errorf("failed to create new demesne: %v", err)
	}
	return Demesne{
		ID:   id,
		Name: name,
	}, nil
}

type Repository struct {
	pool platform.Pool
}

// GetByID retrieves a Demesne by its UUIDv7 ID value.
func (r Repository) GetByID(ctx context.Context, id uuid.UUID) (Demesne, error) {
	const q = `SELECT id, name FROM demesne WHERE id = $1`

	var d Demesne
	err := r.pool.QueryRow(ctx, q, id)
	if err != nil {
		return Demesne{}, fmt.Errorf("unable to get user with ID %s: %v", id, err)
	}

	return d, nil
}

// DeleteByID deletes a Demesne by its UUIDv7 ID value.
func (r Repository) DeleteByID(ctx context.Context, id uuid.UUID) error {
	const q = `DELETE FROM demesne WHERE id = $1`

	_, err := r.pool.Exec(ctx, q, id)
	if err != nil {
		return fmt.Errorf("unable to delete demesne with the ID: %s, %w", id, err)
	}

	return nil
}

func (r Repository) Rename(ctx context.Context, id uuid.UUID, newName string) error {
	//TODO implement me
	panic("implement me")
}
