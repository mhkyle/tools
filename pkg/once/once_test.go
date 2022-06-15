package once

import (
	"testing"
)

func TestOnce_Do(t *testing.T) {
	var one Once

	f1 := func() error {
		t.Logf("Do it one time")
		return nil
	}

	one.Do(f1)
	one.Do(f1)
}
