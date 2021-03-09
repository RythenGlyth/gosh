package shared

// IGosh describes the main type of a gosh instance.
// It collects all modules of a gosh shell and
// binds functionality to each module.
type IGosh interface {
	SetDebugger(IDebugger)
	GetDebugger() IDebugger
	DebugMessage(ModuleIdentifier, string)

	GetPluginManager() IPluginManager

	GetEventHandler() IEventHandler
}
