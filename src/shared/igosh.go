package shared

// IGosh describes the main type of a gosh instance.
// It collects all modules of a gosh shell and
// binds functionality to each module.
type IGosh interface {
	WriteString(string) (n int, err error)
	Init() error

	// Interactive returns a return code and the last error.
	Interactive() (n int, err error)
	Close()

	RegisterBuiltin(b Builtin)

	// ChangeWD changes the current directory to the specified one.
	ChangeWD(string) error
	// GetWD returns the current working directory.
	GetWD() (string, error)

	GetAliasManager() AliasManager

	SetDebugger(IDebugger)
	GetDebugger() IDebugger
	DebugMessage(ModuleIdentifier, string)

	GetPluginManager() IPluginManager

	GetEventHandler() IEventHandler
}

type Builtin interface {
	Match(line string) bool
	Eval(g IGosh, line string) error
}
