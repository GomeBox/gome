package shared

import (
	"errors"
	"github.com/GomeBox/gome/adapters/shared/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnloader_Unload(t *testing.T) {
	unloadCalled := false
	retErr := false
	adapterMock := &mocks.Unloader{
		OnUnload: func() error {
			unloadCalled = true
			if retErr {
				return errors.New("test")
			}
			return nil
		},
	}
	unloader := Unloader{Adapter: adapterMock}
	err := unloader.Unload()
	assert.NoError(t, err)
	assert.True(t, unloadCalled)
	retErr = true
	err = unloader.Unload()
	assert.Error(t, err)
	retErr = false
	err = unloader.Unload()
	assert.Error(t, err, "already unloaded. Should return error")
}

func TestUnloader_Unloaded(t *testing.T) {
	retErr := false
	adapterMock := &mocks.Unloader{
		OnUnload: func() error {
			if retErr {
				return errors.New("test")
			}
			return nil
		},
	}
	unloader := Unloader{Adapter: adapterMock}
	assert.False(t, unloader.Unloaded())
	_ = unloader.Unload()
	assert.True(t, unloader.Unloaded())

	unloader = Unloader{Adapter: adapterMock}
	retErr = true
	_ = unloader.Unload()
	assert.False(t, unloader.Unloaded())
}
