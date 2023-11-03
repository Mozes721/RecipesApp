package main

import (
	"flag"

	"github.com/RecepieApp/server/app"
	"github.com/RecepieApp/server/runtime"
)

var _ = flag.Bool("debug", false, "Enable Bun Debug log")

func main() {
	flag.Parse()

	a := app.Application{}

	err := a.LoadConfigurations()
	if err != nil {
		panic(err)
	}

	err = runtime.Start(&a)
	if err != nil {
		panic(err)
	}
}
