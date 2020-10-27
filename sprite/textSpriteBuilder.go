package sprite

import (
	"github.com/GomeBox/gome/graphics"
	"github.com/GomeBox/gome/internal/sprite"
	"github.com/GomeBox/gome/primitives"
)

//TextSpriteBuilder contains functions to create a sprite containing Text
type TextSpriteBuilder interface {
	//Get creates the sprite
	Get() Interface
	//SetPosition sets the initial position that the sprite will have after it is created
	SetPosition(x, y float32) TextSpriteBuilder
}

//FromText returns a TextSpriteBuilder that can be used to create a sprite
func FromText(text graphics.Text) TextSpriteBuilder {
	return &textSpriteBuilder{
		text: text,
		pos:  primitives.PointF{},
	}
}

type textSpriteBuilder struct {
	text graphics.Text
	pos  primitives.PointF
}

func (builder *textSpriteBuilder) Get() Interface {
	drawer := sprite.NewTextDrawer(builder.text)
	s := sprite.New(drawer, builder.pos)
	return s
}

func (builder *textSpriteBuilder) SetPosition(x, y float32) TextSpriteBuilder {
	builder.pos.X = x
	builder.pos.Y = y
	return builder
}
