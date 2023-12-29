package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"os"
)

var (
	user     = os.Getenv("SSE_MYSQL_USER")
	password = os.Getenv("SSE_MYSQL_PASSWORD")
	host     = os.Getenv("SSE_MYSQL_HOST")
	database = os.Getenv("SSE_MYSQL_DB")
	dsnTpl   = "%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "mysql/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	gormdb, _ := gorm.Open(mysql.Open(fmt.Sprintf(dsnTpl, user, password, host, database)))
	g.UseDB(gormdb)

	g.ApplyBasic(
		g.GenerateModel("user"),
		g.GenerateModel("book"),
		g.GenerateModel("order"),
		g.GenerateModel("publisher"),
	)

	g.Execute()
}
