package mxchecker

import "testing"

func TestVerifyEmail(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"gokul@tutecircle.com", "valid"},
		{"tp.gokul@tutecircle.com", "invalid"},
		{"gokul@yesware.com", "accept_all"},
	}
	for _, c := range cases {
		got, err  := VerifyEmail(c.in)
		if err != nil {
			t.Errorf("Thrown unhandled exception %q",  err)
		}
		if got != c.want {
			t.Errorf("VerifyEmail(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
