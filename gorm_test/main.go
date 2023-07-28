package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type Users struct {
	ID       uint   `gorm:"primaryKey;column:id"`
	Username string `gorm:"column:username"`
	Email    string `gorm:"column:email"`
}

func main() {
	db := getDB()
	//var count int64
	//var users []Users

	users := make([]*Users, 0)

	//aaa := &Users{}

	//if err := db.Find(&users); err != nil {
	//	fmt.Println(err)
	//}

	WithPage(db, nil, users)

}

func getDB() *gorm.DB {
	dsn := "root:123456@tcp(localhost:3307)/mydatabase?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // 使用標準輸出作為記錄器
			logger.Config{
				LogLevel: logger.Info, // 設置日誌級別，這裡設置為 Info 可印出 SQL 語句
			},
		),
	})
	if err != nil {
		panic("failed to connect database")
	}
	//db.Callback().Query().Before("gorm:query").Register("debug", func(db *gorm.DB) {
	//	sql := db.Statement.SQL.String()
	//	fmt.Println("sql: " + sql)
	//})

	return db
}
