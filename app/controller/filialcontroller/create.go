package filialcontroller

import (
	"filial-go/app/core/dto"
)

func (controller controller) Create(filialMessage []byte) error {

	createFilialCommand, err := dto.FromJSONCreateFilialCommand(filialMessage)
	if err != nil {
		return err
	}

	if err = controller.usecase.Create(createFilialCommand); err != nil {
		return err
	}

	return nil
}
