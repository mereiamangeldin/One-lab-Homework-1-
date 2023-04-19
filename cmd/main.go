package main

import (
	"context"
	"fmt"
	"github.com/mereiamangeldin/One-lab-Homework-1/config"
	"github.com/mereiamangeldin/One-lab-Homework-1/repository"
	"github.com/mereiamangeldin/One-lab-Homework-1/service"
	"github.com/mereiamangeldin/One-lab-Homework-1/transport/http"
	"github.com/mereiamangeldin/One-lab-Homework-1/transport/http/handler"
	"log"
	"os"
	"os/signal"
)

// @title           One lab hw
// @version         1.0
// @description     This is a sample server celler server.

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey  ApiKeyAuth
// @in header
// @name Authorization
func main() {
	log.Fatalln(fmt.Sprintf("Service shut down:%s", run()))
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	gracefullyShutdown(cancel)
	conf, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}
	repo, err := repository.New(conf)
	if err != nil {
		log.Fatal(err.Error())
	}
	svc, err := service.NewManager(repo)
	if err != nil {
		log.Fatal(err.Error())
	}
	h := handler.NewManager(conf, svc)
	srv := http.NewServer(conf, h)
	return srv.Run(ctx)
}

func gracefullyShutdown(c context.CancelFunc) {
	osC := make(chan os.Signal, 1)
	signal.Notify(osC, os.Interrupt)
	go func() {
		log.Print(<-osC)
		c()
	}()
}
