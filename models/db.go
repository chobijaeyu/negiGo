package models

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// func init() {
// 	var (
// 		err                                                     error
// 		dbType, dbName, user, password, host, tablePrefix, port string
// 	)
// 	dbType = os.Getenv("DB_TYPE")
// 	dbName = os.Getenv("DB_NAME")
// 	user = os.Getenv("DB_USER")
// 	password = os.Getenv("DB_PASSWORD")
// 	host = os.Getenv("DB_HOST")
// 	port = os.Getenv("DB_PORT")
// 	tablePrefix = os.Getenv("DB_TABLE_PREFIX")

// 	// db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
// 	db, err = gorm.Open(dbType, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",z
// 		host,
// 		port,
// 		user,
// 		dbName,
// 		password,
// 	))

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableNameHandler string) string {
// 		return tablePrefix + defaultTableNameHandler
// 	}

// 	db.SingularTable(true)
// 	db.DB().SetMaxIdleConns(10)
// 	db.DB().SetMaxOpenConns(100)
// }

func init() {
	db = gormConnect()
	db.AutoMigrate(NegiField{})
}

func gormConnect() *gorm.DB {

	dbUser := os.Getenv("DB_USERNAME")
	dbPwd := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME_negi")

	var dbURI string
	fmt.Println(gin.Mode())
	if gin.Mode() == gin.ReleaseMode {
		socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
		if !isSet {
			socketDir = "/cloudsql"
		}
		instanceConnectionName := os.Getenv("instanceConnectionName")
		dbURI = fmt.Sprintf("%s:%s@unix(/%s/%s)/%s", dbUser, dbPwd, socketDir, instanceConnectionName, dbName)

	} else {
		PROTOCOL := fmt.Sprintf("tcp(%s:%s)", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
		dbURI = dbUser + ":" + dbPwd + "@" + PROTOCOL + "/" + dbName
	}

	dsn := dbURI + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}
