package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Post() {
	fmt.Println("Performing Http Post...")
	kubeObject := KubeObject{ObjectName: "DaemonSet", YamlContent: "This is a daemonset"}
	jsonReq, err := json.Marshal(kubeObject)

	resp, err := http.Post("http://localhost:10000/kubeObject", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	// Convert response body to KubeObject struct
	var kubeObjectStruct KubeObject
	json.Unmarshal(bodyBytes, &kubeObjectStruct)
	fmt.Printf("%+v\n", kubeObjectStruct)
}
