package filialcontroller

import "filial-go/app/core/domain"

type controller struct {
	usecase domain.FilialUseCase
}

// New returns contract implementation of FilialController
func New(usecase domain.FilialUseCase) domain.FilialController {
	return &controller{
		usecase: usecase,
	}
}
