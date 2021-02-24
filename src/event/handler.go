package event

import (
	"gosh/src/shared"

	"github.com/scrouthtv/termios"
)

type EventHandler struct {
	parent               shared.IGosh
	preSendKeyListeners  []func(g shared.IGosh, k *termios.Key) bool
	postSendKeyListeners []func(g shared.IGosh, k *termios.Key) bool
}

func NewEventHandler(parent shared.IGosh) *EventHandler {
	return &EventHandler{parent, nil, nil}
}

func (h *EventHandler) PreSendKey(g shared.IGosh, k *termios.Key) bool {
	var ok bool = true

	for _, f := range h.preSendKeyListeners {
		ok = f(h.parent, k)

		if !ok {
			h.parent.DebugMessage(shared.ModPluginLoader, "Key event was cancelled by a plugin")
			return false
		}
	}

	return true
}

func (h *EventHandler) PostSendKey(g shared.IGosh, k *termios.Key) bool {
	var ok bool = true

	for _, f := range h.postSendKeyListeners {
		ok = f(h.parent, k)

		if !ok {
			h.parent.DebugMessage(shared.ModPluginLoader, "Key event was cancelled by a plugin")
			return false
		}
	}

	return true
}
