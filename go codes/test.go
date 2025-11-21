package main

// import (
//     "html/template"
//     "net/http"
// )

// type PageData struct {
//     Keywords []string
// }

// func main() {

//     http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

//         keywords := []string{"iphone", "insurance", "loans"}

//         tmpl := template.Must(template.ParseFiles("test.html"))

//         data := PageData{Keywords: keywords}

//         tmpl.Execute(w, data)
//     })

//     http.ListenAndServe(":8080", nil)
// }
