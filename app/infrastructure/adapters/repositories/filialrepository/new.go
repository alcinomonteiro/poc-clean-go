package filialrepository

import (
	"context"
	"filial-go/app/core/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	ctx        context.Context
	collection mongo.Collection
}

// New returns contract implementation of FilialRepository
func New(ctx context.Context, collection mongo.Collection) domain.FilialRepository {
	return &repository{
		ctx:        ctx,
		collection: collection,
	}
}
