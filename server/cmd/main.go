package main

import (
	"flag"
	. "github.com/RecepieApp/server/runtime"
	//
	//"github.com/RecepieApp/server/db"
	//"github.com/RecepieApp/server/helpers"
)

var _ = flag.Bool("debug", false, "Enable Bun Debug log")

func main() {
	flag.Parse()
	//// Initialize the runtime
	//_, err := a.LoadConfigurations()
	//fmt.Println(a.Client)
	//if err != nil {
	//	panic(err)
	//}
	err := Start()
	if err != nil {
		panic(err)
	}
}
