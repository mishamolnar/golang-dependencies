package golang_dependencies

import dp "github.com/mishamolnar/golang-dependency-two"

const (
	G float32 = 9.8
)

func Level() string {
	return dp.CapitalizeString("String from direct dependency", "another")
}
