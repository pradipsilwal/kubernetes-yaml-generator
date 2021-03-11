package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/99-devops/kubernetes-yaml-generator/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type KubeObject struct {
	ObjectName  string `json:"ObjectName"`
	YamlContent string `json:"YamlContent"`
}

//CreateConnection creates a connection to MongoDB Atlas database and returns *mongo.Collection
//along with context.CancelFunc
func CreateConnection() (*mongo.Collection, context.CancelFunc) {

	//Defining the credentials to be taken from environment variables
	DATABASE_NAME := os.Getenv("DATABASE_NAME")
	COLLECTION_NAME := os.Getenv("COLLECTION_NAME")
	DATABASE_USERNAME := os.Getenv("DATABASE_USERNAME")
	DATABASE_PASSWORD := os.Getenv("DATABASE_PASSWORD")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	//Creating the connection
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://"+DATABASE_USERNAME+":"+DATABASE_PASSWORD+"@cluster0.u43qj.mongodb.net/"+DATABASE_NAME+"?retryWrites=true&w=majority",
	))
	if err != nil {
		log.Fatal(err)
	}

	//Get collection as ref'
	collection := client.Database(DATABASE_NAME).Collection(COLLECTION_NAME)

	return collection, cancel
}

//InsertSIngleDocument inserts the Single document to mongodb collection
func InsertSingleDocument(byteKubeObject []byte, collection *mongo.Collection) *mongo.InsertOneResult {
	var kubeObject KubeObject
	json.Unmarshal(byteKubeObject, &kubeObject)
	fmt.Println("Inserting document with object name: ", kubeObject.ObjectName)
	insertOneResult, e := collection.InsertOne(context.TODO(), kubeObject)
	utils.CheckError(e)
	return insertOneResult
}

//GetSingleDocument gets single document from database that matches the kubernetes Object Name
func GetSingleDocument(objectName string, collection *mongo.Collection) []byte {
	var kubeObject KubeObject
	filter := bson.D{{Key: "objectname", Value: objectName}} // 'Key:' and 'Value:' keywords can be omitted		filter := bson.D{{"name", name}}	this works as well
	e := collection.FindOne(context.TODO(), filter).Decode(&kubeObject)
	utils.CheckError(e)
	byteKubeObject, e := json.Marshal(kubeObject)
	utils.CheckError(e)
	return byteKubeObject
}

//DeleteSingleDocument deletes the object document from database that matches the given object name
func DeleteSingleDocument(objectName string, collection *mongo.Collection) *mongo.DeleteResult {
	filter := bson.D{{Key: "objectname", Value: objectName}}
	fmt.Println("Deleting document with object name: ", objectName)
	deleteResult, e := collection.DeleteMany(context.TODO(), filter)
	utils.CheckError(e)
	return deleteResult
}

//UpdateDocument adds new kubernetes object in the database
func UpdateDocument(updatedKubeObject KubeObject, collection *mongo.Collection) {
	filter := bson.D{{Key: "objectname", Value: updatedKubeObject.ObjectName}}
	update := bson.D{
		{Key: "$set", Value: bson.D{ // '$set' set the value of the field in the document
			{Key: "yamlcontent", Value: updatedKubeObject.YamlContent},
		}},
	}

	fmt.Println("Updating document with object name: ", updatedKubeObject.ObjectName)
	_, e := collection.UpdateOne(context.TODO(), filter, update)
	utils.CheckError(e)
}

//GetAllObjects gets all the objects from mongo collection
func GetAllObjects(collection *mongo.Collection) []byte {
	var results []*KubeObject
	findOptions := options.Find()
	cur, e := collection.Find(context.TODO(), bson.D{{}}, findOptions) // 'bosn.D{{}}' to apply any filter. Here the filter is empty.
	utils.CheckError(e)

	// Iterating through the cursor for each value
	for cur.Next(context.TODO()) {
		var object KubeObject
		e := cur.Decode(&object)
		utils.CheckError(e)
		results = append(results, &object)
	}

	e = cur.Err()
	utils.CheckError(e)
	fmt.Println()
	cur.Close(context.TODO())
	byteAllKubeObjects, e := json.Marshal(results)
	utils.CheckError(e)
	return byteAllKubeObjects
}

// func InsertMultipleDocument(collection *mongo.Collection, multipleObjects []interface{}) {
// 	_, e := collection.InsertMany(context.TODO(), multipleObjects)
// 	utils.CheckError(e)
// }
