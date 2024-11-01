package biz

import (
	"context"
	"github.com/CrazyGang97/go-clean-arch/internal/repo"
)

type UserRepo interface {
	Create(ctx context.Context, user *repo.UserPO) error
	Get(ctx context.Context, id string) (*repo.UserPO, error)
}

type User struct {
	Name string
	Age  int64
}
