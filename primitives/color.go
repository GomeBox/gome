package primitives

type Color struct {
	R, G, B, A uint8
}

var definedColors *DefinedColors

func Colors() *DefinedColors {
	if definedColors == nil {
		definedColors = new(DefinedColors)
	}
	return definedColors
}

type DefinedColors struct {
}

func (_ DefinedColors) White() Color {
	return Color{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}
}
