package main

import (
	"log"

	"github.com/xiatechs/jsonata-playbook/app"
)

func main() {
	err := app.Start()
	if err != nil {
		log.Fatal(err)
	}
}
