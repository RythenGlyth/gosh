package main

import (
	"log"
	"os"

	"github.com/RythenGlyth/gosh/src/debug"
	"github.com/RythenGlyth/gosh/src/gosh"
)

var myGosh *gosh.Gosh

func main() {

	myGosh = gosh.NewGosh()
	myGosh.Init()

	if len(os.Args) > 1 {
		if os.Args[1] == "start-debug" {
			debug.StartDebugServer()
			return
		} else if os.Args[1] == "debug" {
			debug, err := debug.NewClient()
			if err != nil {
				log.Println(err)
				return
			}
			myGosh.SetDebugClient(debug)
		} else if os.Args[1] == "plugin" {
			debug, err := debug.NewClient()
			if err != nil {
				log.Println(err)
				return
			}
			myGosh.SetDebugClient(debug)

			myGosh.LoadPlugin("/home/lenni/example-mod.so")
		}
	}

	myGosh.SendKey("asdf")
	retcode, err := myGosh.Interactive()
	if err != nil {
		os.Stdout.WriteString(err.Error())
	}

	os.Exit(retcode)
}
