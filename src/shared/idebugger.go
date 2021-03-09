package shared

// IDebugger implementations can be used to send debugging messages.
type IDebugger interface {
	SendMessage(ModuleIdentifier, string)
}
