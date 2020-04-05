package dev

import (
	"github.com/cusspvz/ble"
	"github.com/cusspvz/ble/darwin"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return darwin.NewDevice(opts...)
}
