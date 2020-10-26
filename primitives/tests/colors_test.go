package tests

import (
	"github.com/GomeBox/gome/primitives"
	"testing"
)

type colorTestParams struct {
	colorName     string
	testFunction  func() primitives.Color
	expectedColor primitives.Color
}

func TestColors(t *testing.T) {
	testParams := []colorTestParams{
		{
			colorName:    "White",
			testFunction: primitives.Colors().White,
			expectedColor: primitives.Color{
				R: 255,
				G: 255,
				B: 255,
				A: 255,
			},
		},
	}
	for _, c := range testParams {
		got := c.testFunction()
		want := c.expectedColor
		if got != want {
			t.Errorf("Color %q (%v) not matching expected (%v)", c.colorName, got, want)
		}
	}
}
