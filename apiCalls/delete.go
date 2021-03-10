package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Delete() {
	fmt.Println("Performing Http Delete...")
	kubeObject := KubeObject{"Deployment", "Af"}
	jsonReq, err := json.Marshal(kubeObject)
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:10000/kubeObject/"+kubeObject.ObjectName, bytes.NewBuffer(jsonReq))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}
