package assert

import "testing"

// Equal _
func Equal(
	t *testing.T,
	expected interface{},
	actual interface{},
) {
	if actual != expected {
		t.Errorf("Assertion failed, %v ist not equal to %v", actual, expected)
	}
}
