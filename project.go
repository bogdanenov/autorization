
package main

import (
  "fmt"
  "net/http"
  "html/template"
)

func main(){
  fmt.Println("Listening on port: 3000")
  http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
  http.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("./image/"))))
  http.HandleFunc("/", indexHandler)
  http.ListenAndServe(":3000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  t,err := template.ParseFiles("templates/index.html")
  if err != nil {
    fmt.Fprintf(w, err.Error())
  }
  t.ExecuteTemplate(w, "index", nil)
}
