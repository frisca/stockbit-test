package main

import (
	"stockbit-test/omdb/routes"
)

func main() {
	//setup routes
	r := routes.SetupRouter()

	// running
	r.Run()

}
