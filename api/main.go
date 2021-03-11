package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/99-devops/kubernetes-yaml-generator/database"
	"github.com/gorilla/mux"
)

type KubeObject struct {
	ObjectName  string `json:"ObjectName"`
	YamlContent string `json:"YamlContent"`
}

//Home page of the api
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

//getAllKubeObjects gets all objects by calling database.GetAllObjects
func getAllKubeObjects(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllKubeObjects")
	collection, cancel := database.CreateConnection()
	defer cancel()
	byteAllKubeObjects := database.GetAllObjects(collection)
	var allKubeObject []KubeObject
	json.Unmarshal(byteAllKubeObjects, &allKubeObject)
	json.NewEncoder(w).Encode(allKubeObject)

}

//getSingleKubeObject get the document by calling database.GetSingleDocument
func getSingleKubeObject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	objectName := vars["ObjectName"]
	fmt.Println("Endpoint Hit: returnSingleKubeObjects")
	collection, cancel := database.CreateConnection()
	defer cancel()
	byteYamlContent := database.GetSingleDocument(objectName, collection)
	var object KubeObject
	json.Unmarshal(byteYamlContent, &object)
	json.NewEncoder(w).Encode(object)
}

//createNewObjectYaml creates new object by calling database.InsertSingleDocument
func createNewObjectYaml(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	collection, cancel := database.CreateConnection()
	defer cancel()
	insertOneResult := database.InsertSingleDocument(reqBody, collection)
	fmt.Println("Insert one result: ", insertOneResult)
}

//deleteObjectYaml deletes the object by calling database.DeleteSingleDocument
func deleteObjectYaml(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete API called...")
	vars := mux.Vars(r)
	objectName := vars["ObjectName"]
	collection, cancel := database.CreateConnection()
	defer cancel()
	deleteResult := database.DeleteSingleDocument(objectName, collection)
	fmt.Println("Delete Result: ", deleteResult)
}

//handleRequests handle all the request that the server gets and processed by respective functions
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/kubeObjects", getAllKubeObjects)
	myRouter.HandleFunc("/kubeObject", createNewObjectYaml).Methods("POST")
	myRouter.HandleFunc("/kubeObject/{ObjectName}", deleteObjectYaml).Methods("DELETE")
	myRouter.HandleFunc("/kubeObjects/{ObjectName}", getSingleKubeObject)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("API Server Started...")
	handleRequests()
}
