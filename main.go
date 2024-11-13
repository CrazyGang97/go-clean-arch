package main

import (
	"github.com/CrazyGang97/go-clean-arch/internal/control"
	"github.com/CrazyGang97/go-clean-arch/internal/service"
)

func main() {
	service.NewCleanArchService(control.New())
}
