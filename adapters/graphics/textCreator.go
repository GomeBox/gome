package graphics

import "github.com/GomeBox/gome/primitives"

//TextCreator is the adapter that is returned by FontLoader.
//It is used to create a text with the loaded font
type TextCreator interface {
	//Create returns a drawable text
	Create(text string, color primitives.Color) (TextDrawer, error)
}
