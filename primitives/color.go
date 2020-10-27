package primitives

import "fmt"

//Color is a RGBA-Color
type Color struct {
	R, G, B, A uint8
}

func (color Color) String() string {
	return fmt.Sprintf("R:%d, G:%d, B:%d, A:%d", color.R, color.G, color.B, color.A)
}

var definedColors *DefinedColors

//Colors contains predefined colors
func Colors() *DefinedColors {
	if definedColors == nil {
		definedColors = new(DefinedColors)
	}
	return definedColors
}

type DefinedColors struct {
}

//White returns the Color white
func (_ DefinedColors) White() Color {
	return Color{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}
}
