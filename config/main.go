package main

import (
	"fmt"

	"github.com/yoshikipom/go/config/internal/config"
)

func main() {
	// call only once
	err := config.Initialize("./config/config.yml")
	if err != nil {
		panic(err)
	}

	// you can get config anywhere
	c := config.GetConfig()
	fmt.Printf("%+v", c)
}
