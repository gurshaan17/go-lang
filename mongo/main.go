package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/gurshaan17/go-lang/mongo/controllers"
    "github.com/julienschmidt/httprouter"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    r := httprouter.New()
    client := getClient()

    uc := controllers.NewUserController(client)

    r.GET("/user/:id", uc.GetUser)
    r.POST("/user", uc.CreateUser)
    r.DELETE("/user/:id", uc.DeleteUser)

    fmt.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}

func getClient() *mongo.Client {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatalf("Failed to ping MongoDB: %v", err)
    }

    fmt.Println("Connected to MongoDB!")
    return client
}