package filialrepository

import (
	"context"
	"fmt"

	"filial-go/app/core/domain"
	"filial-go/app/core/dto"
)

func (repository repository) Create(filialInput *dto.CreateFilialInput) (*domain.Filial, error) {
	result, err := repository.collection.InsertOne(context.TODO(), filialInput)
	if err != nil {
		return nil, err
	}

	filial := domain.Filial{
		ID:   fmt.Sprintf("%v", result.InsertedID),
		Name: filialInput.Name,
		Code: filialInput.Code,
		Type: filialInput.Type,
	}
	return &filial, nil
}
