package repo

import (
	"context"
	"github.com/CrazyGang97/go-clean-arch/internal/infra/db"
	"github.com/CrazyGang97/go-clean-arch/internal/infra/db/model"
)

type userRepo struct {
}

func NewUserRepo() *userRepo {
	return &userRepo{}
}

func (r *userRepo) Create(ctx context.Context, user *model.UserPO) error {
	return db.UserPO.WithContext(ctx).Create(user)
}

func (r *userRepo) GetByName(ctx context.Context, name string) (*model.UserPO, error) {
	return db.UserPO.WithContext(ctx).Where(db.UserPO.Name.Eq(name)).First()
}
