package main

import (
	"fmt"

	"github.com/99-devops/kubernetes-yamal-generator/utils"
)

type KubeObject struct {
	ObjectName  string `json:"ObjectName"`
	YamlContent string `json:"YamlContent"`
}

func main() {
	// delete()
	// post()
	lines, e := utils.GetStrings("pod.yaml")
	printLines(lines)
	utils.CheckError(e)
	fmt.Println(lines)
}

func printLines(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}
