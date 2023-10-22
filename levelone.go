package golang_dependencies

import "fmt"

const (
	G float32 = 9.8
)

func Level() string {
	return "String from direct dependency"
}

func Direct() {
	fmt.Println("Hello from direct dependency func")
}
