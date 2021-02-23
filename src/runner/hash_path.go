package runner

import (
	"os"
	"sync"
)

const (
	readAmount = 4096
)

func (r *Runner) hashFolder(root string, wg *sync.WaitGroup, px *sync.RWMutex) {
	var err error
	var f *os.File
	var entries []os.FileInfo
	var ok bool
	f, err = os.Open(root)
	//if err != nil {
	// TODO debug error
	//}
	var execs []string

	// for every entry set
	for err == nil {
		// TODO: go 1.16 provides f.ReadDir which should be faster
		entries, err = f.Readdir(readAmount)

		// For every file:
		for _, info := range entries {
			// Do we already have it?
			_, ok = r.path[info.Name()]
			if !ok {
				// Is it an executable?
				if isExecutable(info) {
					execs = append(execs, info.Name())
				}
			}
		}
	}

	px.Lock()
	for _, exec := range execs {
		r.path[exec] = &root
	}
	px.Unlock()

	wg.Done()
}
