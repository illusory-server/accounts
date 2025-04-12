package job

import (
	"github.com/illusory-server/accounts/pkg/ecosystem"
	"time"
)

type HttpJob struct{}

type HttpJobParam struct {
	Address        string
	RequestTimeout time.Duration
}

func NewHttpJob[T any](
	param HttpJobParam,
) (*ecosystem.HttpJob[T], error) {
	//httpJoh := ecosystem.NewHttpJobBuilder().
	//	Address(param.Address).
	//	RequestTimeout(param.RequestTimeout).
	return nil, nil
}
