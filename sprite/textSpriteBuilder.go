package sprite

import (
	"github.com/GomeBox/gome/graphics"
	"github.com/GomeBox/gome/primitives"
)

type TextSpriteBuilder interface {
	Get() Interface
	SetPosition(x, y float32) TextSpriteBuilder
}

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
	drawer := newTextDrawer(builder.text)
	sprite := sprite{
		drawer:   drawer,
		position: builder.pos,
	}
	return &sprite
}

func (builder *textSpriteBuilder) SetPosition(x, y float32) TextSpriteBuilder {
	builder.pos.X = x
	builder.pos.Y = y
	return builder
}
