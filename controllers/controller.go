package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	model "github.com/The-Fermits/Golang/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://lco:hitesh@cluster0.tuuzb.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

const dbName = "netflix"
const collectionName = "watchlist"

// very important for connect with above url of mongoDB
var collection *mongo.Collection

// connect with mongoDB
// whenever we perform any mongoDB operation we need to provide context ALWAYS

func init() { // this execute only once
	// client - option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection Success")

	collection = client.Database(dbName).Collection(collectionName)

	//  collection reference is ready
	fmt.Println("collection instance is ready ")
}

func checkErr(err error) {
	if err != nil {
		log.Println("Error: ", err) // Log the error but don't exit
	}
}

// insert data in mongoDB with helper functions

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Println("Error inserting movie:", err)
		return
	}
	fmt.Println("Inserted movie with ID:", inserted.InsertedID)
}

func updateOneMovie(movieID string) {
	// grab the _id that mongoDB understands
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	checkErr(err)
	fmt.Println("Movie status is updated and modified count is: ", result.ModifiedCount)
}

// delete one movie form mongoDB
func deleteOneMovie(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	deleteCount, _ := collection.DeleteOne(context.Background(), filter)
	fmt.Println("One Movie is deleted  with count: ", deleteCount.DeletedCount)
}

// delete all movies (just change filter )
func delteAllMovies() {

	// filter := bson.D{{}} pass nothing
	deleteCount, _ := collection.DeleteMany(context.Background(), bson.D{{}})
	fmt.Println("All movies are deleted with count: ", deleteCount.DeletedCount)
}

// read all collection of movies in mongoDB , something new is here
func getAllMovies() []primitive.M {
	// cur -> used for cursor , cursor has all data form mongoDB, it is like a linkedList
	cur, err := collection.Find(context.Background(), bson.D{{}}) // find without id means filter = {{}}
	checkErr(err)

	var movies []primitive.M // we will return this

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		checkErr(err)
		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())
	return movies
}

// Now write Actual Controllers which are exportable
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "/application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie model.Netflix

	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		log.Println("JSON decode error:", err)
		return
	}

	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "/application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "/application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "/application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	delteAllMovies()
	json.NewEncoder(w).Encode("all deleted")
}

// serve Home Route
func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>I am HTML who serving Home route here </h1>"))
}
