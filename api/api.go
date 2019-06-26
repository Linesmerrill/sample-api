package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"go.uber.org/zap"

	"github.com/gorilla/mux"
)

type product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

// Init function to bring a new logger online
func Init() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
}

// Run starts the new mux router
func Run() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/product/{id}", productsHandler).Methods("GET")
	r.HandleFunc("/products", getProducts)
	http.Handle("/", r)
	log.Println("server running at localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(`{"health": "ok"}`))
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	thisProduct, err := getProduct(id)
	if err != nil {
		zap.NamedError("error getting product", err)
		return
	}

	out, err := json.Marshal(thisProduct)
	if err != nil {
		zap.NamedError("error marshaling product", err)
	}
	_, _ = w.Write(out)
}

func getProduct(productID string) (product, error) {
	file, err := ioutil.ReadFile(fmt.Sprintf("./test/products/%s", productID))
	if err != nil {
		zap.NamedError("error opening file", err)
	}

	myProduct := product{}
	_ = json.Unmarshal(file, &myProduct)
	return myProduct, nil
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(`{}`))
}
