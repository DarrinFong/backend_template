package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/darrinfong/backend_template/routers"
	"github.com/darrinfong/backend_template/services"
	"github.com/darrinfong/backend_template/utils"
	"github.com/gorilla/mux"
)

func initDatabase() {
	// databaseInfo := fmt.Sprintf(
	// 	"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 	"127.0.0.1",
	// 	5432,
	// 	"",
	// 	"",
	// 	"")

	// db, err := sql.Open("postgres", databaseInfo)
	// if err != nil {
	// 	log.Fatalf("Could not create DB")
	// 	log.Fatalf(err.Error())
	// }

	services.InventoryRepo = &utils.HMockDB{}
	// services.InventoryRepo = &dbinterface.SQLDB{DB: db}
}

func init() {
	// seed random for mocks
	rand.Seed(time.Now().UnixNano())
	initDatabase()
}

func main() {
	r := mux.NewRouter()
	routers.InventoryHandler(r)
	http.Handle("/", r)

	log.Println("Backend listening on http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
