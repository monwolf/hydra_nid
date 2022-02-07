package contextx

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/ory/hydra/driver/config"
)

type (
	Contextualizer interface {
		Network(ctx context.Context, network uuid.UUID) uuid.UUID
		Config(ctx context.Context, config *config.Provider) *config.Provider
	}
	ContextualizerProvider interface {
		Contextualizer() Contextualizer
	}
	DefaultContextualizer struct{}
)

var _ Contextualizer = (*DefaultContextualizer)(nil)

func (d *DefaultContextualizer) Network(ctx context.Context, network uuid.UUID) uuid.UUID {
	return network
}

func (d *DefaultContextualizer) Config(ctx context.Context, config *config.Provider) *config.Provider {
	return config
}
