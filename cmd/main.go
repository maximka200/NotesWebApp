package main

import (
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	todo "github.com/maximka200/NotesWebApp"
	handler "github.com/maximka200/NotesWebApp/pkg/handler"
	"github.com/maximka200/NotesWebApp/pkg/repository"
	"github.com/maximka200/NotesWebApp/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// задаем формат логгов в виде json
	logrus.SetFormatter(new(logrus.JSONFormatter))
	// проверяем конфиг
	if err := InitConfig(); err != nil {
		logrus.Fatalf("config initialization error: %s", err.Error())
	}
	// проверяем загрузку .env
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env: %s", err.Error())
	}
	// поднимаем дб
	db, err := repository.NewPostgresDb(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to init db: %s", err.Error())
	}
	// инициализируем главные сущности для работы с логикой
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	/// *gin.Engine реализует интерфейс http.Handler
	handlers := handler.NewHandler(services)
	// инициализируем end-points
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
