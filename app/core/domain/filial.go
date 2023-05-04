package domain

import (
	"fmt"

	"filial-go/app/core/dto"
)

// Filial is entity of table filial database column
type Filial struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Code int32  `json:"code"`
	Type string `json:"type"`
}

func (f Filial) String() string {
	return fmt.Sprintf("Filial{ID: %s,Name: %s,Code: %d,Type: %s}", f.ID, f.Name, f.Code, f.Type)
}

// FilialController is a Filial of me adapter layer
type FilialController interface {
	Create(filialMessage []byte) error
}

// FilialUseCase is a Filial of business rule layer
type FilialUseCase interface {
	Create(filialCommand *dto.CreateFilialCommand) error
}

// FilialRepository is a Filial of database connection adapter layer
type FilialRepository interface {
	Create(filialInput *dto.CreateFilialInput) error
}
