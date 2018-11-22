package main

import (
	"flag"
	"log"

	"github.com/ilgooz/service-articles/articles"
	"github.com/mesg-foundation/core/x/xsignal"
	mesg "github.com/mesg-foundation/go-service"
)

var (
	mongoAddr = flag.String("mongoAddr", "", "MongoDB's address")
	dbName    = flag.String("dbName", "", "Database name")
)

func main() {
	flag.Parse()

	service, err := mesg.New()
	if err != nil {
		log.Fatal(err)
	}

	storage, err := articles.NewMongoStorage(*mongoAddr, *dbName)
	if err != nil {
		log.Fatal(err)
	}

	a, err := articles.New(service, storage)
	if err != nil {
		log.Fatal(err)
	}

	// start the articles service.
	go func() {
		log.Println("articles service has been started")

		if err := a.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	// wait for interrupt and gracefully shutdown the articles service.
	<-xsignal.WaitForInterrupt()

	log.Println("shutting down...")

	if err := a.Close(); err != nil {
		log.Fatal(err)
	}

	log.Println("shutdown")
}
