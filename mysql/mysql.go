package mysql

import (
	"db_lab_library/mysql/query"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"sync"
)

var (
	DB       *gorm.DB
	once     sync.Once
	user     = os.Getenv("SSE_MYSQL_USER")
	password = os.Getenv("SSE_MYSQL_PASSWORD")
	host     = os.Getenv("SSE_MYSQL_HOST")
	database = os.Getenv("SSE_MYSQL_DB")
	dsnTpl   = "%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
)

func Init() {
	once.Do(func() {
		var err error
		DB, err = gorm.Open(mysql.Open(fmt.Sprintf(dsnTpl, user, password, host, database)))
		if err != nil {
			panic(err)
		}
	})
	query.SetDefault(DB)
}
