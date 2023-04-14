package assert

import (
	"testing"
)

func Equal[T comparable](t *testing.T, actual T, expected T) {
	t.Helper()

	if actual != expected {
		t.Fatalf("expected: %v; got it: %v", expected, actual)
	}
}
