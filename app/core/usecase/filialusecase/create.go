package filialusecase

import (
	"filial-go/app/core/domain"
	"filial-go/app/core/dto"
)

func (usecase usecase) Create(filialCommand *dto.CreateFilialCommand) (*domain.Filial, error) {

	filialInput := dto.CreateFilialInput{
		Name: filialCommand.Name,
		Code: filialCommand.Code,
		Type: filialCommand.Type,
	}

	filial, err := usecase.repository.Create(&filialInput)

	if err != nil {
		return nil, err
	}

	return filial, nil
}
