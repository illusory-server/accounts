package account

import "context"

func (a *AccountsUseCase) UpdateInfoById(ctx context.Context, id, firstName, lastName string) (*WithoutPassword, error) {
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
