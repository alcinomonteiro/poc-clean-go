package filialusecase

import "filial-go/app/core/domain"

type usecase struct {
	repository domain.FilialRepository
}

// New returns contract implementation of FilialUseCase
func New(repository domain.FilialRepository) domain.FilialUseCase {
	return &usecase{
		repository: repository,
	}
}
