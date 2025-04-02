package account

import (
	"context"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/illusory-server/accounts/internal/app/factory"
	"github.com/illusory-server/accounts/internal/domain/aggregate"
	"github.com/illusory-server/accounts/internal/domain/entity"
	"github.com/illusory-server/accounts/internal/domain/event"
	"github.com/illusory-server/accounts/internal/domain/repository"
	"github.com/illusory-server/accounts/internal/domain/vo"
	mockRepo "github.com/illusory-server/accounts/internal/mock/repo"
	"github.com/illusory-server/accounts/pkg/errors/codex"
	"github.com/illusory-server/accounts/pkg/errors/errx"
	"github.com/illusory-server/accounts/pkg/logger"
	"github.com/illusory-server/accounts/pkg/logger/log"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
	"time"
)

type commandCounter struct {
	repository.AccountCommand
	updateCount int
}

func (c *commandCounter) Update(ctx context.Context, account *aggregate.Account) error {
	c.updateCount++
	return c.AccountCommand.Update(ctx, account)
}

type timer struct {
	t time.Time
}

func (t *timer) Now() time.Time {
	return t.t
}

type genID struct {
	id string
}

func (g genID) GenerateID() string {
	if g.id == "" {
		g.id = uuid.New().String()
	}
	return g.id
}

func createUpdateInfoByIdTestAccount(t *testing.T, idp, firstName, lastName string, ti time.Time) *aggregate.Account {
	// Подготовка тестовых данных
	id, err := vo.NewID(idp)
	require.NoError(t, err)

	info, err := vo.NewAccountInfo(firstName, lastName, "test@example.com")
	require.NoError(t, err)

	role, err := vo.NewRole(vo.RoleUser)
	require.NoError(t, err)

	password, err := vo.NewPassword("secure_password123")
	require.NoError(t, err)

	acc, err := entity.NewAccount(
		id,
		info,
		role,
		"test_nickname",
		password,
		ti,
		ti,
	)
	require.NoError(t, err)

	agg, err := aggregate.NewAccount(acc)
	require.NoError(t, err)

	return agg
}

func setupLogger() (logger.Logger, *logger.OutMultiDump) {
	out := logger.NewOutMultiDump()
	l := log.NewLogger(&log.Options{
		Out:   out,
		Level: logger.DebugLvl,
	})
	return l, out
}

