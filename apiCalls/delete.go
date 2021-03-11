package main

import (
	"fmt"
	"net/http"

	"github.com/99-devops/kubernetes-yaml-generator/utils"
)

func delete(kubeObjectName string) {
	fmt.Println("Performing Http Delete...")
	request, e := http.NewRequest("DELETE", "http://localhost:10000/kubeObject/"+kubeObjectName, nil)
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	http.DefaultClient.Do(request)

	fmt.Println("Request: ", request)
	utils.CheckError(e)
	fmt.Println("Deleted")
}
