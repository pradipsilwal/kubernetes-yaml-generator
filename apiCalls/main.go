package main

type KubeObject struct {
	ObjectName  string `json:"ObjectName"`
	YamlContent string `json:"YamlContent"`
}

func main() {
	delete("Deployment")
	// post()
}
