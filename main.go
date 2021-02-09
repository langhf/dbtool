package main

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"os"
)

type DbBase struct {
	statement string
	dbName    string
	tbName    string
	dbNum     int
	tbNum     int
	dbCoon    sql.DB
}

func newDbBase() *DbBase {
	return &DbBase{
		statement: "",
		dbName:    "",
		tbName:    "",
		dbNum:     0,
		tbNum:     0,
	}
}

var Config = viper.New()

func parseSQL(dbFile string, db *DbBase) error {
	log.Info("parseSQL")
	if _, err := os.Stat(dbFile); err != nil {
		dbFile = "./db.sql"
	}
	if statement, err := ioutil.ReadFile(dbFile); err != nil {
		log.Info(statement)
		db.statement = string(statement)
		return nil
	}
	return io.EOF
}

func main() {
	dbBase := newDbBase()
	parseSQL("db.sql", dbBase)
}
