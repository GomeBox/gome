package primitives

import "testing"

func TestNewRectangle(t *testing.T) {
	x, y := 100, 234
	width, height := 23, 799
	rect := NewRectangle(x, y, width, height)
	if rect.X != x || rect.Y != y || rect.Width != width || rect.Height != height {
		t.Errorf("Rectangle has not expected values. Got x %v, y %v, w %v, h %v. Want x %v, y %v, w %v, h %v",
			rect.X, rect.Y, rect.Width, rect.Height,
			x, y, width, height)
	}
}
