package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/lpernett/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
	Router *mux.Router
}


type AppConfig struct {
	AppName string
	AppEnv string
	AppPort string
}

type DBConfig struct {
	DBHost string
	DBUser string
	DBPass string
	DBName string
	DBPort string
}

func (server *Server) Initialize(appConfig AppConfig, dbConfig DBConfig) {
	fmt.Println("Welcome to "+ appConfig.AppName + " API server!")

	server.ConnectDB(dbConfig)
	server.initializeRoutes()
}


func (server *Server) Run(addr string) {
	fmt.Printf("Server is running on port %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func (server *Server) ConnectDB(dbConfig DBConfig) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbConfig.DBHost, dbConfig.DBUser, dbConfig.DBPass, dbConfig.DBName, dbConfig.DBPort)
	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
    panic(err)
  } else {
    fmt.Println("Connected to the database!")
  }

	for _, model := range RegisterModesl() {
		err = server.DB.Debug().AutoMigrate(model.Model)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Database migrated successfully!")
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Run() {
	var server = Server{}
	var appConfig = AppConfig{}
	var dbConfig = DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appConfig.AppName = getEnv("APP_NAME", "Go API")
	appConfig.AppEnv = getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT", "9000")

	dbConfig.DBHost = getEnv("DB_HOST", "localhost")
	dbConfig.DBUser = getEnv("DB_USER", "postgres")
	dbConfig.DBPass = getEnv("DB_PASS", "romi")
	dbConfig.DBName = getEnv("DB_NAME", "go_commerce")
	dbConfig.DBPort = getEnv("DB_PORT", "5432")

	server.Initialize(appConfig, dbConfig)
	server.Run(":" + appConfig.AppPort)
}