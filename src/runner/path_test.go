package runner

import (
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	start := time.Now()

	r, err := Init()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	elapsed := time.Since(start)

	t.Logf("Read %d executables in %d ms:\n", len(r.path), elapsed.Milliseconds())
}
