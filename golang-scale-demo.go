package main

// This is a simple Golang Martini Application that is meant to be demoed with
// the Pivotal Web Service App Manager. Information on Martini can be found
// at http://martini.codegangsta.io/
//
// The Application shows the Instance ID and address to demonstrate the load
// balance capability plus the ability to terminate and instance and watch the
// recovery in the apps manager
//
// Charles Wu (chwu@pivotal.io)
//
import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"os"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		instance := os.Getenv("CF_INSTANCE_INDEX")
		instance_addr := os.Getenv("CF_INSTANCE_ADDR")
		data := struct {
			Instance      string
			Instance_addr string
		}{
			instance,
			instance_addr}
		r.HTML(200, "home", data)
	})

	m.Get("/kill", func() {
		os.Exit(9)
	})

	m.Run()
}
