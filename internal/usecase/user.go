package usecase

import (
	"context"
	"github.com/CrazyGang97/go-clean-arch/internal/control"
	"github.com/CrazyGang97/go-clean-arch/internal/infra/db/model"
	"github.com/CrazyGang97/go-clean-arch/internal/repo"
)

type UserRepo interface {
	Create(ctx context.Context, user *model.UserPO) error
	GetByName(ctx context.Context, name string) (*model.UserPO, error)
}

type UserDO struct {
	Name  string
	Age   int32
	Books []*BookDO
}

type User struct {
	cc   *control.Control
	repo UserRepo
}

func NewUser(cc *control.Control) *User {
	return &User{
		cc:   cc,
		repo: repo.NewUserRepo(),
	}
}

func (r *User) Create(ctx context.Context, user *UserDO) error {
	// create user
	return r.repo.Create(ctx, &model.UserPO{
		Name: user.Name,
		Age:  user.Age,
	})
}

func (r *User) Get(ctx context.Context, name string) (*UserDO, error) {
	// get user first
	user, err := r.repo.GetByName(ctx, name)
	if err != nil {
		// handle error
		return nil, err
	}
	books, err := r.cc.Book.List(ctx, repo.WithAuthorID(user.ID))
	if err != nil {
		return nil, err
	}
	return buildUserDO(user, books), nil
}

func buildUserDO(user *model.UserPO, books []*BookDO) *UserDO {
	return &UserDO{
		Name:  user.Name,
		Age:   user.Age,
		Books: books,
	}
}
