package game

import (
	"github.com/GomeBox/gome/primitives"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSettings_WindowSettings(t *testing.T) {
	settings := newSettings()
	assert.NotNil(t, settings)
	assert.NotNil(t, settings.WindowSettings())
	assert.Equal(t, settings.windowSettings.Resolution, primitives.Dimensions{
		Width:  800,
		Height: 600,
	})
}
