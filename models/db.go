package models

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"

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
}

func gormConnect() *gorm.DB {

	dbUser := os.Getenv("DB_USERNAME")
	dbPwd := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME_negi")

	var dbURI string
	fmt.Println(gin.Mode())
	if gin.Mode() == gin.ReleaseMode {
		// socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
		// if !isSet {
		// 	socketDir = "/cloudsql"
		// }
		instanceConnectionName := os.Getenv("instanceConnectionName")
		dbURI = fmt.Sprintf("%s:%s@cloudsql(%s)/%s", dbUser, dbPwd, instanceConnectionName, dbName)

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

func GetFirebaseAuthClient(c *gin.Context) *auth.Client {
	app, err := firebase.NewApp(c.Request.Context(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	authClient, err := app.Auth(c.Request.Context())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	return authClient
}

//Setup database ctx and client
func Setup() (ctx context.Context, client *firestore.Client) {
	// [START fs_initialize]
	// Sets your Google Cloud Platform project ID.
	projectID := os.Getenv("gcloud_projectID")

	// Get a Firestore client.
	ctx = context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Printf("Failed to create client: %v", err)
	}

	// Close client when done.
	// defer client.Close()
	// [END fs_initialize]
	return
}
