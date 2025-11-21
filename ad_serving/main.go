package main

import (
	"ad_serving/web/routes"
	"fmt"
	"net/http"
)

func main() {

	routes.Kbbroutes()

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)

}
