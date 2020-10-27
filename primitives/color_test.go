package primitives

import (
	"testing"
)

type colorTestParams struct {
	colorName     string
	testFunction  func() Color
	expectedColor Color
}

func TestColors(t *testing.T) {
	testParams := []colorTestParams{
		{
			colorName:    "White",
			testFunction: Colors().White,
			expectedColor: Color{
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
