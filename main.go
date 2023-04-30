package main

import (
	"log"
	"os"

	"github.com/zlucasbiasi/searchIP/app"
)

func main() {

	aplication := app.Generate()

	if err := aplication.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
