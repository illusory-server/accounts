package account

import "context"

func (a *AccountsUseCase) UpdateById(ctx context.Context, id, firstName, lastName, nick string) (*WithoutPassword, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AccountsUseCase) UpdateEmailById(ctx context.Context, id, email string) (*WithoutPassword, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AccountsUseCase) UpdatePasswordById(ctx context.Context, id, oldPassword, password string) (*WithoutPassword, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AccountsUseCase) UpdateRoleById(ctx context.Context, id, role string) (*WithoutPassword, error) {
	//TODO implement me
	panic("implement me")
}
