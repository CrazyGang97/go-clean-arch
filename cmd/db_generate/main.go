package main

import (
	"github.com/CrazyGang97/go-clean-arch/internal/infra/db"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../../internal/infra/db",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	db.Init()
	g.UseDB(db.DB())
}
