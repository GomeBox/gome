package graphics

import (
	"github.com/GomeBox/gome/internal/graphics/interfaces"
	"github.com/GomeBox/gome/primitives"
)

//Text represents a drawable Text
type Text interface {
	//Draw draws the text to the passed position of the game window
	Draw(position *primitives.Point) error
	//DrawF draws the text to the passed position of the game window
	DrawF(position *primitives.PointF) error
	//Dimensions returns the size of the text in pixels
	Dimensions() primitives.Dimensions
}

type text struct {
	internalText interfaces.Text
}

func (t *text) Draw(position *primitives.Point) error {
	return t.internalText.Draw(position)
}

func (t *text) DrawF(position *primitives.PointF) error {
	return t.internalText.DrawF(position)
}

func (t *text) Dimensions() primitives.Dimensions {
	return t.internalText.Dimensions()
}
