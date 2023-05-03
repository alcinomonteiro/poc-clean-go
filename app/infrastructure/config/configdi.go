package config

import (
	"filial-go/app/controller/filialcontroller"
	"filial-go/app/core/domain"
	"filial-go/app/core/usecase/filialusecase"
	"filial-go/app/infrastructure/adapters/repositories/filialrepository"

	"go.mongodb.org/mongo-driver/mongo"
)

// ConfigFilialDI return a FilialController abstraction with dependency injection configuration
func ConfigFilialDI(collection *mongo.Collection) domain.FilialController {
	filialRepository := filialrepository.New(*collection)
	filialUseCase := filialusecase.New(filialRepository)
	filialController := filialcontroller.New(filialUseCase)

	return filialController
}
