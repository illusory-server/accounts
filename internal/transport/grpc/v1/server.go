package v1

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	v1 "github.com/illusory-server/accounts/gen/accounts/v1"
	"github.com/illusory-server/accounts/internal/app/usecase/account"
	"github.com/illusory-server/accounts/internal/domain/vo"
	"github.com/illusory-server/accounts/pkg/fn"
	"time"
)

type Server struct {
	accountUseCase account.UseCase
}

const timeFormat = time.RFC3339

func accountWithoutPassToTransport(acc *account.WithoutPassword) *v1.Account {
	return &v1.Account{
		Id:         acc.ID().Value(),
		FirstName:  acc.FirstName(),
		LastName:   acc.LastName(),
		Email:      acc.Email(),
		Nickname:   acc.Nickname(),
		Role:       acc.Role(),
		AvatarLink: acc.AvatarURL(),
		UpdatedAt:  acc.UpdatedAt().Format(timeFormat),
		CreatedAt:  acc.CreatedAt().Format(timeFormat),
	}
}

func (s *Server) Create(ctx context.Context, req *v1.CreateAccountRequest) (*v1.Account, error) {
	res, err := s.accountUseCase.Create(
		ctx,
		req.GetFirstName(),
		req.GetLastName(),
		req.GetEmail(),
		req.GetNickname(),
		req.GetPassword(),
	)
	if err != nil {
		return nil, err
	}
	return accountWithoutPassToTransport(res), err
}

func (s *Server) UpdateInfoById(ctx context.Context, req *v1.UpdateInfoByIdRequest) (*empty.Empty, error) {
	err := s.accountUseCase.UpdateInfoById(ctx, req.GetAccountId(), req.GetFirstName(), req.GetLastName())
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *Server) UpdateNicknameById(ctx context.Context, req *v1.UpdateNickByIdRequest) (*empty.Empty, error) {
	err := s.accountUseCase.UpdateNickById(ctx, req.GetAccountId(), req.GetNick())
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *Server) UpdateRoleById(ctx context.Context, req *v1.UpdateRoleByIdRequest) (*empty.Empty, error) {
	err := s.accountUseCase.UpdateRoleById(ctx, req.GetAccountId(), req.GetRole())
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *Server) UpdateEmailById(ctx context.Context, req *v1.UpdateEmailByIdRequest) (*empty.Empty, error) {
	err := s.accountUseCase.UpdateEmailById(ctx, req.GetAccountId(), req.GetEmail())
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *Server) UpdatePasswordById(ctx context.Context, req *v1.UpdatePasswordByIdRequest) (*empty.Empty, error) {
	err := s.accountUseCase.UpdatePasswordById(ctx, req.GetAccountId(), req.GetOldPassword(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *Server) AddAvatarLink(ctx context.Context, req *v1.AddAvatarLinkRequest) (*empty.Empty, error) {
	err := s.accountUseCase.AddAvatarLink(ctx, req.GetAccountId(), req.GetAvatarLink())
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *Server) DeleteById(ctx context.Context, id *v1.Id) (*empty.Empty, error) {
	err := s.accountUseCase.DeleteById(ctx, id.GetId())
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *Server) DeleteManyById(ctx context.Context, ids *v1.Ids) (*empty.Empty, error) {
	err := s.accountUseCase.DeleteManyByIds(ctx, ids.GetIds())
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *Server) GetAccountById(ctx context.Context, id *v1.Id) (*v1.Account, error) {
	res, err := s.accountUseCase.GetById(ctx, id.GetId())
	if err != nil {
		return nil, err
	}
	return accountWithoutPassToTransport(res), err
}

func (s *Server) GetAccountByEmail(ctx context.Context, str *v1.String) (*v1.Account, error) {
	res, err := s.accountUseCase.GetByEmail(ctx, str.GetValue())
	if err != nil {
		return nil, err
	}
	return accountWithoutPassToTransport(res), err
}

func (s *Server) GetAccountByNickname(ctx context.Context, str *v1.String) (*v1.Account, error) {
	res, err := s.accountUseCase.GetByNickname(ctx, str.GetValue())
	if err != nil {
		return nil, err
	}
	return accountWithoutPassToTransport(res), err
}

func (s *Server) GetAccountsByIds(ctx context.Context, ids *v1.Ids) (*v1.Accounts, error) {
	res, err := s.accountUseCase.GetByIds(ctx, ids.GetIds())
	if err != nil {
		return nil, err
	}
	return &v1.Accounts{
		Accounts: fn.Map(res, accountWithoutPassToTransport),
	}, nil
}

func convertOrder(order v1.QueryOrder) vo.QueryOrder {
	switch order {
	case v1.QueryOrder_ASK:
		return vo.Asc
	case v1.QueryOrder_DESK:
		return vo.Desc
	}
	return ""
}

func (s *Server) GetAccountsByQuery(ctx context.Context, req *v1.QueryRequest) (*v1.QueryAccountsResponse, error) {
	query, err := vo.NewQuery(
		uint(req.GetPage()),
		uint(req.GetLimit()),
		req.GetSortBy(),
		convertOrder(req.GetOrderBy()),
	)
	if err != nil {
		return nil, err
	}
	res, pageCount, err := s.accountUseCase.GetByQuery(ctx, query)
	if err != nil {
		return nil, err
	}
	return &v1.QueryAccountsResponse{
		Accounts:  fn.Map(res, accountWithoutPassToTransport),
		PageCount: uint64(pageCount),
	}, nil
}