func TestAccountsUseCase_UpdateInfoById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.TODO()
	accFactory := factory.NewAccountFactory(&timer{}, genID{})

	type params struct {
		id, firstName, lastName string
	}

	createdAt := time.Now().Add(-time.Hour)
	eventTime := time.Now().Add(time.Millisecond)

	getInfoByParam := func(t *testing.T, p params, acc *aggregate.Account) vo.AccountInfo {
		info, err := vo.NewAccountInfo(p.firstName, p.lastName, acc.Account().Info().Email())
		assert.NoError(t, err)
		return info
	}

	anyErr := errors.New("any error")

	testCases := []struct {
		name              string
		params            params
		expectedErr       error
		expectedErrCode   codex.Code
		eventTime         time.Time
		logNotEmptyFields []string
		setupCommand      func(*testing.T, params) *mockRepo.MockAccountCommand
		setupQuery        func(*testing.T, params) *mockRepo.MockAccountQuery
	}{
		{
			name: "Should update account info by id",
			params: params{
				id:        uuid.New().String(),
				firstName: "UPDATED_John",
				lastName:  "UPDATED_Doe",
			},
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				agg := createUpdateInfoByIdTestAccount(t, p.id, p.firstName, p.lastName, createdAt)
				repo.EXPECT().GetById(gomock.Any(), p.id).Return(agg, nil)
				return repo
			},
			setupCommand: func(t *testing.T, p params) *mockRepo.MockAccountCommand {
				repo := mockRepo.NewMockAccountCommand(ctrl)
				agg := createUpdateInfoByIdTestAccount(t, p.id, p.firstName, p.lastName, createdAt)
				info := getInfoByParam(t, p, agg)
				err := agg.ChangeAccountInfo(info, eventTime)
				assert.NoError(t, err)
				repo.EXPECT().Update(gomock.Any(), agg).Return(nil)
				return repo
			},
		},
		{
			name: "Should handle update error",
			params: params{
				id:        uuid.New().String(),
				firstName: "UPDATED_John",
				lastName:  "UPDATED_Doe",
			},
			logNotEmptyFields: []string{"aggregate"},
			expectedErr:       errx.New(codex.Internal, "connect refuse"),
			expectedErrCode:   codex.Internal,
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				agg := createUpdateInfoByIdTestAccount(t, p.id, "current_f_name", "current_l_name", createdAt)
				repo.EXPECT().GetById(gomock.Any(), p.id).Return(agg, nil)
				return repo
			},
			setupCommand: func(t *testing.T, p params) *mockRepo.MockAccountCommand {
				repo := mockRepo.NewMockAccountCommand(ctrl)
				agg := createUpdateInfoByIdTestAccount(t, p.id, p.firstName, p.lastName, createdAt)
				info := getInfoByParam(t, p, agg)
				err := agg.ChangeAccountInfo(info, eventTime)
				assert.NoError(t, err)
				repo.EXPECT().Update(gomock.Any(), agg).Return(errx.New(codex.Internal, "connect refuse"))
				return repo
			},
		},
		{
			name: "Should handle incorrect first name update error",
			params: params{
				id:        uuid.New().String(),
				firstName: "i",
				lastName:  "UPDATED_Doe",
			},
			logNotEmptyFields: []string{"first_name", "last_name", "email"},
			expectedErr:       anyErr,
			expectedErrCode:   codex.InvalidArgument,
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				agg := createUpdateInfoByIdTestAccount(t, p.id, "current_f_name", "current_l_name", createdAt)
				repo.EXPECT().GetById(gomock.Any(), p.id).Return(agg, nil)
				return repo
			},
			setupCommand: func(t *testing.T, p params) *mockRepo.MockAccountCommand {
				return nil
			},
		},
		{
			name: "Should handle incorrect last name update error",
			params: params{
				id:        uuid.New().String(),
				firstName: "UPDATED_John",
				lastName:  "d",
			},
			logNotEmptyFields: []string{"first_name", "last_name", "email"},
			expectedErr:       anyErr,
			expectedErrCode:   codex.InvalidArgument,
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				agg := createUpdateInfoByIdTestAccount(t, p.id, "current_f_name", "current_l_name", createdAt)
				repo.EXPECT().GetById(gomock.Any(), p.id).Return(agg, nil)
				return repo
			},
			setupCommand: func(t *testing.T, p params) *mockRepo.MockAccountCommand {
				return nil
			},
		},
		{
			name: "Should handle incorrect get by id",
			params: params{
				id:        uuid.New().String(),
				firstName: "UPDATED_John",
				lastName:  "UPDATED_Doe",
			},
			logNotEmptyFields: []string{"id"},
			expectedErr:       errx.New(codex.NotFound, "user not found"),
			expectedErrCode:   codex.NotFound,
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				repo.EXPECT().GetById(gomock.Any(), p.id).Return(nil, errx.New(codex.NotFound, "user not found"))
				return repo
			},
			setupCommand: func(t *testing.T, p params) *mockRepo.MockAccountCommand {
				return nil
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			command := tc.setupCommand(t, tc.params)
			query := tc.setupQuery(t, tc.params)
			timeDep := &timer{
				t: eventTime,
			}

			counter := &commandCounter{
				AccountCommand: command,
			}
			l, dump := setupLogger()

			useCase, err := NewUseCase(l, accFactory, query, counter, timeDep)
			assert.NoError(t, err)

			err := useCase.UpdateInfoById(ctx, tc.params.id, tc.params.firstName, tc.params.lastName)
			if tc.expectedErr != nil {
				assert.Error(t, err)
				if !errors.Is(tc.expectedErr, anyErr) {
					assert.Equal(t, errors.Cause(tc.expectedErr).Error(), errors.Cause(err).Error())
				}
				c := errx.Code(err)
				assert.Equal(t, tc.expectedErrCode, c)
				assert.Equal(t, 1, len(dump.Dumps))
				logData := dump.Dumps[0]

				m := map[string]interface{}{}
				errU := json.Unmarshal(logData, &m)
				assert.NoError(t, errU)
				assert.NotEmpty(t, m["message"])
				assert.NotEmpty(t, m["error"])

				for _, f := range tc.logNotEmptyFields {
					assert.NotEmpty(t, m[f])
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, 1, counter.updateCount)
				if len(dump.Dumps) > 0 {
					logData := dump.Dumps[0]
					m := map[string]interface{}{}
					errU := json.Unmarshal(logData, &m)
					assert.NoError(t, errU)
					assert.NotEmpty(t, m["message"])
					data, ok := m["aggregate"]
					assert.True(t, ok)
					t.Log(data)
					temp, ok := data.(map[string]interface{})
					assert.True(t, ok)
					tmp := temp["events"].([]interface{})

					assert.Equal(t, []interface{}{string(event.AccountChangeInfoType)}, tmp)
				}
			}
		})
	}
}

