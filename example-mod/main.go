package main

import (
	"gosh/src/gosh"
	"time"

	"github.com/scrouthtv/termios"
)

// OnKey only returns true on even minutes.
func OnKey(g *gosh.Gosh, k *termios.Key) bool {
	return time.Now().Minute()%2 == 0
}
