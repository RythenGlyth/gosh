package main

import (
	"os"

	"gosh/gosh"
)

var MyGosh *gosh.Gosh

func main() {

	MyGosh = gosh.NewGosh()

	retcode, err := MyGosh.Interactive()
	if err != nil {
		os.Stdout.WriteString(err.Error())
	}

	os.Exit(retcode)
}
