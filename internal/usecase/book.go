package usecase

import (
	"context"
	"github.com/CrazyGang97/go-clean-arch/internal/control"
	"github.com/CrazyGang97/go-clean-arch/internal/infra/db/model"
	"github.com/CrazyGang97/go-clean-arch/internal/repo"
)

type BookRepo interface {
	List(ctx context.Context, ops ...repo.BookOption) ([]*model.BookPO, error)
}

type BookDO struct {
	Name     string
	AuthorID int64
}

type Book struct {
	cc   *control.Control
	repo BookRepo
}

func NewBook(cc *control.Control) *Book {
	return &Book{
		cc:   cc,
		repo: repo.NewBookRepo(),
	}
}

func (r *Book) List(ctx context.Context, ops ...repo.BookOption) ([]*BookDO, error) {
	books, err := r.repo.List(ctx, ops...)
	if err != nil {
		return nil, err
	}
	return bookPO2DOBatch(books), nil
}

func bookPO2DO(book *model.BookPO) *BookDO {
	return &BookDO{
		Name:     book.Name,
		AuthorID: book.AuthorID,
	}
}

func bookPO2DOBatch(books []*model.BookPO) []*BookDO {
	res := make([]*BookDO, len(books))
	for i, book := range books {
		res[i] = bookPO2DO(book)
	}
	return res
}
