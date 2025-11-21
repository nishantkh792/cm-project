package routes

import (
	"ad_serving/web/controllers"
	"fmt"
	"net/http"
)

func Kbbroutes() {
	// Serve tag.js
	fmt.Println("inside routes")
	http.HandleFunc("/server/tag.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "app/pkg/static/tag1.js")
	})

	// Serve ad.html
	http.HandleFunc("/server/ad.html", controllers.Kbbcontrollers)
}
