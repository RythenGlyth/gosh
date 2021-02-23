package main

import (
	"gosh/src/gosh"
	"time"

	"github.com/scrouthtv/termios"
)

func OnKey(g *gosh.Gosh, k *termios.Key) bool {
	return time.Now().Minute()%2 == 0
}
