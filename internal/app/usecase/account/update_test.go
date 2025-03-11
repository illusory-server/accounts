package account

import (
	"github.com/golang/mock/gomock"
	mock_repo "github.com/illusory-server/accounts/internal/mock/repo"
	"github.com/illusory-server/accounts/pkg/logger"
	"github.com/illusory-server/accounts/pkg/logger/log"
	"testing"
)

func setupLogger() (logger.Logger, *log.OutMultiDump) {
	out := log.NewOutMultiDump()
	l := log.NewLogger(&log.Options{
		Out:   out,
		Level: logger.DebugLvl,
	})
	return l, out
}

func TestUpdateById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	//l, _ := setupLogger()
	//ctx := context.TODO()
	//accFactory := factory.NewAccountFactory()

	type param struct {
	}

	testCases := []struct {
		name         string
		setupCommand func() *mock_repo.MockAccountCommand
		setupQuery   func() *mock_repo.MockAccountQuery
	}{
		{},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			//command := tc.setupCommand()
			//query := tc.setupQuery()

			//useCase := NewUseCase(l, accFactory, query, command)

			//useCase.UpdateById(ctx)
		})
	}
}
