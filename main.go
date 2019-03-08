package main

import (
	"fmt"
	"github.com/onesafe/k8s_practice/resource"
)

func main() {
	fmt.Println("starting...")
	pasPod, err := resource.GetDeployment("prophet-ee-dev", "pas")
	if err != nil {
		fmt.Println("Get Deployment error: ", err.Error())
	} else {
		fmt.Println(pasPod)
	}
}
