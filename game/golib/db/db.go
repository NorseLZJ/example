package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
)

var (
	modelMap = make(map[string][]model)
)

type ConnDns interface {
	MySQLDns(dbs []string) []string
}

type model struct {
	mod  Mode
	init func(dbMap *gorp.DbMap)
}

type Mode interface {
	SetDbMap(dbMap *gorp.DbMap)
	DbMap() *gorp.DbMap
	SetDb(db *sql.DB)
	Db() *sql.DB
}

type baseModel struct {
	dbMap *gorp.DbMap
	db    *sql.DB
}

func (c *baseModel) SetDbMap(dbMap *gorp.DbMap) { c.dbMap = dbMap }
func (c *baseModel) DbMap() *gorp.DbMap         { return c.dbMap }
func (c *baseModel) SetDb(db *sql.DB)           { c.db = db }
func (c *baseModel) Db() *sql.DB                { return c.db }

type SqlLog struct{}

func (l *SqlLog) Printf(format string, v ...interface{}) {
	log.Println(fmt.Sprintf(format, v...))
}

func Reg(key string, mod Mode, init func(dbMap *gorp.DbMap)) {
	v, ok := modelMap[key]
	if ok {
		for _, it := range v {
			if mod == it.mod {
				return
			}
			modelMap[key] = append(v, model{
				mod:  mod,
				init: init,
			})
		}
	} else {
		v = make([]model, 0, 5)
		modelMap[key] = append(v, model{
			mod:  mod,
			init: init,
		})
	}
}

func Init(dnsF ConnDns) {
	for key := range modelMap {
		dns := dnsF.MySQLDns(key)
		if dns == "" {
			log.Fatal("dns is nil key:", key)
		}
		db, err := sql.Open("mysql", dns)
		if err != nil {
			log.Fatal(key)
		}
		if err = db.Ping(); err != nil {
			log.Fatal(err)
		}
		db.SetMaxIdleConns(16)
		db.SetMaxOpenConns(32)
		dbMap := &gorp.DbMap{
			Db:      db,
			Dialect: gorp.MySQLDialect{},
		}
		dbMap.TraceOn("[gorp]", &SqlLog{})
		mapIt, ok := modelMap[key]
		if ok {
			for _, it := range mapIt {
				it.mod.SetDbMap(dbMap)
				it.mod.SetDb(db)
				it.init(dbMap)
			}
		}
	}
	modelMap = make(map[string][]model)
}
