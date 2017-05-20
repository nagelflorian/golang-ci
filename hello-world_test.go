package main

import "testing"

func TestGreeting(t *testing.T) {
	cases := []struct {
		in, expect string
	}{
		{"Felix", "Hi, Felix"},
		{"Regina", "Hi, Regina"},
	}

	for _, c := range cases {
		got := Greeting(c.in)
		if got != c.expect {
			t.Errorf("Greeting(%q) == %q, expect %q", c.in, got, c.expect)
		}
	}
}
