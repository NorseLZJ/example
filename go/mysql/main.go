package main

import (
	"database/sql"
	"fmt"

	"github.com/davyxu/golog"
	_ "github.com/go-sql-driver/mysql"
)

const (
	// connStr user:root  pw:123456  addr:127.0.0.1:3306
	connStr = "%s:%s@tcp(%s)/?charset=utf8"
)

var (
	log = golog.New("webserver")
)

func main() {

	dns := fmt.Sprintf(connStr, "root", "123456", "127.0.0.1:3307")
	log.Infof("dns:%s\n", dns)

	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Errorf("Open mysql database error: %s\n", err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Errorln(err)
		return
	}

	var (
		id, playerId, job, level string
	)
	rows, err := db.Query("SELECT ID,PlayerID,Job,Level FROM data.collect")
	if err != nil {
		log.Errorln(err)
	}
	count := 0
	for rows.Next() {
		rows.Scan(&id, &playerId, &job, &level)
		//log.Infof("id:%s,playerId:%s,job:%s,level:%s\n", id, playerId, job, level)
		count++
	}
	log.Infof("count:%d\n", count)
}
