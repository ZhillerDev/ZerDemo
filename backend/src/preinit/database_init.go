package preinit

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB
var err error

func DatabaseInit() {
	// 设置数据库链接
	dsn := "root:123456@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"

	// 这是默认的链接数据库的代码，err进行可否链接判断
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("数据已成功链接！")

	// 数据库属性值配置
	var db *sql.DB
	db, _ = DB.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)
	db.SetConnMaxIdleTime(time.Hour)
	db.SetConnMaxLifetime(time.Hour)
}
