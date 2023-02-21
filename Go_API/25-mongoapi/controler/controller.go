package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/arjun/modules/25-mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://krishna:2812@netflix.r5erkdk.mongodb.net/?retryWrites=true&w=majority"
const dbName = "Netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	collection = (*mongo.Collection)(client.Database(dbName).Collection(colName))
	// collection instance is ready

}

// helpers method
// insert one record
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("inserted one movie with ID", inserted.InsertedID)
}

// update one record
func updateOneMovie(moveId string) {
	id, _ := primitive.ObjectIDFromHex(moveId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified counter update", result.ModifiedCount)
}

// delete one record
func deleteOneMovie(moveId string) {
	id, _ := primitive.ObjectIDFromHex(moveId)

	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got deleted", deleteCount)
}

// delete many record
func deleteManyMovie() int64 {
	deletecount, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("delete all recors count", deletecount.DeletedCount)
	return deletecount.DeletedCount
}

// final -collect all movies
func collectAllMovies() []primitive.M {
	result, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M
	for result.Next(context.Background()) {
		var movie bson.M
		err := result.Decode(movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer result.Close(context.Background())
	return movies
}

// Actual controller via api call - we can put all these in another files
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	allMovies := collectAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Controll-Allow-Methods", "POST")

	var movies model.Netflix

	json.NewDecoder(r.Body).Decode(&movies)
	insertOneMovie(movies)
	json.NewEncoder(w).Encode(movies)

}

func MarckAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Controll-Allow-Methods", "PUST")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode("Updated mark as watched")
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Controll-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}
func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Controll-Allow-Methods", "DELETE")

	count := deleteManyMovie()
	json.NewEncoder(w).Encode(count)

}
