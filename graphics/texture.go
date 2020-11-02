package graphics

import (
	"github.com/GomeBox/gome/internal/graphics/interfaces"
	"github.com/GomeBox/gome/primitives"
)

//Texture represents a drawable Texture
type Texture interface {
	//Draw draws the part of the texture, defined by the source-rectangle to the part of the game window,
	//that is defined by the dest-rectangle. If source is nil, the whole texture is drawn
	Draw(source, dest *primitives.Rectangle) error
	//DrawF draws the part of the texture, defined by the source-rectangle to the part of the game window,
	//that is defined by the dest-rectangle. If source is nil, the whole texture is drawn
	DrawF(source *primitives.Rectangle, dest *primitives.RectangleF) error
	//Dimensions returns the textures dimensions
	Dimensions() primitives.Dimensions
}

type texture struct {
	internalTexture interfaces.Texture
}

func (t *texture) Draw(source, dest *primitives.Rectangle) error {
	return t.internalTexture.Draw(source, dest)
}

func (t *texture) DrawF(source *primitives.Rectangle, dest *primitives.RectangleF) error {
	return t.internalTexture.DrawF(source, dest)
}

func (t *texture) Dimensions() primitives.Dimensions {
	return t.internalTexture.Dimensions()
}
