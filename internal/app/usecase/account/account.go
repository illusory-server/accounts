package account

type (
	UseCase interface {
	}

	account struct{}
)

func NewUseCase() UseCase {
	return &account{}
}
