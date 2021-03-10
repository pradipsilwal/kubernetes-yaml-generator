package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/99-devops/kubernetes-yamal-generator/utils"
)

func post() {
	fmt.Println("Performing Http Post...")

	stringYamlContent := utils.GetStringFromFile("deployment.yaml")

	kubeObject := KubeObject{ObjectName: "Deployment", YamlContent: stringYamlContent}

	jsonReq, e := json.Marshal(kubeObject)
	utils.CheckError(e)

	_, e = http.Post("http://localhost:10000/kubeObject", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	utils.CheckError(e)
}
