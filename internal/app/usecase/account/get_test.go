package account

import (
	"context"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/illusory-server/accounts/internal/app/factory"
	mockRepo "github.com/illusory-server/accounts/internal/mock/repo"
	"github.com/illusory-server/accounts/pkg/errors/codex"
	"github.com/illusory-server/accounts/pkg/errors/errx"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestAccountsUseCase_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	type params struct {
		id string
	}

	createdAt := time.Now().Add(-time.Hour)
	eventTime := time.Now()

	testCases := []struct {
		name         string
		params       params
		expected     func(*testing.T, params) *WithoutPassword
		expectedErr  error
		expectedCode codex.Code
		logLevel     string
		setupQuery   func(*testing.T, params) *mockRepo.MockAccountQuery
	}{
		{
			name: "Should correctly get account by id",
			params: params{
				id: uuid.New().String(),
			},
			expected: func(t *testing.T, p params) *WithoutPassword {
				accFactory := factory.NewAccountFactory(&timer{createdAt}, genID{p.id})
				acc, err := accFactory.CreateAccount(
					"John",
					"Doe",
					"john.doe@gmail.com",
					"john",
					"@CorrectPassword123%",
				)
				assert.NoError(t, err)
				return ConvertAccountAggregateToWithoutPassword(acc)
			},
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				accFactory := factory.NewAccountFactory(&timer{createdAt}, genID{p.id})
				acc, err := accFactory.CreateAccount(
					"John",
					"Doe",
					"john.doe@gmail.com",
					"john",
					"@CorrectPassword123%",
				)
				assert.NoError(t, err)
				repo.EXPECT().GetById(gomock.Any(), p.id).Return(acc, nil)
				return repo
			},
		},
		{
			name: "Should error canceled correctly handle",
			params: params{
				id: uuid.New().String(),
			},
			expected: func(t *testing.T, p params) *WithoutPassword {
				return nil
			},
			expectedErr:  errx.New(codex.Canceled, "connected refused"),
			expectedCode: codex.Canceled,
			logLevel:     "error",
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				repo.EXPECT().GetById(gomock.Any(), p.id).Return(nil, errx.New(codex.Canceled, "connected refused"))
				return repo
			},
		},
		{
			name: "Should not error level error handle",
			params: params{
				id: uuid.New().String(),
			},
			expected: func(t *testing.T, p params) *WithoutPassword {
				return nil
			},
			expectedErr:  errx.New(codex.NotFound, "acc by id not found"),
			expectedCode: codex.NotFound,
			logLevel:     "info",
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				repo.EXPECT().GetById(gomock.Any(), p.id).Return(nil, errx.New(codex.NotFound, "acc by id not found"))
				return repo
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			query := tc.setupQuery(t, tc.params)
			timeDep := &timer{
				t: eventTime,
			}

			l, dump := setupLogger()

			useCase, err := NewUseCase(l, nil, query, nil, timeDep)
			assert.NoError(t, err)

			result, err := useCase.GetById(ctx, tc.params.id)
			if tc.expectedErr != nil {
				assert.Error(t, err)
				c := errx.Code(err)
				assert.Equal(t, tc.expectedCode, c)
				assert.Equal(t, 1, len(dump.Dumps))
				logData := dump.Dumps[0]

				m := map[string]interface{}{}
				errU := json.Unmarshal(logData, &m)
				assert.NoError(t, errU)
				assert.NotEmpty(t, m["message"])
				assert.NotEmpty(t, m["level"])
				lvl := m["level"].(string)
				assert.Equal(t, tc.logLevel, strings.ToLower(lvl))
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected(t, tc.params), result)
			}
		})
	}
}

