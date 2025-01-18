package app

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Romiz-Lab/BE-go-ecommerce/database/seeders"
	"github.com/gorilla/mux"
	"github.com/lpernett/godotenv"
	"github.com/urfave/cli/v2"
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

}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func (server *Server) dbMigrate() {
	for _, model := range RegisterModesl() {
		err := server.DB.Debug().AutoMigrate(model.Model)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Database migrated successfully!")
}

func (server *Server) initCommands(dbConfig DBConfig) {
	server.ConnectDB(dbConfig)

	cmdApp := &cli.App{
		Name: "Go Ecommerce",
		Usage: "CLI for Go Ecommerce",
		Commands: []*cli.Command{
			{
				Name: "db:migrate",
				Usage: "Migrate database",
				Action: func(c *cli.Context) error {
					server.dbMigrate()
					return nil
				},
			},
			{
				Name: "db:seed",
				Usage: "Seed the database with sample data",
				Action: func(c *cli.Context) error {
					err := seeders.DBSeed(server.DB)
					if err != nil {
						log.Fatal(err)
					}
					return nil
				},
			},
		},
	}

	if err := cmdApp.Run(os.Args); err != nil {
		log.Fatal(err)
	}
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

	flag.Parse()
	arg := flag.Arg(0)
	if arg != "" {
		server.initCommands(appConfig, dbConfig)
	} else {
		server.Initialize(appConfig, dbConfig)
		server.Run(":" + appConfig.AppPort)
	}
}