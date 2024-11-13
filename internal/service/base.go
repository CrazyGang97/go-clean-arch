package service

import (
	"github.com/CrazyGang97/go-clean-arch/internal/control"
)

type CleanArchService struct {
	cc *control.Control
}

func NewCleanArchService(cc *control.Control) *CleanArchService {
	return &CleanArchService{cc: cc}
}
