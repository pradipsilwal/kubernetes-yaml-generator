package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type KubeObject struct {
	ObjectName  string `json:"ObjectName"`
	YamlContent string `json:"YamlContent"`
}

var KubeObjects []KubeObject

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllKubeObjects(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllKubeObjects")
	json.NewEncoder(w).Encode(KubeObjects)
}

func returnSingleKubeObject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["ObjectName"]
	fmt.Println("Endpoint Hit: returnSingleKubeObjects")
	for _, object := range KubeObjects {
		if object.ObjectName == key {
			json.NewEncoder(w).Encode(object)
		}
	}
}

func createNewObjectYaml(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var kubeObject KubeObject
	json.Unmarshal(reqBody, &kubeObject)
	fmt.Println(kubeObject)
	KubeObjects = append(KubeObjects, kubeObject)
	json.NewEncoder(w).Encode(kubeObject)
}

func deleteObjectYaml(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["ObjectName"]
	for index, kubeObject := range KubeObjects {
		if kubeObject.ObjectName == name {
			KubeObjects = append(KubeObjects[:index], KubeObjects[index+1:]...)
		}
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/kubeObjects", returnAllKubeObjects)
	myRouter.HandleFunc("/kubeObject", createNewObjectYaml).Methods("POST")
	myRouter.HandleFunc("/kubeObject/{ObjectName}", deleteObjectYaml).Methods("DELETE")
	myRouter.HandleFunc("/kubeObjects/{ObjectName}", returnSingleKubeObject)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	KubeObjects = []KubeObject{
		KubeObject{ObjectName: "Pod", YamlContent: "This is a content"},
		KubeObject{ObjectName: "Service", YamlContent: "This is a service"},
	}
	handleRequests()
}
