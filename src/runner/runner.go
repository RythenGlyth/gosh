package runner

import (
	"errors"
	"os"
	"strings"
)

// ErrNoEnvPath is returned if no path could be determined (e. g. $PATH is unset).
var ErrNoEnvPath error = errors.New("no $PATH variable") //nolint revive // They want to remove the error, which does not work.

// Runner finds the location of commands.
type Runner struct {
	path map[string]*string
}

// Init creates a new runner and reads all available executables from the $PATH.
// If $PATH is unset, ErrNoEnvPath is returned.
func Init() (*Runner, error) {
	var envPath string
	var ok bool

	envPath, ok = os.LookupEnv("PATH")

	if !ok {
		return nil, ErrNoEnvPath
	}

	r := Runner{make(map[string]*string)}

	for _, folder := range strings.Split(envPath, pathSep) {
		r.hashFolder(folder)
	}

	return &r, nil
}
