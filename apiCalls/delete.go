package main

import (
	"fmt"
	"net/http"

	"github.com/pradipsilwal/kubernetes-yaml-generator/utils"
)

//delete sends delete request to the api server to delete the object
func delete(kubeObjectName string) {
	fmt.Println("Performing Http Delete...")

	//Creating the request
	request, e := http.NewRequest("DELETE", "http://localhost:10000/kubeObject/"+kubeObjectName, nil)

	//Setting the content type of the request header
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	//making a request
	http.DefaultClient.Do(request)

	fmt.Println("Request: ", request)
	utils.CheckError(e)
	fmt.Println("Deleted")
}
