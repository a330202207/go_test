package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"go_blog/pkg/setting"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "无法获取 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("DB_NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	connection = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)

	db, err = gorm.Open(dbType, connection)

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}
