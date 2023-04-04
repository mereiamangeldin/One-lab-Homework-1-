package main

import (
	"fmt"
	"github.com/mereiamangeldin/One-lab-Homework-1/config"
	"github.com/mereiamangeldin/One-lab-Homework-1/repository"
	"github.com/mereiamangeldin/One-lab-Homework-1/service"
	"github.com/mereiamangeldin/One-lab-Homework-1/transport/http"
	"github.com/mereiamangeldin/One-lab-Homework-1/transport/http/handler"
	"log"
)

func main() {
	log.Fatalln(fmt.Sprintf("Service shut down:%s", run()))
}

func run() error {
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
	return srv.Run(conf.Port, srv.InitRoutes())
}
