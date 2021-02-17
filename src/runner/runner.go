package runner

import "os"
import "errors"
import "sync"
import "strings"

var NoEnvPath error = errors.New("No $PATH variable")

type Runner struct {
	path map[string]*string
}

func Init() (*Runner, error) {
	var envPath string
	var ok bool

	var wg sync.WaitGroup
	var px sync.RWMutex

	envPath, ok = os.LookupEnv("PATH")
	if !ok {
		return nil, NoEnvPath
	}

	var r Runner = Runner{make(map[string]*string)}

	for _, folder := range strings.Split(envPath, path_sep) {
		wg.Add(1)
		go r.hashFolder(folder, &wg, &px)
	}
	wg.Wait()

	return &r, nil
}
