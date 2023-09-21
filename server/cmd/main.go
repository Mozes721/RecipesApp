package main

import (
	"flag"

	"github.com/RecepieApp/server/app"
	"github.com/RecepieApp/server/runtime"
	// . "github.com/RecepieApp/server/runtime"
	//
	//"github.com/RecepieApp/server/db"
	//"github.com/RecepieApp/server/helpers"
)

var _ = flag.Bool("debug", false, "Enable Bun Debug log")

func main() {
	flag.Parse()
	client, err := app.LoadConfigurations()
	if err != nil {
		panic(err)
	}

	err = runtime.Start(client)
	if err != nil {
		panic(err)
	}
}
