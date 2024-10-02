package main

import (
	"flag"
	"log" 
	"github.com/RecepieApp/server/app"
	"github.com/RecepieApp/server/runtime"
	
)

var _ = flag.Bool("debug", false, "Enable Bun Debug log")

func main() {
	flag.Parse()

	a := app.Application{}

	if err := a.LoadConfigurations(); err != nil {
        log.Fatalf("Failed to load configurations: %v", err)
    }

    if err := runtime.Start(&a); err != nil {
        log.Fatalf("Failed to start the application: %v", err)
    }

}

