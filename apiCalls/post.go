package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/99-devops/kubernetes-yaml-generator/utils"
)

//post sends post request to the api server to add new object
func post() {
	fmt.Println("Performing Http Post...")

	//Get content from the file
	stringYamlContent := utils.GetStringFromFile("deployment.yaml")

	kubeObject := KubeObject{ObjectName: "Deployment", YamlContent: stringYamlContent}

	//encode the data in json for sending the request
	jsonReq, e := json.Marshal(kubeObject)
	utils.CheckError(e)

	//sending post request to api server
	_, e = http.Post("http://localhost:10000/kubeObject", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	utils.CheckError(e)
}
