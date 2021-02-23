package main

import "github.com/RythenGlyth/gosh/src/gosh"
import "github.com/scrouthtv/termios"

// OnKey returns false if the key is an 'x'.
func OnKey(g *gosh.Gosh, k *termios.Key) bool {
	return k.Value != 'x'
}
