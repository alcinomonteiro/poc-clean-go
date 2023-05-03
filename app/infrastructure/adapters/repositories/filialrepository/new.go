package filialrepository

import (
	"filial-go/app/core/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	collection mongo.Collection
}

// New returns contract implementation of FilialRepository
func New(collection mongo.Collection) domain.FilialRepository {
	return &repository{
		collection: collection,
	}
}
