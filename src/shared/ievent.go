package shared

import "github.com/scrouthtv/termios"

type IEventHandler interface {
	PreSendKey(IGosh, *termios.Key) bool
	PostSendKey(IGosh, *termios.Key) bool
}
