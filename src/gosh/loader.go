package gosh

import (
	"plugin"
)

// Handler handles all event-based plugins.
// It is capable of loading new plugins from a specified path using the Load() method,
// as well as sending an event to all registered plugins
type Handler struct {
	keyListener []func(s string) bool
	parent      *Gosh
}

// NewHandler creates a new plugin handler with 0 plugins pre-loaded
// and the specified gosh as parent.
// g shall not be nil.
func NewHandler(g *Gosh) *Handler {
	return &Handler{nil, g}
}

// Load loads a plugin (shared library) from the specified path.
// It returns an error if the file could not be found.
func (h *Handler) Load(path string) error {
	p, err := plugin.Open(path)

	if err != nil {
		return err
	}

	loaded := h.loadKeyListeners(p)

	if loaded {
		h.parent.DebugMessage(2, "Loaded a key listener")
	} else {
		h.parent.DebugMessage(2, "Couldn't load a key listener")
	}

	return nil
}

// OnKey sends the event to all plugins and
// returns false as soon as the first plugin returns false
// or true if all loaded plugins returned true.
func (h *Handler) OnKey(s string) bool {
	var ok bool = true

	for _, f := range h.keyListener {
		ok = f(s)
		if !ok {
			return false
		}
	}

	return true
}

func (h *Handler) loadKeyListeners(p *plugin.Plugin) bool {
	s, err := p.Lookup("OnKey")

	if err != nil {
		// OnKey() does not exist
		h.parent.DebugMessage(2, "Couldn't find the OnKey() method")
		return false
	}

	f, ok := s.(func(s string) bool)
	if !ok {
		h.parent.DebugMessage(2, "OnKey() has wrong signature or isn't a function")
		return false
	}

	h.keyListener = append(h.keyListener, f)

	h.parent.DebugMessage(2, "Succesfully loaded a key listener")
	return true
}
