package db

import (
	"fmt"
	"log"
	"testing"
	"time"

	"gopkg.in/gorp.v1"
)

var mc = &MySQLConfig{
	User:    "root",
	Pass:    "123456",
	Net:     "tcp",
	Host:    "127.0.0.1",
	Port:    3306,
	DbName:  "game",
	MaxConn: 10,
	MaxIdle: 10,
}

type User struct {
	Id         int       `db:"id"`
	NickName   string    `db:"nickName"`
	CreateTime time.Time `db:"createTime"`
	Sex        int       `db:"sex"`
}

const (
	dbName = "game"
)

type UserModel struct {
	CommonModel
}

var (
	userModel  = &UserModel{}
	userFieldS = objToString(User{})
)

func init() {
	Reg(dbName, userModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(User{}, "user").SetKeys(false, "id")
	})
}

func getUserModel() *UserModel {
	return userModel
}

func TestInit(t *testing.T) {
	Init(mc)

	sql := "SELECT `id`,`nickName`,`createTime` FROM `user`;"
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		u := &User{}
		err = rows.Scan(u.Id, u.NickName, u.CreateTime)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(u.Id, " - ", u.NickName, " - ", u.Sex, " - ", u.CreateTime)
	}
}
