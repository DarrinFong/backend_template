package routers

import (
	"log"
	"net/http"

	"github.com/darrinfong/backend_template/services"
	"github.com/gorilla/mux"
)

//InventoryHandler : GET inventory
func InventoryHandler(r *mux.Router) {
	//GET User Inventory
	r.HandleFunc("/v1/inventory/GET", func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		log.Println(vars)

		bytes, err := services.ListUserInventory(1)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	})
}
