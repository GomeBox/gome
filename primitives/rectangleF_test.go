package primitives

import "testing"

func TestNewRectangleF(t *testing.T) {
	x, y := float32(12.34), float32(356.903)
	width, height := float32(783.34), float32(22.0093)
	rect := NewRectangleF(x, y, width, height)
	if rect.X != x || rect.Y != y || rect.Width != width || rect.Height != height {
		t.Errorf("Rectangle has not expected values. Got x %v, y %v, w %v, h %v. Want x %v, y %v, w %v, h %v",
			rect.X, rect.Y, rect.Width, rect.Height,
			x, y, width, height)
	}
}
