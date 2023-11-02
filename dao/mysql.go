package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn = "gosql:11223344@tcp(123.56.17.185:3306)/bubble?charset=utf8mb4&parseTime=True"

var (
	DB *gorm.DB
)

// InitDB MySQL数据库配置
func InitDB() (err error) {
	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		return fmt.Errorf("connect db fail: %w", err)
	}
	return DB.Error
}
