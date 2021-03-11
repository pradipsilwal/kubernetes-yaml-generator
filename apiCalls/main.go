package main

type KubeObject struct {
	ObjectName  string `json:"ObjectName"`
	YamlContent string `json:"YamlContent"`
}

func main() {

	//performing delete request to the delete api
	delete("Deployment")

	//sending post request to post api
	// post()
}
