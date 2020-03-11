package main

import (
	"strings"

	"github.com/ije/rex"
)

const indexHTML = `
	<h1>Welcome to use REX!</h1>
	<p><a href="/users/bob">User Bob</a></p>
	<p><a href="/v2">V2 API</a></p>
	<p><a href="/v3">V3 API</a></p>
`
const v2HTML = `
	<h1>V2 API</h1>
	<p><a href="/v2/users/bob">User Bob</a></p> 
	<p><a href="/">Home</a></p>
`
const v3HTML = `
	<h1>V3 API</h1>
	<p><a href="/v3/users/bob">User Bob</a></p> 
	<p><a href="/">Home</a></p>
`

func main() {
	rex.Use(rex.Header("X-Version", "default"), rex.Header("Foo", "bar"))

	rex.Get("/", func(ctx *rex.Context) {
		ctx.HTML(indexHTML)
	})

	v1 := rex.New()
	v1.Use(rex.Header("X-Version", "v1"))
	v1.Group("/users", func(r *rex.REST) {
		r.Get("/:id", func(ctx *rex.Context) {
			ctx.Ok("Hello, I'm " + strings.Title(ctx.URL.Param("id")) + "!")
		})
	})

	v2 := rex.New("v2")
	v2.Use(rex.Header("X-Version", "v2"))
	v2.Get("/", func(ctx *rex.Context) {
		ctx.HTML(v2HTML)
	})
	v2.Group("/users", func(r *rex.REST) {
		r.Get("/:id", func(ctx *rex.Context) {
			ctx.Ok("[v2] Hello, I'm " + strings.Title(ctx.URL.Param("id")) + "!")
		})
	})

	v3 := rex.New("v3")
	v3.Use(rex.Header("X-Version", "v3"))
	v3.Get("/", func(ctx *rex.Context) {
		ctx.HTML(v3HTML)
	})
	v3.Group("/users", func(r *rex.REST) {
		r.Get("/:id", func(ctx *rex.Context) {
			ctx.Ok("[v3] Hello, I'm " + strings.Title(ctx.URL.Param("id")) + "!")
		})
	})

	rex.Start(8080)
}
