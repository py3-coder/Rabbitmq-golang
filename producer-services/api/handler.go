package api

import (
	"context"
	"math/rand"
	"message-queue/db"
	"message-queue/models"
	"message-queue/producer-services/producer"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type ProductServiceInterface interface {
	AddProduct(wr http.ResponseWriter, args *models.Product) *models.ProductResponse
}
type productservice struct {
	dbc db.MongoDBConnection
}

func ProductServices(dbc db.MongoDBConnection) ProductServiceInterface {
	return &productservice{dbc: dbc}
}

func (pd *productservice) AddProduct(wr http.ResponseWriter, args *models.Product) *models.ProductResponse {
	resp := models.ProductResponse{}
	collection := pd.dbc.ConnectToMongoDB().Collection("product-service-db")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	val := Random()
	temp := args
	data := models.ProductDetails{}
	data.ProductID = val
	data.ProductDescription = temp.ProductDescription
	data.ProductImages = temp.ProductImages
	data.ProductName = temp.ProductName
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	data.UserID = temp.UserID

	_, err1 := collection.InsertOne(ctx, data)
	if err1 != nil {
		log.Error("Found Error Inside Mongo dump it", err1)
		resp.Status = "Failed"
		resp.StatusCode = 400
		return &resp
	}
	go func() {
		err := producer.SendProductToQueue(val)
		if err != nil {
			log.Error("Found Error Inside Mongo dump it", err1)
		}
	}()
	resp.Status = "Sucess"
	resp.StatusCode = 200
	return &resp
}

func Random() int {
	return rand.Intn(100000)
}
