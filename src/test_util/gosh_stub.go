package test_util

import (
	"gosh/src/shared"
	"testing"
)

type GoshStub struct {
	T *testing.T
}

func (g *GoshStub) ChangeWD(string) error                { return nil }
func (g *GoshStub) Close()                               {}
func (g *GoshStub) GetAliasManager() shared.AliasManager { return nil }

func (g *GoshStub) GetDebugger() shared.IDebugger                { return nil }
func (g *GoshStub) SetDebugger(shared.IDebugger)                 {}
func (g *GoshStub) DebugMessage(shared.ModuleIdentifier, string) {}

func (g *GoshStub) GetEventHandler() shared.IEventHandler   { return nil }
func (g *GoshStub) GetPluginManager() shared.IPluginManager { return nil }
func (g *GoshStub) GetWD() (string, error)                  { return "", nil }
func (g *GoshStub) Init() error                             { return nil }
func (g *GoshStub) Interactive() (int, error)               { return 0, nil }
func (g *GoshStub) RegisterBuiltin(shared.Builtin)          {}
func (g *GoshStub) Write(bytes []byte) (int, error)         { g.T.Log(string(bytes)); return len(bytes), nil }
func (g *GoshStub) WriteString(s string) (int, error)       { g.T.Log(s); return len(s), nil }

func (g *GoshStub) Size() (int, int) { return 80, 24 }
