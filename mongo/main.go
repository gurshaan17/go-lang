package main

import (
  "net/http"
  "github.com/julienschmidt/httprouter"
  "gopkg.in/mgo.v2"
  "github.com/gurshaan17/go-lang/mongo/controllers"
)

func main(){
  r :=  httprouter.New()  
  uc := controllers.NewUserController(getSession())
  r.GET("/user/:id" , uc.GetUser)
  r.POST("/user", uc.CreateUser)
  r.DELETE("/user/:id", uc.DeleteUser)
  http.ListenAndServe(":8080", r)
}

func getSession() *mgo.Session{
  s, err := mgo.Dial("mongo url")
  if err != nil {
    panic(err)
  }
  return s;
}