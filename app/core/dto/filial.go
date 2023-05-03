package dto

import "encoding/json"

// CreateFilialCommand is an representation message listener to create a new Filial
type CreateFilialCommand struct {
	Name string `json:"name"`
	Code int32  `json:"code"`
	Type string `json:"type"`
}

// FromJSONCreateFilialCommand converts json message listener to a CreateFilialCommand struct
func FromJSONCreateFilialCommand(msg []byte) (*CreateFilialCommand, error) {
	createFilialCommand := CreateFilialCommand{}

	if err := json.Unmarshal(msg, &createFilialCommand); err != nil {
		return nil, err
	}

	return &createFilialCommand, nil
}

type CreateFilialInput struct {
	Name string `json:"name"`
	Code int32  `json:"code"`
	Type string `json:"type"`
}
