package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Food represents each food item
type Food struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name"`
	Energy       int                `json:"energy" bson:"energy"`
	Protein      float64            `json:"protein" bson:"protein"`
	Fat          float64            `json:"fat" bson:"fat"`
	Carbohydrate float64            `json:"carbohydrate" bson:"carbohydrate"`
	Sugars       float64            `json:"sugars" bson:"sugars"`
	DietaryFibre float64            `json:"dietary-fibre" bson:"dietary-fibre"`
	Sodium       float64            `json:"sodium" bson:"sodium"`
}

type errorJSON struct {
	Statuscode int    `json:"statuscode"`
	Err        string `json:"err"`
}

//Helper function to pretty print json responses.
func printJSON(w io.Writer, result interface{}) {
	bytedata, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Fprint(w, err)
	}
	fmt.Fprint(w, string(bytedata))
}

//CreateDataHandler For CREATE operation
func CreateDataHandler(collection *mongo.Collection) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("content-type", "application/json")
		var food Food
		if err := json.NewDecoder(request.Body).Decode(&food); err != nil {
			panic(err)
		}
		result, err := collection.InsertOne(context.TODO(), food)
		if err != nil {
			panic(err)
		}

		response.WriteHeader(http.StatusCreated)
		printJSON(response, result)

	}
}

//ViewAllDataHandler for READ operation
func ViewAllDataHandler(collection *mongo.Collection) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("content-type", "application/json")
		var foods []Food
		ctx := context.TODO()
		cursor, err := collection.Find(ctx, bson.D{})
		if err != nil {
			panic(err)
		}

		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var food Food
			cursor.Decode(&food)
			foods = append(foods, food)
		}

		if err := cursor.Err(); err != nil {
			panic(err)
		}
		printJSON(response, foods)

	}
}

//ViewDataByIDHandler for READ operation
func ViewDataByIDHandler(collection *mongo.Collection) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("content-type", "application/json")
		vars := mux.Vars(request)
		id := vars["id"]
		objectID, _ := primitive.ObjectIDFromHex(id)
		filter := bson.D{{"_id", objectID}}
		var food Food
		err := collection.FindOne(context.TODO(), filter).Decode(&food)
		if err != nil {
			printJSON(response, errorJSON{Statuscode: http.StatusNotFound, Err: err.Error()})
		} else {
			response.WriteHeader(http.StatusFound)
			printJSON(response, food)
		}
	}

}

//DeleteDataByIDHandler for DELETE operation
func DeleteDataByIDHandler(collection *mongo.Collection) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		id := vars["id"]
		objectID, _ := primitive.ObjectIDFromHex(id)
		filter := bson.D{{"_id", objectID}}
		var food Food
		collection.FindOneAndDelete(context.TODO(), filter).Decode(&food)
		printJSON(response, food)
	}
}

//UpdateDataByIDHandler for UPDATE operation
func UpdateDataByIDHandler(collection *mongo.Collection) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("content-type", "application/json")
		vars := mux.Vars(request)
		id := vars["id"]
		objectID, _ := primitive.ObjectIDFromHex(id)
		filter := bson.D{{"_id", objectID}}
		var food Food
		if err := json.NewDecoder(request.Body).Decode(&food); err != nil {
			panic(err)
		}
		result := collection.FindOneAndUpdate(context.Background(), filter, bson.M{"$set": food}, options.FindOneAndUpdate().SetReturnDocument(1))
		decoded := Food{}
		if err := result.Decode(&decoded); err != nil {
			panic(err)
		}

		printJSON(response, decoded)
	}

}
