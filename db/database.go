package db

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDBConnection interface {
	ConnectToMongoDB() *mongo.Database
}
type customermongodb struct {
}

func MongodbDBInstance() MongoDBConnection {
	return &customermongodb{}
}

var (
	Client *mongo.Client
	ctxb   = context.Background()
)

func CreateNewDBConnection() bool {
	ctx, cancel := context.WithTimeout(ctxb, 30*time.Second)
	defer cancel()
	clientOpts := options.Client().ApplyURI("mongodb://localhost:27017/")

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Error("MongoDb connection drop:", err)
		return false
	}
	Client = client
	log.Info("Database is connected")
	return Ping()
}
func reconnect(client *mongo.Client) {
	for {
		if CreateNewDBConnection() {
			client = Client
			break
		}
		time.Sleep(time.Second * 5)
	}
}
func HandlePoolMonitor(evt *event.PoolEvent) {
	switch evt.Type {
	case event.PoolClosedEvent:
		log.Error("DB connection closed and tring to reconnect.")
		reconnect(Client)
	}
}

func Ping() bool {
	if err := Client.Ping(ctxb, readpref.Primary()); err != nil {
		return false
	}
	return true
}

func (cm *customermongodb) ConnectToMongoDB() *mongo.Database {
	return Client.Database("service-db")
}

func DisconnectToMongoDB() {
	defer func() {
		if err := Client.Disconnect(ctxb); err != nil {
			log.Error("DB Disconnector break due to address port", err)
		}
	}()

}
