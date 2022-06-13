package main

import (
	"Aggregator"
	"Aggregator/handlers"
	"Aggregator/jobs"
	"Aggregator/repository"
	"Aggregator/service"
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing gonfigs %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables %s", err.Error())
	}

	db, err := repository.NewPostgresDB(context.Background(), repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db %s", err.Error())
	}

	go func() {
		defer wg.Done()
		if err := jobs.Collect(db); err != nil {
			logrus.Printf("an error occured while collecting data %s", err.Error())
		}
	}()
	wg.Wait()

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handlers.NewHandler(services)

	srv := new(Aggregator.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			logrus.Fatalf("an error occured while running http server %s", err.Error())
		}
	}()
	logrus.Println("Application started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("application is shutting down")

	if err = srv.ShutDown(context.Background()); err != nil {
		logrus.Errorf("an error occured while server shutting down: %s", err.Error())
	}

	db.Close()
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
