package main

import (
	"github.com/xiatechs/jsonata-playbook/app"
)

func main() {
	app.Endpoint = "127.0.0.1:7085"
	app.Start()
}
