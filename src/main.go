package main

import (
	"gosh/src/debug"
	"gosh/src/gosh"
	"log"
	"os"
)

func main() {
	myGosh := gosh.NewGosh()

	err := myGosh.Init()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "start-debug" {
			debug.StartDebugServer()
			return
		} else if os.Args[1] == "debug" {
			dc, err := debug.NewClient()
			if err != nil {
				log.Println(err)
				return
			}
			myGosh.SetDebugger(dc)
		} else if os.Args[1] == "plugin" {
			ds, err := debug.NewClient()
			if err != nil {
				log.Println(err)
				return
			}
			myGosh.SetDebugger(ds)

			ds.SendMessage(1, "Loading example-mod.so")
			myGosh.GetPluginManager().Load("/home/lenni/example-mod.so")
		}
	}

	retcode, err := myGosh.Interactive()
	if err != nil {
		os.Stdout.WriteString(err.Error())
	}

	os.Exit(retcode)
}
