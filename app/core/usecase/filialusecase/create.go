package filialusecase

import (
	"filial-go/app/core/dto"
)

func (usecase usecase) Create(filialCommand *dto.CreateFilialCommand) error {

	filialInput := dto.CreateFilialInput{
		Name: filialCommand.Name,
		Code: filialCommand.Code,
		Type: filialCommand.Type,
	}

	err := usecase.repository.Create(&filialInput)

	if err != nil {
		return err
	}

	return nil
}
