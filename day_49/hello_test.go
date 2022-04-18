package day_49

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello Lurr!"

	t.Logf("expect to return Hello Lurr!")

	if got != want {
		t.Errorf("go %q want %q", got, want)
	}
}
