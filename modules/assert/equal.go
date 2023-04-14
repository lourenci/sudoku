package assert

import (
	"testing"
)

func Equal[T comparable](t *testing.T, expected T, actual T) {
	t.Helper()

	if actual != expected {
		t.Fatalf("expected: %v; got it: %v", expected, actual)
	}
}
