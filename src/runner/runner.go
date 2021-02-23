package runner

import (
	"errors"
	"os"
	"strings"
	"sync"
)

// ErrNoEnvPath is returned if no path could be determined (e. g. $PATH is unset)
var ErrNoEnvPath error = errors.New("no $PATH variable")

// Runner finds the location of commands.
type Runner struct {
	path map[string]*string
}

// Init creates a new runner and reads all available executables from the $PATH.
// If $PATH is unset, ErrNoEnvPath is returned.
func Init() (*Runner, error) {
	var envPath string
	var ok bool

	var wg sync.WaitGroup
	var px sync.RWMutex

	envPath, ok = os.LookupEnv("PATH")

	if !ok {
		return nil, ErrNoEnvPath
	}

	var r Runner = Runner{make(map[string]*string)}

	for _, folder := range strings.Split(envPath, pathSep) {
		wg.Add(1)
		go r.hashFolder(folder, &wg, &px)
	}

	wg.Wait()

	return &r, nil
}
