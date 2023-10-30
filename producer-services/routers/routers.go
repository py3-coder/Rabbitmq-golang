package routers

import (
	"encoding/json"
	"message-queue/db"
	"message-queue/models"
	"message-queue/producer-services/api"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var (
	datab = db.MongodbDBInstance()
	pd    = api.ProductServices(datab)
)

func NewRouter() *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/", indexGetHandler).Methods("GET")
	//add product data::
	route.HandleFunc("/addproduct", addProductServices).Methods("POST")
	return route
}

// Common Util Routers
func indexGetHandler(wr http.ResponseWriter, r *http.Request) {
	wr.Header().Set("Content-Type", "application/json")
	wr.WriteHeader(http.StatusAccepted)
	wr.Write([]byte("Unauthorized Page"))
}

// addproduct::
func addProductServices(wr http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data := models.Product{}
	if r.Body == nil {
		http.Error(wr, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Error("Decorder in Controller:", err)
		http.Error(wr, err.Error(), 400)
		return
	}
	result := pd.AddProduct(wr, &data)
	wr.Header().Set("Content-Type", "application/json")
	wr.WriteHeader(http.StatusOK)
	json.NewEncoder(wr).Encode(result)
}
