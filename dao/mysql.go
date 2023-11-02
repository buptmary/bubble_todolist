package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn = "user:password@tcp(ip)/bubble?charset=utf8mb4&parseTime=True"

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
