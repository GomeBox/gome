package primitives

//Dimensions defines a width and a height of an entity
type Dimensions struct {
	Width, Height int
}

func (dimensions *Dimensions) ToDimensionsF() DimensionsF {
	return DimensionsF{
		Width:  float32(dimensions.Width),
		Height: float32(dimensions.Height),
	}
}
