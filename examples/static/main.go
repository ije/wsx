package main

import (
	"github.com/ije/rex"
)

func main() {
	rex.Query("*", func(ctx *rex.Context) interface{} {
		return rex.Static("./www", "e404.html")
	})

	rex.Use(rex.Debug())
	rex.Start(8080)
}
