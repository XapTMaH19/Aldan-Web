package main

import (
	"github.com/XapTMaH19/AldanWeb"
	"github.com/XapTMaH19/AldanWeb/internal/handler"
	"github.com/XapTMaH19/AldanWeb/internal/service"
	"github.com/XapTMaH19/AldanWeb/internal/storage"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %v", err.Error())
	}

	/*if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %v", err.Error())
	}*/

	/*db, err := storage.NewPostgresDB(storage.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %v", err.Error())
	}*/
	r, err := storage.NewClientRabbitMQ(storage.RabbitMQConfig{
		Username:          viper.GetString("rabbitmq.username"),
		Password:          viper.GetString("rabbitmq.password"),
		Host:              viper.GetString("rabbitmq.host"),
		Port:              viper.GetString("rabbitmq.port"),
		RegisterNameQueue: viper.GetString("rabbitmq.queues.register"),
		LoginNameQueue:    viper.GetString("rabbitmq.queues.login"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize rabbit: %v", err.Error())
	}
	storage := storage.NewStorage(r)
	services := service.NewService(storage)
	handlers := handler.NewHandler(services)
	srv := new(AldanWeb.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("произошла ошибка при запуске http-сервера: %v", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
