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
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccountsUseCase_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.TODO()

	createdAt := time.Now()
	accFactory := factory.NewAccountFactory(&timer{
		t: createdAt,
	}, genID{
		id: uuid.New().String(),
	})
	eventTime := time.Now().Add(time.Millisecond)

	anyErr := errors.New("any error")

	type params struct {
		firstName, lastName, email, nick, password string
	}

	testCases := []struct {
		name                string
		params              params
		expectedErr         error
		expectedCode        codex.Code
		eventTime           time.Time
		logNotEmptyFields   []string
		expected            func(*testing.T, params) *WithoutPassword
		setupCommand        func(*testing.T, params) *mockRepo.MockAccountCommand
		setupQuery          func(*testing.T, params) *mockRepo.MockAccountQuery
		setupAccountFactory func(*testing.T, params) factory.AccountFactory
	}{
		{
			name: "Should create account successfully",
			params: params{
				firstName: "John",
				lastName:  "Smith",
				email:     "john.smith@gmail.com",
				nick:      "john",
				password:  "#CorrectPass124$$",
			},
			eventTime: eventTime,
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				repo.EXPECT().HasByNickname(gomock.Any(), p.nick).Return(false, nil)
				repo.EXPECT().HasByEmail(gomock.Any(), p.email).Return(false, nil)
				return repo
			},
			setupCommand: func(t *testing.T, p params) *mockRepo.MockAccountCommand {
				repo := mockRepo.NewMockAccountCommand(ctrl)
				acc, err := accFactory.CreateAccount(p.firstName, p.lastName, p.email, p.nick, p.password)
				assert.NoError(t, err)
				repo.EXPECT().Create(gomock.Any(), acc).Return(acc, nil)
				return repo
			},
			expected: func(t *testing.T, p params) *WithoutPassword {
				acc, err := accFactory.CreateAccount(p.firstName, p.lastName, p.email, p.nick, p.password)
				assert.NoError(t, err)
				return ConvertAccountAggregateToWithoutPassword(acc)
			},
		},
		{
			name: "Should create account error handle",
			params: params{
				firstName: "John",
				lastName:  "Smith",
				email:     "john.smith@gmail.com",
				nick:      "john",
				password:  "#CorrectPass124$$",
			},
			eventTime:         eventTime,
			expectedErr:       errx.New(codex.AlreadyExists, "account already exists"),
			expectedCode:      codex.AlreadyExists,
			logNotEmptyFields: []string{"account"},
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				repo.EXPECT().HasByNickname(gomock.Any(), p.nick).Return(false, nil)
				repo.EXPECT().HasByEmail(gomock.Any(), p.email).Return(false, nil)
				return repo
			},
			setupCommand: func(t *testing.T, p params) *mockRepo.MockAccountCommand {
				repo := mockRepo.NewMockAccountCommand(ctrl)
				acc, err := accFactory.CreateAccount(p.firstName, p.lastName, p.email, p.nick, p.password)
				assert.NoError(t, err)
				repo.EXPECT().Create(gomock.Any(), acc).Return(nil, errx.New(codex.AlreadyExists, "account already exists"))
				return repo
			},
			expected: func(t *testing.T, p params) *WithoutPassword {
				return nil
			},
		},
		{
			name: "Should query has by nick error handle",
			params: params{
				firstName: "John",
				lastName:  "Smith",
				email:     "john.smith@gmail.com",
				nick:      "john",
				password:  "#CorrectPass124$$",
			},
			eventTime:         eventTime,
			expectedErr:       errx.New(codex.Canceled, "connect refused"),
			expectedCode:      codex.Canceled,
			logNotEmptyFields: []string{"nickname"},
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				repo.EXPECT().HasByNickname(gomock.Any(), p.nick).Return(false, errx.New(codex.Canceled, "connect refused"))
				return repo
			},
			setupCommand: func(t *testing.T, p params) *mockRepo.MockAccountCommand {
				return nil
			},
			expected: func(t *testing.T, p params) *WithoutPassword {
				return nil
			},
		},
		{
			name: "Should query has by email error handle",
			params: params{
				firstName: "John",
				lastName:  "Smith",
				email:     "john.smith@gmail.com",
				nick:      "john",
				password:  "#CorrectPass124$$",
			},
			eventTime:         eventTime,
			expectedErr:       errx.New(codex.Canceled, "connect refused"),
			expectedCode:      codex.Canceled,
			logNotEmptyFields: []string{"email"},
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				repo.EXPECT().HasByNickname(gomock.Any(), p.nick).Return(false, nil)
				repo.EXPECT().HasByEmail(gomock.Any(), p.email).Return(false, errx.New(codex.Canceled, "connect refused"))
				return repo
			},
			setupCommand: func(t *testing.T, p params) *mockRepo.MockAccountCommand {
				return nil
			},
			expected: func(t *testing.T, p params) *WithoutPassword {
				return nil
			},
		},
		{
			name: "Should query has by nick candidate error handle",
			params: params{
				firstName: "John",
				lastName:  "Smith",
				email:     "john.smith@gmail.com",
				nick:      "john",
				password:  "#CorrectPass124$$",
			},
			eventTime:         eventTime,
			expectedErr:       ErrNicknameExists,
			expectedCode:      errx.Code(ErrNicknameExists),
			logNotEmptyFields: []string{"nickname"},
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				repo.EXPECT().HasByNickname(gomock.Any(), p.nick).Return(true, nil)
				return repo
			},
			setupCommand: func(t *testing.T, p params) *mockRepo.MockAccountCommand {
				return nil
			},
			expected: func(t *testing.T, p params) *WithoutPassword {
				return nil
			},
		},
		{
			name: "Should query has by nick candidate error handle",
			params: params{
				firstName: "John",
				lastName:  "Smith",
				email:     "john.smith@gmail.com",
				nick:      "john",
				password:  "#CorrectPass124$$",
			},
			eventTime:         eventTime,
			expectedErr:       ErrEmailExists,
			expectedCode:      errx.Code(ErrEmailExists),
			logNotEmptyFields: []string{"email"},
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				repo.EXPECT().HasByNickname(gomock.Any(), p.nick).Return(false, nil)
				repo.EXPECT().HasByEmail(gomock.Any(), p.email).Return(true, nil)
				return repo
			},
			setupCommand: func(t *testing.T, p params) *mockRepo.MockAccountCommand {
				return nil
			},
			expected: func(t *testing.T, p params) *WithoutPassword {
				return nil
			},
		},
		{
			name: "Should incorrect first name error handle",
			params: params{
				firstName: "P",
				lastName:  "Smith",
				email:     "john.smith@gmail.com",
				nick:      "john",
				password:  "#CorrectPass124$$",
			},
			eventTime:         eventTime,
			expectedErr:       anyErr,
			expectedCode:      codex.InvalidArgument,
			logNotEmptyFields: []string{"first_name", "last_name", "email", "nickname"},
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				repo.EXPECT().HasByNickname(gomock.Any(), p.nick).Return(false, nil)
				repo.EXPECT().HasByEmail(gomock.Any(), p.email).Return(false, nil)
				return repo
			},
			setupCommand: func(t *testing.T, p params) *mockRepo.MockAccountCommand {
				return nil
			},
			expected: func(t *testing.T, p params) *WithoutPassword {
				return nil
			},
		},
		{
			name: "Should incorrect last name error handle",
			params: params{
				firstName: "John",
				lastName:  "S",
				email:     "john.smith@gmail.com",
				nick:      "john",
				password:  "#CorrectPass124$$",
			},
			eventTime:         eventTime,
			expectedErr:       anyErr,
			expectedCode:      codex.InvalidArgument,
			logNotEmptyFields: []string{"first_name", "last_name", "email", "nickname"},
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				repo.EXPECT().HasByNickname(gomock.Any(), p.nick).Return(false, nil)
				repo.EXPECT().HasByEmail(gomock.Any(), p.email).Return(false, nil)
				return repo
			},
			setupCommand: func(t *testing.T, p params) *mockRepo.MockAccountCommand {
				return nil
			},
			expected: func(t *testing.T, p params) *WithoutPassword {
				return nil
			},
		},
		{
			name: "Should incorrect email error handle",
			params: params{
				firstName: "John",
				lastName:  "Smith",
				email:     "incorrect",
				nick:      "john",
				password:  "#CorrectPass124$$",
			},
			eventTime:         eventTime,
			expectedErr:       anyErr,
			expectedCode:      codex.InvalidArgument,
			logNotEmptyFields: []string{"first_name", "last_name", "email", "nickname"},
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				repo.EXPECT().HasByNickname(gomock.Any(), p.nick).Return(false, nil)
				repo.EXPECT().HasByEmail(gomock.Any(), p.email).Return(false, nil)
				return repo
			},
			setupCommand: func(t *testing.T, p params) *mockRepo.MockAccountCommand {
				return nil
			},
			expected: func(t *testing.T, p params) *WithoutPassword {
				return nil
			},
		},
		{
			name: "Should incorrect nickname error handle",
			params: params{
				firstName: "John",
				lastName:  "Smith",
				email:     "john.smith@gmail.com",
				nick:      "j",
				password:  "#CorrectPass124$$",
			},
			eventTime:         eventTime,
			expectedErr:       anyErr,
			expectedCode:      codex.InvalidArgument,
			logNotEmptyFields: []string{"first_name", "last_name", "email", "nickname"},
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				repo.EXPECT().HasByNickname(gomock.Any(), p.nick).Return(false, nil)
				repo.EXPECT().HasByEmail(gomock.Any(), p.email).Return(false, nil)
				return repo
			},
			setupCommand: func(t *testing.T, p params) *mockRepo.MockAccountCommand {
				return nil
			},
			expected: func(t *testing.T, p params) *WithoutPassword {
				return nil
			},
		},
		{
			name: "Should incorrect password error handle",
			params: params{
				firstName: "John",
				lastName:  "Smith",
				email:     "john.smith@gmail.com",
				nick:      "john",
				password:  "no",
			},
			eventTime:         eventTime,
			expectedErr:       anyErr,
			expectedCode:      codex.InvalidArgument,
			logNotEmptyFields: []string{"first_name", "last_name", "email", "nickname"},
			setupQuery: func(t *testing.T, p params) *mockRepo.MockAccountQuery {
				repo := mockRepo.NewMockAccountQuery(ctrl)
				repo.EXPECT().HasByNickname(gomock.Any(), p.nick).Return(false, nil)
				repo.EXPECT().HasByEmail(gomock.Any(), p.email).Return(false, nil)
				return repo
			},
			setupCommand: func(t *testing.T, p params) *mockRepo.MockAccountCommand {
				return nil
			},
			expected: func(t *testing.T, p params) *WithoutPassword {
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
			useCase := NewUseCase(l, accFactory, query, counter, timeDep)

			result, err := useCase.Create(
				ctx, tc.params.firstName, tc.params.lastName,
				tc.params.email, tc.params.nick, tc.params.password,
			)
			if tc.expectedErr != nil {
				assert.Error(t, err)
				if !errors.Is(tc.expectedErr, anyErr) {
					assert.Equal(t, errors.Cause(tc.expectedErr).Error(), errors.Cause(err).Error())
				}
				c := errx.Code(err)
				assert.Equal(t, tc.expectedCode, c)
				assert.Equal(t, 1, len(dump.Dumps))
				logData := dump.Dumps[0]

				m := map[string]interface{}{}
				errU := json.Unmarshal(logData, &m)
				assert.NoError(t, errU)
				assert.NotEmpty(t, m["message"])

				for _, f := range tc.logNotEmptyFields {
					assert.NotEmpty(t, m[f])
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected(t, tc.params), result)
			}
		})
	}
}
