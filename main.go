package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	inventory "github.com/darrinfong/backend_template/server/inventory"
)

func init() {
	// Declare mocks
	// Use DatabaseRepo or MockRepo depending on env.
	inventory.InventoryRepo = inventory.MockRepo{}

	// Attach handlers
	http.HandleFunc("/inventory/GET", inventory.GETUserInventory)
}

func main() {
	// seed random for mock function
	rand.Seed(time.Now().UnixNano())
	log.Println("Backend listening on http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
