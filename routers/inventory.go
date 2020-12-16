package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/darrinfong/backend_template/models"
	"github.com/darrinfong/backend_template/services"
	"github.com/gorilla/mux"
)

//InventoryHandler : GET inventory
func InventoryHandler(r *mux.Router) {
	//GET User Inventory
	r.HandleFunc("/inventory/", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			items, err := services.GetInventory()
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			json.NewEncoder(w).Encode(items)
		case "POST":
			var i models.NewItem
			err := json.NewDecoder(req.Body).Decode(&i)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			item, _ := services.CreateItem(i)
			json.NewEncoder(w).Encode(item)
		}
	}).Methods("GET", "POST")

	//GET Item
	r.HandleFunc("/inventory/{itemID:[0-9]+}", func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		itemID, err := strconv.ParseInt(params["itemID"], 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		switch req.Method {
		case "GET":
			item, _ := services.GetItem(int(itemID))
			json.NewEncoder(w).Encode(item)
		case "PUT":
		case "DELETE":
		}
	}).Methods("GET", "POST", "PUT", "DELETE")
}
