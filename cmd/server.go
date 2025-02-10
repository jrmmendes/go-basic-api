package main

import (
	"api/pkg/core"
)

func main() {
	app := core.Bootstrap()
	app.Listen(":3000")
}
