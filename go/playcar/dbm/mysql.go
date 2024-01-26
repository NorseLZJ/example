package dbm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MysqlConn struct {
	*gorm.DB
}

func (a *MysqlConn) Close() {
	db, err := a.DB.DB()
	if err != nil {
		return
	}
	_ = db.Close()
}

func NewMysqlConn(dns string, logLevel logger.LogLevel) (*MysqlConn, *gorm.DB) {
	/*
		dns example
		dns := "lzj:123456@tcp(127.0.0.1:3306)/information_schema?charset=utf8mb4&parseTime=True&loc=Local"
	*/
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		panic(err)
	}
	return &MysqlConn{DB: db}, db
}
