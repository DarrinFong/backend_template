package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/darrinfong/backend_template/routers"
	"github.com/darrinfong/backend_template/services"
	"github.com/gorilla/mux"
)

// Define active DB: 0-DB, 1-Mock
const dbType int = 1

func init() {
	// seed random for mocks
	rand.Seed(time.Now().UnixNano())

	// Set active repo
	services.SetInventoryDB(dbType)
}

func main() {
	r := mux.NewRouter()
	routers.InventoryHandler(r)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8000", nil))
	log.Println("Backend listening on http://localhost:8000/")
}
