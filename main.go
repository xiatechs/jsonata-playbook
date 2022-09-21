package main

import (
	"github.com/xiatechs/jsonata-playbook/app"
)

func main() {
	app.SetEndpoint(":8050")
	app.Start()
}
