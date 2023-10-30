package main

import (
	"context"
	"message-queue/consumer-services/consumer"
	"message-queue/db"
	"os"
	"os/signal"
	"time"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&formatter.Formatter{
		HideKeys:      true,
		NoFieldsSpace: false,
		TrimMessages:  true,
	})

}
func main() {
	log.Info("Consumer serivices")
	db.CreateNewDBConnection()

	go func() {
		err := consumer.ConsumeQueue()
		if err != nil {
			log.Info("Error Starting server:", err)
			db.DisconnectToMongoDB()
			os.Exit(1)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)
	sig := <-c
	log.Info("Got signal:", sig)
	_, err := context.WithTimeout(context.Background(), 30*time.Second)
	if err != nil {
		log.Info("Error:", err)
	}
}
