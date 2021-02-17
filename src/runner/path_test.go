package runner

import "time"
import "testing"

func TestInit(t *testing.T) {
	start := time.Now()
	r, _ := Init()
	elapsed := time.Now().Sub(start)

	t.Logf("Read %d executables in %d ms:\n", len(r.path), elapsed.Milliseconds())
	for exec, loc := range r.path {
		t.Log(*loc + " " + exec)
	}
}
