package main

import (
	"app"
)

func hello(val string) string {
	return "This is test"
}
func main() {
	app.Run(":6070")
	app.Get("/", hello)
}
