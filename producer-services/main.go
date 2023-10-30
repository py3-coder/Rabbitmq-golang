package main

import (
	"context"
	"message-queue/db"
	"message-queue/producer-services/routers"
	"net/http"
	"os"
	"os/signal"
	"time"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/rs/cors"
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
	log.Info("Producer serivices")
	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	s := http.Server{
		Addr:    ":8080",
		Handler: corsOpts.Handler(routers.NewRouter()),
	}
	db.CreateNewDBConnection()

	go func() {
		err := s.ListenAndServe()
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
	ctx, err := context.WithTimeout(context.Background(), 30*time.Second)
	if err != nil {
		log.Info("Error:", err)
	}
	s.Shutdown(ctx)
}
