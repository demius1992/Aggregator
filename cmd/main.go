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
	"time"
)

func main() {
	var wg sync.WaitGroup
	tickTime := 30 * time.Minute //ticker time in minutes for periodical collection the data

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

	if err := jobs.Collect(db, &wg); err != nil {
		logrus.Printf("an error occured while collecting data %s", err.Error())
	}

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

	ticker := time.NewTicker(tickTime)
	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				ticker.Stop()
				return
			case <-ticker.C:
				if err := jobs.Collect(db, &wg); err != nil {
					logrus.Printf("an error occured while collecting data %s", err.Error())
				}
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	done <- struct{}{}
	logrus.Println("ticker stopped")

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
