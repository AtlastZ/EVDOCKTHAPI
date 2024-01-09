package main

import (
	"context"
	"fmt"
	"encoding/json"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type Person struct {
    Name string `bson:"name"`
    Age  int `bson:"age"`
	Brand string `bson:"brand"`
}

func insertINFO(client *mongo.Client) *mongo.InsertOneResult {
	collection := client.Database("EVDOCK").Collection("info")
	result, err := collection.InsertOne(context.TODO(), bson.D{{Key: "name", Value: "HelloWorld2"}, {Key: "age", Value: 30}})
	fmt.Println("Insert! ->")
	if err != nil {
		fmt.Println(err.Error())
	}
	return result
}

func deleteINFO(client *mongo.Client) *mongo.DeleteResult {
	collection := client.Database("EVDOCK").Collection("info")
	result, err := collection.DeleteOne(context.TODO(), bson.D{{Key: "Brand", Value: "Neta"}})
	// collection.Find(context.TODO(), bson.D{{Key: "", Value: ""}})
	fmt.Println("Delete! ->")
	if err != nil {
		fmt.Println(err.Error())
	}
	return result
}

func getINFO(client *mongo.Client) any {
	collection := client.Database("EVDOCK").Collection("info")
	// filter := bson.D{{Key: "age", Value: 30}} // Adjust filter as needed

    cursor , err := collection.Find(context.TODO(),bson.D{})
    if err != nil {
        log.Fatal(err)
    }
	fmt.Println("got data from database")
    defer cursor.Close(context.TODO())

    var results []Person
    if err := cursor.All(context.TODO(), &results); err != nil {
        log.Fatal(err)
    }

    jsonData, err := json.MarshalIndent(results, "", "  ") // Pretty-print JSON
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(jsonData))
	return jsonData
}
func updateINFO(){

}

func main() {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://admin:kV4vxejori@atlascluster.hz1anm1.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	fmt.Println("try to insert a new value in the database")
	getINFO(client)
}
