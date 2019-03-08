package main

import (
	"fmt"
	"github.com/onesafe/k8s_practice/pod"
)

func main() {
	fmt.Println("starting...")
	pasPod := pod.GetPod("prophet-ee-dev", "pas")
	fmt.Println(pasPod)
}
