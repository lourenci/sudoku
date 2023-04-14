package assert

import (
	"testing"
)

func Equals[T comparable](t *testing.T, actual T, expected T) {
	t.Helper()

	if actual != expected {
		t.Fatalf(`expected: "%v"; got it: "%v"`, expected, actual)
	}
}

func PanicsWithMessage(t *testing.T, fn func(), msg string) {
	t.Helper()

	defer func() {
		r := recover()

		if r == nil {
			t.Fatalf("expected to panic, but it did not")
		}
		if r != msg {
			t.Fatalf(
				`expected to panic with the message "%s", but it panicked with the message "%s"`,
				msg,
				r,
			)
		}
	}()

	fn()
}
