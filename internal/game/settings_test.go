package game

import (
	"github.com/GomeBox/gome/primitives"
	"testing"
)

func TestSettings_NewSettings(t *testing.T) {
	settings := NewSettings()
	if settings == nil {
		t.Fatal("NewSettings() returned nil")
	}
	windowSettings := settings.WindowSettings()
	if windowSettings == nil {
		t.Fatal("WindowSettings() returned nil")
	}
}

func TestSettings_WindowSettingsResolution(t *testing.T) {
	windowSettings := NewSettings().WindowSettings()
	want := primitives.Dimensions{
		Width:  800,
		Height: 600,
	}
	got := windowSettings.Resolution
	if got != want {
		t.Errorf("resolution default values not as expected. Got %v, want %v", got, want)
	}
}

func TestSettings_WindowSettingsFullscreen(t *testing.T) {
	windowSettings := NewSettings().WindowSettings()
	want := false
	got := windowSettings.Fullscreen
	if got != want {
		t.Errorf("fullscreen default value not as expected. Got %t, want %t", got, want)
	}
}

func TestSettings_WindowSettingsTitle(t *testing.T) {
	windowSettings := NewSettings().WindowSettings()
	want := ""
	got := windowSettings.Title
	if got != want {
		t.Errorf("title default value not as expected. Got %q, want %q", got, want)
	}
}
