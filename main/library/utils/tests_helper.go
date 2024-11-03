package utils

import "testing"

func AssertTrue(t *testing.T, condition bool, message string) {
	if !condition {
		t.Error(message)
	}
}
