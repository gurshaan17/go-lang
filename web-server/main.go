package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
  if err := r.ParseForm(); err!= nil {
    fmt.Fprintf(w, "parseform error")
    return 
  }
  fmt.Fprintf(w, "post successful")
  name := r.FormValue("name")
  address := r.FormValue("address")
  fmt.Fprintf(w, "Name: %s\n", name)
  fmt.Fprintf(w, "Address: %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/hello" {
    http.Error(w, "port not found", http.StatusNotFound)
  }
  if r.Method != "GET" {
    http.Error(w, "method not supported", http.StatusNotFound)
  }
  fmt.Fprintf(w, "hello ")
}

func main(){
  fileServer := http.FileServer(http.Dir("./static"))
  http.Handle("/",fileServer)
  http.HandleFunc("/form",formHandler)
  http.HandleFunc("/hello",helloHandler)

  fmt.Println("Server starting at port 8080")
  if err := http.ListenAndServe(":8080", nil); err != nil{
    log.Fatal(err)
  }
}