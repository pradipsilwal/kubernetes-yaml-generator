package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/99-devops/kubernetes-yamal-generator/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type KubeObject struct {
	ObjectName  string
	YamlContent string
}

func main() {
	DATABASE_NAME := os.Getenv("DATABASE_NAME")
	COLLECTION_NAME := os.Getenv("COLLECTION_NAME")
	DATABASE_USERNAME := os.Getenv("DATABASE_USERNAME")
	DATABASE_PASSWORD := os.Getenv("DATABASE_PASSWORD")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://"+DATABASE_USERNAME+":"+DATABASE_PASSWORD+"@cluster0.u43qj.mongodb.net/"+DATABASE_NAME+"?retryWrites=true&w=majority",
	))
	if err != nil {
		log.Fatal(err)
	}

	// get collection as ref'
	collection := client.Database(DATABASE_NAME).Collection(COLLECTION_NAME)

	lines, e := utils.GetStrings("pod.yaml")
	utils.CheckError(e)
	stringLines := utils.GetStringFromSlice(lines)
	kubeObject := KubeObject{"Pod", stringLines}
	// insertSingleDocument(kubeObject, collection)
	printSingleDocument(kubeObject.ObjectName, collection)
	// updatedLines, e := utils.GetStrings("podupdate.yaml")
	// utils.CheckError(e)
	// updatedStringLines := utils.GetStringFromSlice(updatedLines)
	// updatedKubeObject := KubeObject{"Pod", updatedStringLines}
	// updateDocument(updatedKubeObject, collection)
	// deleteSingleDocument(kubeObject.ObjectName, collection)
	// printSingleDocument(kubeObject.ObjectName, collection)

}

func insertSingleDocument(kubeObject KubeObject, collection *mongo.Collection) {
	fmt.Println("Inserting document with object name: ", kubeObject.ObjectName)
	_, e := collection.InsertOne(context.TODO(), kubeObject)
	utils.CheckError(e)
}

func printSingleDocument(objectName string, collection *mongo.Collection) {
	var kubeObject KubeObject
	filter := bson.D{{Key: "objectname", Value: objectName}} // 'Key:' and 'Value:' keywords can be omitted		filter := bson.D{{"name", name}}	this works as well
	fmt.Println("Printing document with object name: ", objectName)
	e := collection.FindOne(context.TODO(), filter).Decode(&kubeObject)
	utils.CheckError(e)
	fmt.Println(kubeObject.YamlContent)
}

func deleteSingleDocument(objectName string, collection *mongo.Collection) {
	filter := bson.D{{Key: "objectname", Value: objectName}}
	fmt.Println("Deleting document with object name: ", objectName)
	_, e := collection.DeleteMany(context.TODO(), filter)
	utils.CheckError(e)

}

func updateDocument(updatedKubeObject KubeObject, collection *mongo.Collection) {
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

func getAllObjects(collection *mongo.Collection) {
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
	for _, value := range results {
		fmt.Println(*value)
	}
}

func insertMultipleDocument(collection *mongo.Collection, multipleObjects []interface{}) {
	_, e := collection.InsertMany(context.TODO(), multipleObjects)
	utils.CheckError(e)
}
