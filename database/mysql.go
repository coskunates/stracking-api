package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

const (
	mysqlHost     = "mysql_host"
	mysqlUserName = "mysql_username"
	mySqlPassword = "mysql_password"
	mySqlSchema   = "mysql_schema"
)

var (
	DB       *gorm.DB
	host     = "127.0.0.1" //os.Getenv(mysqlHost)
	username = "root"      //os.Getenv(mysqlUserName)
	password = "root"      //os.Getenv(mySqlPassword)
	schema   = "stock"     //os.Getenv(mySqlSchema)
)

func init() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,          // Disable color
		},
	)

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{
		Logger:                 newLogger,
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		fmt.Println(err)
	}

	DB = db
}

func GetClient() *gorm.DB {
	return DB
}

func GetMockClient() *gorm.DB {
	DB.DryRun = true
	return DB
}
