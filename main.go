package main

import (
	"github.com/0x21F/inquizative/routes"
)

func main() {
	controller := routes.RouteController{}
	e := controller.ToServer()

	e.Start(":8080")

}
