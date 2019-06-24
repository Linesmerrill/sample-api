package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Run starts the new mux router
func Run() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	//r.HandleFunc("/products", ProductsHandler)
	//r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)
	log.Println("server running at localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(`{"health": "ok"}`))
}
