package graphics

import "github.com/GomeBox/gome/primitives"

//TextDrawer is the adapter that is used to draw text
type TextDrawer interface {
	//Draw renders the text in its original size to the defined position
	Draw(position *primitives.Point) error
	//DrawScaled draws the text into the defined destRect
	DrawScaled(destRect *primitives.Rectangle) error
	//DrawF renders the text in its original size to the defined position
	DrawF(position *primitives.PointF) error
	//DrawScaledF draws the text into the defined destRect
	DrawScaledF(destRect *primitives.RectangleF) error
}
