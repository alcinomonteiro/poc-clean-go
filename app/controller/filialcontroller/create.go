package filialcontroller

import (
	"fmt"

	"filial-go/app/core/dto"
)

func (controller controller) Create(filialMessage []byte) error {

	createFilialCommand, err := dto.FromJSONCreateFilialCommand(filialMessage)
	if err != nil {
		return err
	}

	filial, err := controller.usecase.Create(createFilialCommand)
	if err != nil {
		return err
	}

	fmt.Println(filial)
	return nil
}
