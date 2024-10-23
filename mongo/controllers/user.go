package controllers

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/gurshaan17/go-lang/mongo/models"
    "github.com/julienschmidt/httprouter"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
    client *mongo.Client
}

func NewUserController(client *mongo.Client) *UserController {
    return &UserController{client}
}

// GetUser fetches a user by ID from MongoDB
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    id := p.ByName("id")

    oid, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    u := models.User{}
    collection := uc.client.Database("go").Collection("users")

    err = collection.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&u)
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(u)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    u := models.User{}
    json.NewDecoder(r.Body).Decode(&u)

    u.Id = primitive.NewObjectID() // Generate a new ObjectID

    collection := uc.client.Database("go").Collection("users")
    _, err := collection.InsertOne(context.TODO(), u)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "Error inserting new user: %v\n", err)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(u)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    id := p.ByName("id")

    oid, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    collection := uc.client.Database("go").Collection("users")
    _, err = collection.DeleteOne(context.TODO(), bson.M{"_id": oid})
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Deleted user with ID: %s\n", id)
}