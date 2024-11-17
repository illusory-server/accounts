package account

import "github.com/illusory-server/accounts/internal/app/factory"

type (
	UseCase interface{}

	AccountsUseCase struct {
		accountFactory factory.AccountFactory
	}
)

func NewUseCase(
	accountFactory factory.AccountFactory,
) UseCase {
	return &AccountsUseCase{
		accountFactory: accountFactory,
	}
}
