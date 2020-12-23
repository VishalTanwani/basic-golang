package main

import (

	// api "mini/api"
	// cli "mini/cli"
	_ "mini/db"
	dump "mini/dump"
)

type test struct {
	Name string
	Id   int
}

func main() {
	// cli.Run()
	dump.Run()
	// api.Run()
}