func createUpdateEmailByIdTestAccount(t *testing.T, idp, email string, ti time.Time) *aggregate.Account {
	// Подготовка тестовых данных
	id, err := vo.NewID(idp)
	require.NoError(t, err)

	info, err := vo.NewAccountInfo("John", "Doe", email)
	require.NoError(t, err)

	role, err := vo.NewRole(vo.RoleUser)
	require.NoError(t, err)

	password, err := vo.NewPassword("secure_password123")
	require.NoError(t, err)

	acc, err := entity.NewAccount(
		id,
		info,
		role,
		"test_nickname",
		password,
		ti,
		ti,
	)
	require.NoError(t, err)

	agg, err := aggregate.NewAccount(acc)
	require.NoError(t, err)

	return agg
}

func TestAccountsUseCase_UpdateEmailById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.TODO()
	accFactory := factory.NewAccountFactory(&timer{}, genID{})

	type params struct {
		id, email string
	}

	createdAt := time.Now().Add(-time.Hour)
	eventTime := time.Now().Add(time.Millisecond)

	getInfoByParam := func(t *testing.T, p params, acc *aggregate.Account) vo.AccountInfo {
		info, err := vo.NewAccountInfo(acc.Account().Info().FirstName(), acc.Account().Info().LastName(), p.email)
		assert.NoError(t, err)
		return info
	}

	anyErr := errors.New("any error")

	testCases := []struct {
		name              string
		params            params
		expectedErr       error
		expectedErrCode   codex.Code
		eventTime         time.Time
		logNotEmptyFields []string
		setupCommand      func(*testing.T, params) *mockRepo.MockAccountCommand
		setupQuery        func(*testing.T, params) *mockRepo.MockAccountQuery
	}{
		{
			name: "Should correctly update account email",
			params: params{
				id:    uuid.New().String(),
				email: "correct@gmail.com",
			},
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				agg := createUpdateEmailByIdTestAccount(t, p.id, p.email, createdAt)
				repo.EXPECT().GetById(gomock.Any(), p.id).Return(agg, nil)
				return repo
			},
			setupCommand: func(t *testing.T, p params) *mockRepo.MockAccountCommand {
				repo := mockRepo.NewMockAccountCommand(ctrl)
				agg := createUpdateEmailByIdTestAccount(t, p.id, p.email, createdAt)
				info := getInfoByParam(t, p, agg)
				err := agg.ChangeEmail(info, eventTime)
				assert.NoError(t, err)
				repo.EXPECT().Update(gomock.Any(), agg).Return(nil)
				return repo
			},
		},
		{
			name: "Should handle update error",
			params: params{
				id:    uuid.New().String(),
				email: "correct@gmail.com",
			},
			logNotEmptyFields: []string{"aggregate"},
			expectedErr:       errx.New(codex.Internal, "connect refuse"),
			expectedErrCode:   codex.Internal,
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				agg := createUpdateEmailByIdTestAccount(t, p.id, p.email, createdAt)
				repo.EXPECT().GetById(gomock.Any(), p.id).Return(agg, nil)
				return repo
			},
			setupCommand: func(t *testing.T, p params) *mockRepo.MockAccountCommand {
				repo := mockRepo.NewMockAccountCommand(ctrl)
				agg := createUpdateEmailByIdTestAccount(t, p.id, p.email, createdAt)
				info := getInfoByParam(t, p, agg)
				err := agg.ChangeEmail(info, eventTime)
				assert.NoError(t, err)
				repo.EXPECT().Update(gomock.Any(), agg).Return(errx.New(codex.Internal, "connect refuse"))
				return repo
			},
		},
		{
			name: "Should handle incorrect email update error",
			params: params{
				id:    uuid.New().String(),
				email: "incorrect",
			},
			logNotEmptyFields: []string{"first_name", "last_name", "email"},
			expectedErr:       anyErr,
			expectedErrCode:   codex.InvalidArgument,
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				agg := createUpdateEmailByIdTestAccount(t, p.id, "correct@gmail.com", createdAt)
				repo.EXPECT().GetById(gomock.Any(), p.id).Return(agg, nil)
				return repo
			},
			setupCommand: func(t *testing.T, p params) *mockRepo.MockAccountCommand {
				return nil
			},
		},
		{
			name: "Should handle incorrect get by id",
			params: params{
				id:    uuid.New().String(),
				email: "correct@gmail.com",
			},
			logNotEmptyFields: []string{"id"},
			expectedErr:       errx.New(codex.NotFound, "user not found"),
			expectedErrCode:   codex.NotFound,
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				repo.EXPECT().GetById(gomock.Any(), p.id).Return(nil, errx.New(codex.NotFound, "user not found"))
				return repo
			},
			setupCommand: func(t *testing.T, p params) *mockRepo.MockAccountCommand {
				return nil
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			command := tc.setupCommand(t, tc.params)
			query := tc.setupQuery(t, tc.params)
			timeDep := &timer{
				t: eventTime,
			}

			counter := &commandCounter{
				AccountCommand: command,
			}
			l, dump := setupLogger()

			useCase, err := NewUseCase(l, accFactory, query, counter, timeDep)
			assert.NoError(t, err)

			err := useCase.UpdateEmailById(ctx, tc.params.id, tc.params.email)
			if tc.expectedErr != nil {
				assert.Error(t, err)
				if !errors.Is(tc.expectedErr, anyErr) {
					assert.Equal(t, errors.Cause(tc.expectedErr).Error(), errors.Cause(err).Error())
				}
				c := errx.Code(err)
				assert.Equal(t, tc.expectedErrCode, c)
				assert.Equal(t, 1, len(dump.Dumps))
				logData := dump.Dumps[0]

				m := map[string]interface{}{}
				errU := json.Unmarshal(logData, &m)
				assert.NoError(t, errU)
				assert.NotEmpty(t, m["message"])
				assert.NotEmpty(t, m["error"])

				for _, f := range tc.logNotEmptyFields {
					assert.NotEmpty(t, m[f])
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, 1, counter.updateCount)
				if len(dump.Dumps) > 0 {
					logData := dump.Dumps[0]
					m := map[string]interface{}{}
					errU := json.Unmarshal(logData, &m)
					assert.NoError(t, errU)
					assert.NotEmpty(t, m["message"])
					data, ok := m["aggregate"]
					assert.True(t, ok)
					t.Log(data)
					temp, ok := data.(map[string]interface{})
					assert.True(t, ok)
					tmp := temp["events"].([]interface{})
					assert.Equal(t, []interface{}{string(event.AccountChangeEmailType)}, tmp)
				}
			}
		})
	}
}

