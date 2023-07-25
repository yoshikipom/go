package main

import (
	"fmt"

	"github.com/yoshikipom/json-masking-go/masking"
)

func main() {
	config := masking.MaskingConfig{
		DeniedKeyList: []string{"name"},
	}
	m := masking.New(&config)

	input := `{"key":"value","name":"John"}`
	output := m.Replace([]byte(input))
	fmt.Println(string(output))
}
