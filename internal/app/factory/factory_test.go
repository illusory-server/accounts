package factory

import (
	"github.com/google/uuid"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type (
	mockTimeNower struct {
		mock.Mock
	}

	mockIDGenerator struct {
		mock.Mock
	}
)

func (m *mockTimeNower) Now() time.Time {
	args := m.Called()
	return args.Get(0).(time.Time)
}

func (m *mockIDGenerator) GenerateID() string {
	args := m.Called()
	return args.String(0)
}

func TestAccountFactoryImpl_CreateAccount_Success(t *testing.T) {
	// Setup mocks
	now := time.Now()
	expectedID := uuid.New().String()

	timeNower := new(mockTimeNower)
	timeNower.On("Now").Return(now)

	idGenerator := new(mockIDGenerator)
	idGenerator.On("GenerateID").Return(expectedID, nil)

	// Create factory
	factory := NewAccountFactory(timeNower, idGenerator)

	// Test data
	firstName := "John"
	lastName := "Doe"
	email := "john.doe@example.com"
	nick := "johndoe"
	password := "securePassword123"

	// Execute
	account, err := factory.CreateAccount(firstName, lastName, email, nick, password)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, account)

	// Verify account values
	assert.Equal(t, expectedID, account.Account().ID().Value())
	assert.Equal(t, firstName, account.Account().Info().FirstName())
	assert.Equal(t, lastName, account.Account().Info().LastName())
	assert.Equal(t, email, account.Account().Info().Email())
	assert.Equal(t, nick, account.Account().Nickname())
	assert.Equal(t, password, account.Account().Password().Value())
	assert.Equal(t, now, account.Account().CreatedAt())
	assert.Equal(t, now, account.Account().UpdatedAt())

	// Verify mocks
	timeNower.AssertExpectations(t)
	idGenerator.AssertExpectations(t)
}

func TestAccountFactoryImpl_CreateAccount_InvalidEmail(t *testing.T) {
	// Setup mocks
	now := time.Now()
	expectedID := uuid.New().String()

	timeNower := new(mockTimeNower)
	timeNower.On("Now").Return(now)

	idGenerator := new(mockIDGenerator)
	idGenerator.On("GenerateID").Return(expectedID, nil)

	// Create factory
	factory := NewAccountFactory(timeNower, idGenerator)

	// Test data with invalid email
	firstName := "John"
	lastName := "Doe"
	email := "invalid-email"
	nick := "johndoe"
	password := "securePassword123"

	// Execute
	account, err := factory.CreateAccount(firstName, lastName, email, nick, password)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, account)
	assert.Contains(t, err.Error(), "vo.NewAccountInfo")

	// Verify mocks
	idGenerator.AssertExpectations(t)
}

func TestAccountFactoryImpl_CreateAccount_InvalidFirstName(t *testing.T) {
	// Setup mocks
	now := time.Now()
	expectedID := uuid.New().String()

	timeNower := new(mockTimeNower)
	timeNower.On("Now").Return(now)

	idGenerator := new(mockIDGenerator)
	idGenerator.On("GenerateID").Return(expectedID, nil)

	// Create factory
	factory := NewAccountFactory(timeNower, idGenerator)

	// Test data with invalid email
	firstName := "l"
	lastName := "Doe"
	email := "john.doe@example.com"
	nick := "johndoe"
	password := "securePassword123"

	// Execute
	account, err := factory.CreateAccount(firstName, lastName, email, nick, password)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, account)
	assert.Contains(t, err.Error(), "vo.NewAccountInfo")

	// Verify mocks
	idGenerator.AssertExpectations(t)
}

func TestAccountFactoryImpl_CreateAccount_InvalidLastName(t *testing.T) {
	// Setup mocks
	now := time.Now()
	expectedID := uuid.New().String()

	timeNower := new(mockTimeNower)
	timeNower.On("Now").Return(now)

	idGenerator := new(mockIDGenerator)
	idGenerator.On("GenerateID").Return(expectedID, nil)

	// Create factory
	factory := NewAccountFactory(timeNower, idGenerator)

	// Test data with invalid email
	firstName := "John"
	lastName := "D"
	email := "john.doe@example.com"
	nick := "johndoe"
	password := "securePassword123"

	// Execute
	account, err := factory.CreateAccount(firstName, lastName, email, nick, password)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, account)
	assert.Contains(t, err.Error(), "vo.NewAccountInfo")

	// Verify mocks
	idGenerator.AssertExpectations(t)
}

func TestAccountFactoryImpl_CreateAccount_InvalidPassword(t *testing.T) {
	// Setup mocks
	now := time.Now()
	expectedID := uuid.New().String()

	timeNower := new(mockTimeNower)
	timeNower.On("Now").Return(now)

	idGenerator := new(mockIDGenerator)
	idGenerator.On("GenerateID").Return(expectedID, nil)

	// Create factory
	factory := NewAccountFactory(timeNower, idGenerator)

	// Test data with invalid password (too short)
	firstName := "John"
	lastName := "Doe"
	email := "john.doe@example.com"
	nick := "johndoe"
	password := "short"

	// Execute
	account, err := factory.CreateAccount(firstName, lastName, email, nick, password)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, account)
	assert.Contains(t, err.Error(), "vo.NewPassword")

	// Verify mocks
	idGenerator.AssertExpectations(t)
}

func TestAccountFactoryImpl_CreateAccount_EmptyNick(t *testing.T) {
	// Setup mocks
	now := time.Now()
	expectedID := uuid.New().String()

	timeNower := new(mockTimeNower)
	timeNower.On("Now").Return(now)

	idGenerator := new(mockIDGenerator)
	idGenerator.On("GenerateID").Return(expectedID, nil)

	// Create factory
	factory := NewAccountFactory(timeNower, idGenerator)

	// Test data with empty nick
	firstName := "John"
	lastName := "Doe"
	email := "john.doe@example.com"
	nick := ""
	password := "securePassword123"

	// Execute
	account, err := factory.CreateAccount(firstName, lastName, email, nick, password)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, account)
	assert.Contains(t, err.Error(), "entity.NewAccount")

	// Verify mocks
	timeNower.AssertExpectations(t)
	idGenerator.AssertExpectations(t)
}

func TestAccountFactoryImpl_CreateAccount_IDGenerationError(t *testing.T) {
	// Setup mocks with ID generator error
	timeNower := new(mockTimeNower)
	idGenerator := new(mockIDGenerator)
	idGenerator.On("GenerateID").Return("bad-id")

	// Create factory
	factory := NewAccountFactory(timeNower, idGenerator)

	// Test data
	firstName := "John"
	lastName := "Doe"
	email := "john.doe@example.com"
	nick := "johndoe"
	password := "securePassword123"

	// Execute
	account, err := factory.CreateAccount(firstName, lastName, email, nick, password)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, account)
	assert.Contains(t, err.Error(), "vo.NewID")

	// Verify mocks
	idGenerator.AssertExpectations(t)
}
