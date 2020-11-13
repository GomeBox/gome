package input

import (
	"github.com/GomeBox/gome/primitives"
	"testing"
)

func TestKey_KeyType(t *testing.T) {
	want := primitives.KeyT
	key := key{keyType: want}
	got := key.KeyType()
	if want != got {
		t.Errorf("KeyType() returned unexpected value. Got %v, want %v", got, want)
	}
}

func TestKey_IsPressed(t *testing.T) {
	key := new(key)
	cases := []bool{true, false}
	for _, want := range cases {
		key.isPressed = want
		got := key.IsPressed()
		if got != want {
			t.Errorf("IsPressed() returned unexpected value. Got %t, want %t", got, want)
		}
	}
}

func TestKey_WasPressed(t *testing.T) {
	key := new(key)
	cases := []bool{true, false}
	for _, want := range cases {
		key.wasPressed = want
		got := key.WasPressed()
		if got != want {
			t.Errorf("WasPressed() returned unexpected value. Got %t, want %t", got, want)
		}
	}
}
