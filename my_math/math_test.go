package my_math

import "testing"

func TestSum(t *testing.T) {
	got := Sum(4, 6)
	want := 101

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
