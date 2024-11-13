package main

import (
	"github.com/CrazyGang97/go-clean-arch/internal/infra/db"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "../../internal/infra/db",
		Mode:         gen.WithDefaultQuery, // generate mode
		ModelPkgPath: "../../internal/infra/db/model",
	})

	db.Init()
	g.UseDB(db.DB())
	g.ApplyBasic(
		g.GenerateModelAs("user", "UserPO"),
		g.GenerateModelAs("book", "BookPO"),
	)
	g.Execute()
}
