package k8s_practice

import (
	"fmt"
)

func main() {
	fmt.Println("starting...")
	GetPod("prophet-ee-dev", "pas")
}
