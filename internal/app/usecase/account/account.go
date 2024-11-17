package account

type (
	UseCase interface{}

	AccountsUseCase struct {
	}
)

func NewUseCase() UseCase {
	return &AccountsUseCase{}
}
