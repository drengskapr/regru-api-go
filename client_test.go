package regru

import "testing"

func TestNew(t *testing.T) {
	c := New("user", "pass")
	if c == nil {
		t.Fatal("expected non-nil *Client")
	}
	if c.username != "user" {
		t.Errorf("expected username 'user', got %q", c.username)
	}
	if c.password != "pass" {
		t.Errorf("expected password 'pass', got %q", c.password)
	}
}
