package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "mysql/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	gormdb, _ := gorm.Open(mysql.Open("root:taomingyu211@(127.0.0.1:3306)/bookstore_management_system?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(gormdb)

	g.ApplyBasic(
		g.GenerateModel("user"),
		g.GenerateModel("book"),
		g.GenerateModel("order"),
		g.GenerateModel("publisher"),
	)

	g.Execute()
}
