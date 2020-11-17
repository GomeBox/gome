package mocks

import "github.com/GomeBox/gome/primitives"

type Drawer struct {
	OnDrawTo     func(pos primitives.PointF) error
	OnDimensions func() primitives.Dimensions
}

func (d *Drawer) DrawTo(pos primitives.PointF) error {
	if d.OnDrawTo != nil {
		return d.OnDrawTo(pos)
	}
	return nil
}

func (d *Drawer) Dimensions() primitives.Dimensions {
	if d.OnDimensions != nil {
		return d.OnDimensions()
	}
	return primitives.Dimensions{}
}
