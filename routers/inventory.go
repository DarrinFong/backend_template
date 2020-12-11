package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/darrinfong/backend_template/services"
	"github.com/gorilla/mux"
)

//InventoryHandler : GET inventory
func InventoryHandler(r *mux.Router) {
	//GET User Inventory
	r.HandleFunc("/v1/inventory/GET/{userID:[0-9]+}", func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		userID, err := strconv.ParseInt(params["userID"], 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		items := services.GetSellerInventory(int(userID))
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(items)
	})

	//GET Item
	r.HandleFunc("/v1/item/GET/{itemID:[0-9]+}", func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		itemID, err := strconv.ParseInt(params["itemID"], 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		item := services.GetItem(int(itemID))
		json.NewEncoder(w).Encode(item)
	})
}
