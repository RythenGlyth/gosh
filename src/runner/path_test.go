package runner

import "time"
import "testing"

func TestInit(t *testing.T) {
	start := time.Now()
	Init()
	elapsed := time.Now().Sub(start)

	t.Logf("Read %d executables in %d ms", len(path), elapsed.Milliseconds())
}
