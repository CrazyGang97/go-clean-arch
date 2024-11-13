package control

import "github.com/CrazyGang97/go-clean-arch/internal/usecase"

type Control struct {
	User *usecase.User
	Book *usecase.Book
}

func New() *Control {
	cc := &Control{}
	user := usecase.NewUser(cc)
	book := usecase.NewBook(cc)
	cc.User = user
	cc.Book = book
	return cc
}
