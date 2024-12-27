package main

import (
	"fmt"

	validator "github.com/matheus-gondim/omega-validator"
)

func main() {
	b, e := validator.Compose(
		validator.New("campo_0", "06976741690").FederalDocument(),
		validator.New("campo_1", 12).Min(1),
		validator.New("campo_2", "1").Min(3).Max(5).Contains("1"),
		validator.New("campo_3", "nãotem@nãotem.com").Required().Email(),
	)
	if e != nil {
		fmt.Println(b, e.Errors)
	}

	fmt.Println(b)
}
