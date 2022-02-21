package main

import (
	"go_web_server/internal/handler"
	ws "go_web_server/internal/model"
	"go_web_server/internal/repository"
	"go_web_server/internal/repository/mysql"
	"go_web_server/internal/service"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := mysql.NewMySQLDB(&mysql.Config{
		Username: os.Getenv("username"),
		Password: os.Getenv("password"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		DBName:   viper.GetString("db.dbname"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db_query: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	err = services.CreateTables()
	if err != nil {
		log.Fatalf("failed db_query create table: %s", err.Error())
	}

	err = services.InsertInto()
	if err != nil {
		log.Fatalf("failed db_query insert into: %s", err.Error())
	}

	srv := new(ws.Server)
	if err := srv.Run("8080", handlers.InitHTTPRouter()); err != nil {
		log.Fatalf("error occurred while run http server:%s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
