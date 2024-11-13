package repo

import (
	"context"
	"github.com/CrazyGang97/go-clean-arch/internal/infra/db"
	"github.com/CrazyGang97/go-clean-arch/internal/infra/db/model"
)

type bookRepo struct {
}

func NewBookRepo() *bookRepo {
	return &bookRepo{}
}

type BookOptions struct {
	AuthorID *int64
}

type BookOption func(*BookOptions)

func WithAuthorID(authorID int64) BookOption {
	return func(o *BookOptions) {
		o.AuthorID = &authorID
	}
}

func (r *bookRepo) List(ctx context.Context, ops ...BookOption) ([]*model.BookPO, error) {
	option := &BookOptions{}
	for _, op := range ops {
		op(option)
	}
	q := db.BookPO.WithContext(ctx)

	if option.AuthorID != nil {
		q = q.Where(db.BookPO.ID.Eq(*option.AuthorID))
	}

	books, err := q.Find()
	if err != nil {
		return nil, err
	}
	return books, nil
}
