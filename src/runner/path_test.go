package runner

import (
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	start := time.Now()
	r, _ := Init()
	elapsed := time.Since(start)

	t.Logf("Read %d executables in %d ms:\n", len(r.path), elapsed.Milliseconds())
	for exec, loc := range r.path {
		t.Log(*loc + " " + exec)
	}
}
