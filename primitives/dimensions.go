package primitives

//Dimensions defines a width and a height of an entity
type Dimensions struct {
	Width, Height int
}

//ToDimensionsF returns a new DimensionsF instance with the Width and Height from the current instance cast to float32
func (dimensions *Dimensions) ToDimensionsF() DimensionsF {
	return DimensionsF{
		Width:  float32(dimensions.Width),
		Height: float32(dimensions.Height),
	}
}
