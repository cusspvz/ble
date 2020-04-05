package dev

import (
	"github.com/cusspvz/ble"
	"github.com/cusspvz/ble/linux"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return linux.NewDevice(opts...)
}
