package app

import (
	"go_web_server/internal/repository/store"
	ws "go_web_server/internal/web"
	"go_web_server/internal/web/handler"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func Run() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := store.NewDB(&store.ConfigDB{
		DB:       "mysql",
		Username: os.Getenv("username"),
		Password: os.Getenv("password"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		DBName:   viper.GetString("db.dbname"),
	})

	if err != nil {
		log.Fatalf("failed to initialize db_query: %s", err.Error())
	}

	repos := store.NewRepository(db)
	handlers := handler.NewHandler(repos)

	err = repos.CreateTables()
	if err != nil {
		log.Fatalf("failed db_query create table: %s", err.Error())
	}

	err = repos.InsertInto()
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
