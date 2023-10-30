package consumer

import (
	"context"
	"io"
	"message-queue/db"
	"message-queue/models"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	datab = db.MongodbDBInstance()
	path  = "E:/golang_project/Message-Queue-System/image/"
)

func ConsumeQueue() error {
	rabbitMQURL := "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Info("Error:", err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"product_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Process incoming messages
	for msg := range msgs {
		productID := string(msg.Body)
		log.Info("Received product_id:", productID)
		i, err := strconv.Atoi(productID)
		if err != nil {
			log.Fatal(err)
		}
		err1 := processImages(i)
		if err1 != nil {
			log.Info("Error", err1)
		}
	}
	return nil
}

func processImages(productID int) error {
	log.Info("ProductID:", productID)
	collection := datab.ConnectToMongoDB().Collection("product-service-db")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	resp := models.ProductDetails{}
	filter := bson.D{{Key: "product_id", Value: productID}}
	err := collection.FindOne(ctx, filter).Decode(&resp)
	if err != nil {
		log.Error("Error Found", err)
	}

	for _, imageURL := range resp.ProductImages {
		saveImageLocally(imageURL)
		resp.CompressedProductImages = append(resp.CompressedProductImages, path+imageURL)
	}
	replacement := bson.M{
		"product_id":                productID,
		"product_name":              resp.ProductName,
		"product_description":       resp.ProductDescription,
		"product_images":            resp.ProductImages,
		"product_price":             resp.ProductPrice,
		"compressed_product_images": resp.CompressedProductImages,
		"created_at":                time.Now(),
		"updated_at":                time.Now(),
	}
	// replaceone ::
	result, err := collection.ReplaceOne(ctx, filter, replacement)
	if err != nil {
		log.Info("Error Found", err)
	}
	if result.MatchedCount == 1 {
		log.Info("Compressed Image Path Updated")
	} else {
		log.Info("Not Replaced::")
		return err
	}
	return nil
}

func saveImageLocally(image string) error {
	imageURL := image
	localDirectory := path

	resp, err := http.Get(imageURL)
	if err != nil {
		log.Info("Failed to download image:", err)
		return nil
	}
	defer resp.Body.Close()

	fileName := filepath.Base(imageURL)
	err = os.MkdirAll(localDirectory, os.ModePerm)
	if err != nil {
		log.Info("Failed to create local directory:", err)
		return nil
	}

	filePath := filepath.Join(localDirectory, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		log.Info("Failed to create a local file:", err)
		return nil
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Info("Failed to save image to the local file:", err)
		return nil
	}
	log.Info("Path:", localDirectory+imageURL)
	return nil
}
