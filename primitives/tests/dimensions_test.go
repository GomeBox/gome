package tests

import (
	"github.com/GomeBox/gome/primitives"
	"math"
	"testing"
)

func TestDimensions_ToDimensionsF(t *testing.T) {
	wantWidth := float32(123.0)
	wantHeight := float32(234.0)
	dimensions := primitives.Dimensions{
		Width:  int(wantWidth),
		Height: int(wantHeight),
	}
	dimensionsF := dimensions.ToDimensionsF()
	if !AreEqual(wantWidth, dimensionsF.Width) {
		t.Errorf("got: %f but wanted: %f for Width", dimensionsF.Width, wantWidth)
	}
	if !AreEqual(wantHeight, dimensionsF.Height) {
		t.Errorf("got: %f but wanted: %f for Height", dimensionsF.Height, wantHeight)
	}
}

type testParams struct {
	f1, f2         float32
	expectedResult bool
}

func TestAreEqual(t *testing.T) {
	testParams := []testParams{
		{
			f1:             10.1234,
			f2:             10.1234,
			expectedResult: true,
		},
		{
			f1:             10.1235,
			f2:             10.1234,
			expectedResult: false,
		},
		{
			f1:             -10.1234,
			f2:             10.1234,
			expectedResult: false,
		},
		{
			f1:             10.1234,
			f2:             -10.1234,
			expectedResult: false,
		},
		{
			f1:             0,
			f2:             0,
			expectedResult: true,
		},
		{
			f1:             100,
			f2:             100,
			expectedResult: true,
		},
	}
	for _, params := range testParams {
		got := AreEqual(params.f1, params.f2)
		want := params.expectedResult
		if got != want {
			t.Errorf("got: %t but wanted: %t for f1: %f, f2: %f", got, want, params.f1, params.f2)
		}
	}
}

func AreEqual(f1, f2 float32) bool {
	return math.Abs(float64(f1-f2)) <= 0.00001
}
