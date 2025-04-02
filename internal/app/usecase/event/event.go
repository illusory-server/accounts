package event

import "context"

var _ Event = (*UseCase)(nil)

type (
	Event interface {
		SendEvent(ctx context.Context) error
	}

	UseCase struct {
	}
)

func (u *UseCase) SendEvent(ctx context.Context) error {
	return nil
}

func NewUseCase() *UseCase {
	return &UseCase{}
}
