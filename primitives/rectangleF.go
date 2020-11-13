package primitives

//Rectangle defines a rectangle using floating point numbers
type RectangleF struct {
	PointF
	DimensionsF
}

func NewRectangleF(x, y, width, height float32) RectangleF {
	return RectangleF{
		PointF: PointF{
			X: x,
			Y: y,
		},
		DimensionsF: DimensionsF{
			Width:  width,
			Height: height,
		},
	}
}
