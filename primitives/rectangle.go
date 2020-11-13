package primitives

//Rectangle defines a rectangle using whole numbers
type Rectangle struct {
	Point
	Dimensions
}

func NewRectangle(x, y, width, height int) Rectangle {
	return Rectangle{
		Point: Point{
			X: x,
			Y: y,
		},
		Dimensions: Dimensions{
			Width:  width,
			Height: height,
		},
	}
}
