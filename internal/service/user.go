package service

import (
	"context"
)

type UserDTO struct {
	Name     string
	Age      int32
	BookList []BookDTO
}

func (r *CleanArchService) GetUser(ctx context.Context) error {
	return nil
}
