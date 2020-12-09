package inventory

import (
	"fmt"
	"io"
	"net/http"
)

//GETUserInventory : GET inventory
func GETUserInventory(w http.ResponseWriter, req *http.Request) {
	bytes, err := ListUserInventory(1)
	if err != nil {
		fmt.Println("Can't serialize", bytes)
	}
	io.WriteString(w, string(bytes))
	fmt.Println(w)
	fmt.Println(*req)
}
