package config

import (
	"context"
	"filial-go/app/controller/filialcontroller"
	"filial-go/app/core/domain"
	"filial-go/app/core/usecase/filialusecase"
	"filial-go/app/infrastructure/adapters/repositories/filialrepository"

	"go.mongodb.org/mongo-driver/mongo"
)

// ConfigFilialDI return a FilialController abstraction with dependency injection configuration
func ConfigFilialDI(ctx context.Context, collection *mongo.Collection) domain.FilialController {
	filialRepository := filialrepository.New(ctx, *collection)
	filialUseCase := filialusecase.New(filialRepository)
	filialController := filialcontroller.New(filialUseCase)

	return filialController
}