func createUpdatePasswordByIdTestAccount(t *testing.T, idp, pass string, ti time.Time) *aggregate.Account {
	// Подготовка тестовых данных
	id, err := vo.NewID(idp)
	require.NoError(t, err)

	info, err := vo.NewAccountInfo("John", "Doe", "correct@gmail.com")
	require.NoError(t, err)

	role, err := vo.NewRole(vo.RoleUser)
	require.NoError(t, err)

	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	assert.NoError(t, err)
	password, err := vo.NewPassword(string(hash))
	require.NoError(t, err)

	acc, err := entity.NewAccount(
		id,
		info,
		role,
		"test_nickname",
		password,
		ti,
		ti,
	)
	require.NoError(t, err)

	agg, err := aggregate.NewAccount(acc)
	require.NoError(t, err)

	return agg
}

func TestAccount_UpdatePasswordById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.TODO()
	accFactory := factory.NewAccountFactory(&timer{}, genID{})

	type params struct {
		id, old, new string
	}

	createdAt := time.Now().Add(-time.Hour)
	eventTime := time.Now().Add(time.Millisecond)

	anyErr := errors.New("any error")

	testCases := []struct {
		name              string
		params            params
		expectedErr       error
		expectedErrCode   codex.Code
		eventTime         time.Time
		logNotEmptyFields []string
		setupCommand      func(*testing.T, params) *mockRepo.MockAccountCommand
		setupQuery        func(*testing.T, params) *mockRepo.MockAccountQuery
	}{
		{
			name: "Should update password by id",
			params: params{
				id:  uuid.New().String(),
				old: "Old_password4142@",
				new: "New_password4142@",
			},
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				agg := createUpdatePasswordByIdTestAccount(t, p.id, p.old, createdAt)
				repo.EXPECT().GetById(gomock.Any(), p.id).Return(agg, nil)
				return repo
			},
			setupCommand: func(t *testing.T, p params) *mockRepo.MockAccountCommand {
				repo := mockRepo.NewMockAccountCommand(ctrl)
				agg := createUpdatePasswordByIdTestAccount(t, p.id, p.old, createdAt)
				pass, err := vo.NewPassword(p.new)
				require.NoError(t, err)
				err = agg.ChangePassword(pass, eventTime)
				assert.NoError(t, err)
				repo.EXPECT().Update(gomock.Any(), agg).Return(nil)
				return repo
			},
		},
		{
			name: "Should handle update error",
			params: params{
				id:  uuid.New().String(),
				old: "Old_password4142@",
				new: "New_password4142@",
			},
			logNotEmptyFields: []string{"aggregate"},
			expectedErr:       errx.New(codex.Internal, "connect refuse"),
			expectedErrCode:   codex.Internal,
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				agg := createUpdatePasswordByIdTestAccount(t, p.id, p.old, createdAt)
				repo.EXPECT().GetById(gomock.Any(), p.id).Return(agg, nil)
				return repo
			},
			setupCommand: func(t *testing.T, p params) *mockRepo.MockAccountCommand {
				repo := mockRepo.NewMockAccountCommand(ctrl)
				agg := createUpdatePasswordByIdTestAccount(t, p.id, p.old, createdAt)
				pass, err := vo.NewPassword(p.new)
				require.NoError(t, err)
				err = agg.ChangePassword(pass, eventTime)
				assert.NoError(t, err)
				repo.EXPECT().Update(gomock.Any(), agg).Return(errx.New(codex.Internal, "connect refuse"))
				return repo
			},
		},
		{
			name: "Should handle incorrect password update error",
			params: params{
				id:  uuid.New().String(),
				old: "Old_password4142@_incorrect",
				new: "New_password4142@",
			},
			logNotEmptyFields: []string{"id"},
			expectedErr:       anyErr,
			expectedErrCode:   codex.InvalidArgument,
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				agg := createUpdatePasswordByIdTestAccount(t, p.id, "Old_password4142@", createdAt)
				repo.EXPECT().GetById(gomock.Any(), p.id).Return(agg, nil)
				return repo
			},
			setupCommand: func(t *testing.T, p params) *mockRepo.MockAccountCommand {
				return nil
			},
		},
		{
			name: "Should handle incorrect new password update error",
			params: params{
				id:  uuid.New().String(),
				old: "Old_password4142@",
				new: "no",
			},
			logNotEmptyFields: []string{"id"},
			expectedErr:       anyErr,
			expectedErrCode:   codex.InvalidArgument,
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				agg := createUpdatePasswordByIdTestAccount(t, p.id, p.old, createdAt)
				repo.EXPECT().GetById(gomock.Any(), p.id).Return(agg, nil)
				return repo
			},
			setupCommand: func(t *testing.T, p params) *mockRepo.MockAccountCommand {
				return nil
			},
		},
		{
			name: "Should handle incorrect get by id",
			params: params{
				id:  uuid.New().String(),
				old: "Old_password4142@",
				new: "New_password4142@",
			},
			logNotEmptyFields: []string{"id"},
			expectedErr:       errx.New(codex.NotFound, "user not found"),
			expectedErrCode:   codex.NotFound,
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				repo.EXPECT().GetById(gomock.Any(), p.id).Return(nil, errx.New(codex.NotFound, "user not found"))
				return repo
			},
			setupCommand: func(t *testing.T, p params) *mockRepo.MockAccountCommand {
				return nil
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			command := tc.setupCommand(t, tc.params)
			query := tc.setupQuery(t, tc.params)
			timeDep := &timer{
				t: eventTime,
			}
			counter := &commandCounter{
				AccountCommand: command,
			}
			l, dump := setupLogger()

			useCase, err := NewUseCase(l, accFactory, query, counter, timeDep)
			assert.NoError(t, err)

			err := useCase.UpdatePasswordById(ctx, tc.params.id, tc.params.old, tc.params.new)
			if tc.expectedErr != nil {
				assert.Error(t, err)
				if !errors.Is(tc.expectedErr, anyErr) {
					assert.Equal(t, errors.Cause(tc.expectedErr).Error(), errors.Cause(err).Error())
				}
				c := errx.Code(err)
				assert.Equal(t, tc.expectedErrCode, c)
				assert.Equal(t, 1, len(dump.Dumps))
				logData := dump.Dumps[0]

				m := map[string]interface{}{}
				errU := json.Unmarshal(logData, &m)
				assert.NoError(t, errU)
				assert.NotEmpty(t, m["message"])
				assert.NotEmpty(t, m["error"])

				for _, f := range tc.logNotEmptyFields {
					assert.NotEmpty(t, m[f])
				}
			} else {
				assert.NoError(t, err)
				if len(dump.Dumps) > 0 {
					logData := dump.Dumps[0]
					m := map[string]interface{}{}
					errU := json.Unmarshal(logData, &m)
					assert.NoError(t, errU)
					assert.NotEmpty(t, m["message"])
					data, ok := m["aggregate"]
					assert.True(t, ok)
					t.Log(data)
					temp, ok := data.(map[string]interface{})
					assert.True(t, ok)
					tmp := temp["events"].([]interface{})
					assert.Equal(t, []interface{}{string(event.AccountChangePasswordType)}, tmp)
				}
			}
		})
	}
}
