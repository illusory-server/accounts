package main

import (
	"fmt"
	"github.com/illusory-server/accounts/pkg/errors/codex"
	"github.com/illusory-server/accounts/pkg/errors/errx"
	"github.com/pkg/errors"
)

func repo(i int) error {
	if i == 10 {
		return errx.New(codex.NotFound, "user not found")
	}
	return nil
}

func repo2(i int) error {
	if i == 9 {
		return errx.New(codex.AlreadyExists, "user exists")
	}
	return nil
}

func repo3(i int) error {
	if i == 8 {
		return errx.New(codex.Internal, "internal server error")
	}
	return nil
}

func useCase(i int) error {
	err := repo(i)
	if err != nil {
		return errors.Wrap(err, "[useCase] repo")
	}

	err = repo2(i)
	if err != nil {
		return errors.Wrap(err, "[useCase] repo2")
	}

	err = repo3(i)
	if err != nil {
		return errors.Wrap(err, "[useCase] repo3")
	}

	return nil
}

func main() {
	err := useCase(7)
	if err != nil {
		c := errx.Code(err)
		fmt.Println("code:", c, "message:", err.Error())
		return
	}
	fmt.Println("complete")
}