func TestAccountsUseCase_GetByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	type params struct {
		email string
	}

	createdAt := time.Now().Add(-time.Hour)
	eventTime := time.Now()
	id := uuid.New().String()

	testCases := []struct {
		name         string
		params       params
		expected     func(*testing.T, params) *WithoutPassword
		expectedErr  error
		expectedCode codex.Code
		logLevel     string
		setupQuery   func(*testing.T, params) *mockRepo.MockAccountQuery
	}{
		{
			name: "Should correctly get account by email",
			params: params{
				email: "current@gmail.com",
			},
			expected: func(t *testing.T, p params) *WithoutPassword {
				accFactory := factory.NewAccountFactory(&timer{createdAt}, genID{id})
				acc, err := accFactory.CreateAccount(
					"John",
					"Doe",
					p.email,
					"john",
					"@CorrectPassword123%",
				)
				assert.NoError(t, err)
				return ConvertAccountAggregateToWithoutPassword(acc)
			},
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				accFactory := factory.NewAccountFactory(&timer{createdAt}, genID{id})
				acc, err := accFactory.CreateAccount(
					"John",
					"Doe",
					p.email,
					"john",
					"@CorrectPassword123%",
				)
				assert.NoError(t, err)
				repo.EXPECT().GetByEmail(gomock.Any(), p.email).Return(acc, nil)
				return repo
			},
		},
		{
			name: "Should error canceled correctly handle",
			params: params{
				email: "current@gmail.com",
			},
			expected: func(t *testing.T, p params) *WithoutPassword {
				return nil
			},
			expectedErr:  errx.New(codex.Canceled, "connected refused"),
			expectedCode: codex.Canceled,
			logLevel:     "error",
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				repo.EXPECT().GetByEmail(gomock.Any(), p.email).Return(nil, errx.New(codex.Canceled, "connected refused"))
				return repo
			},
		},
		{
			name: "Should not error level error handle",
			params: params{
				email: "current@gmail.com",
			},
			expected: func(t *testing.T, p params) *WithoutPassword {
				return nil
			},
			expectedErr:  errx.New(codex.NotFound, "acc by email not found"),
			expectedCode: codex.NotFound,
			logLevel:     "info",
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				repo.EXPECT().GetByEmail(gomock.Any(), p.email).Return(nil, errx.New(codex.NotFound, "acc by email not found"))
				return repo
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			query := tc.setupQuery(t, tc.params)
			timeDep := &timer{
				t: eventTime,
			}

			l, dump := setupLogger()

			useCase, err := NewUseCase(l, nil, query, nil, timeDep)
			assert.NoError(t, err)

			result, err := useCase.GetByEmail(ctx, tc.params.email)
			if tc.expectedErr != nil {
				assert.Error(t, err)
				c := errx.Code(err)
				assert.Equal(t, tc.expectedCode, c)
				assert.Equal(t, 1, len(dump.Dumps))
				logData := dump.Dumps[0]

				m := map[string]interface{}{}
				errU := json.Unmarshal(logData, &m)
				assert.NoError(t, errU)
				assert.NotEmpty(t, m["message"])
				assert.NotEmpty(t, m["level"])
				lvl := m["level"].(string)
				assert.Equal(t, tc.logLevel, strings.ToLower(lvl))
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected(t, tc.params), result)
			}
		})
	}
}

func TestAccountsUseCase_GetByNickname(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	type params struct {
		nickname string
	}

	createdAt := time.Now().Add(-time.Hour)
	eventTime := time.Now()
	id := uuid.New().String()

	testCases := []struct {
		name         string
		params       params
		expected     func(*testing.T, params) *WithoutPassword
		expectedErr  error
		expectedCode codex.Code
		logLevel     string
		setupQuery   func(*testing.T, params) *mockRepo.MockAccountQuery
	}{
		{
			name: "Should correctly get account by nickname",
			params: params{
				nickname: "john",
			},
			expected: func(t *testing.T, p params) *WithoutPassword {
				accFactory := factory.NewAccountFactory(&timer{createdAt}, genID{id})
				acc, err := accFactory.CreateAccount(
					"John",
					"Doe",
					"john.doe@gmail.com",
					p.nickname,
					"@CorrectPassword123%",
				)
				assert.NoError(t, err)
				return ConvertAccountAggregateToWithoutPassword(acc)
			},
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				accFactory := factory.NewAccountFactory(&timer{createdAt}, genID{id})
				acc, err := accFactory.CreateAccount(
					"John",
					"Doe",
					"john.doe@gmail.com",
					p.nickname,
					"@CorrectPassword123%",
				)
				assert.NoError(t, err)
				repo.EXPECT().GetByNickname(gomock.Any(), p.nickname).Return(acc, nil)
				return repo
			},
		},
		{
			name: "Should error canceled correctly handle",
			params: params{
				nickname: "john",
			},
			expected: func(t *testing.T, p params) *WithoutPassword {
				return nil
			},
			expectedErr:  errx.New(codex.Canceled, "connected refused"),
			expectedCode: codex.Canceled,
			logLevel:     "error",
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				repo.EXPECT().GetByNickname(gomock.Any(), p.nickname).Return(nil, errx.New(codex.Canceled, "connected refused"))
				return repo
			},
		},
		{
			name: "Should not error level error handle",
			params: params{
				nickname: "john",
			},
			expected: func(t *testing.T, p params) *WithoutPassword {
				return nil
			},
			expectedErr:  errx.New(codex.NotFound, "acc by email not found"),
			expectedCode: codex.NotFound,
			logLevel:     "info",
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				repo.EXPECT().GetByNickname(gomock.Any(), p.nickname).Return(nil, errx.New(codex.NotFound, "acc by email not found"))
				return repo
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			query := tc.setupQuery(t, tc.params)
			timeDep := &timer{
				t: eventTime,
			}

			l, dump := setupLogger()

			useCase, err := NewUseCase(l, nil, query, nil, timeDep)
			assert.NoError(t, err)

			result, err := useCase.GetByNickname(ctx, tc.params.nickname)
			if tc.expectedErr != nil {
				assert.Error(t, err)
				c := errx.Code(err)
				assert.Equal(t, tc.expectedCode, c)
				assert.Equal(t, 1, len(dump.Dumps))
				logData := dump.Dumps[0]

				m := map[string]interface{}{}
				errU := json.Unmarshal(logData, &m)
				assert.NoError(t, errU)
				assert.NotEmpty(t, m["message"])
				assert.NotEmpty(t, m["level"])
				lvl := m["level"].(string)
				assert.Equal(t, tc.logLevel, strings.ToLower(lvl))
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected(t, tc.params), result)
			}
		})
	}
}
