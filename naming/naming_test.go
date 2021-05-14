package naming

import "testing"

func TestColor(t *testing.T) {
	arg := "blue"
	want := "#0000FE"
	got := Color(arg)
	if got != want {
		t.Errorf("Color(%q) = %s; want %s", arg, got, want)
	}
}
