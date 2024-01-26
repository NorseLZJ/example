package dbm

import (
	"gorm.io/gorm/logger"
	"log"
	"strings"
)

var (
	_goDb   *MysqlConn
	_javaDb *MysqlConn
)

func Init(mysql []string) {
	for _, v := range mysql {
		ss := strings.Split(v, "|")
		if len(ss) != 2 {
			log.Fatalf("mysql config has err: %s", v)
		}
		switch ss[0] {
		case "GO":
			_goDb, _ = NewMysqlConn(ss[1], logger.Error)
		case "JAVA":
			_javaDb, _ = NewMysqlConn(ss[1], logger.Error)
		}
	}
	log.Println("mysql connecting success")
}

func GoDb() *MysqlConn {
	return _goDb
}

func JavaDb() *MysqlConn {
	return _javaDb
}
