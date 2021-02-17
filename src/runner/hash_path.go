package runner

import "os"
import "strings"
import "sync"
import "fmt"

const (
	read_amount = 4096
)

var workers int = 0

// maps executable -> folder
var path map[string]*string = make(map[string]*string)

func Init() {
	var envPath string
	var ok bool

	var wg sync.WaitGroup
	var px sync.RWMutex

	envPath, ok = os.LookupEnv("PATH")
	if ok {
		for _, folder := range strings.Split(envPath, path_sep) {
			wg.Add(1)
			workers++
			go hashFolder(folder, &wg, &px)
		}
	}
	wg.Wait()
}

func hashFolder(root string, wg *sync.WaitGroup, px *sync.RWMutex) {
	var err error
	var f *os.File
	var entries []os.FileInfo
	var ok bool
	f, err = os.Open(root)
	if err != nil {
		fmt.Println(err)
	}
	var execs []string

	// for every entry set
	for err == nil {
		// TODO: go 1.16 provides f.ReadDir which should be faster
		entries, err = f.Readdir(read_amount)

		// For every file:
		for _, info := range entries {
			// Do we already have it?
			_, ok = path[info.Name()]
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
		path[exec] = &root
	}
	px.Unlock()

	wg.Done()
	workers--
	//fmt.Printf("%d to go\n", workers)
}
