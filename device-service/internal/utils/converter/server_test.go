package converter

import (
	"github.com/alserov/device-shop/proto/gen/device"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAdminConverter(t *testing.T) {
	c := NewServerConverter()

	// CREATE DEVICE
	createDevice := &device.CreateDeviceReq{
		Title:        "title",
		Description:  "desc",
		Price:        100,
		Manufacturer: "manu",
		Amount:       5,
	}
	convertedCreateDevice := c.Admin.CreateDeviceToService(createDevice)

	require.Equal(t, createDevice.Title, convertedCreateDevice.Title)
	require.Equal(t, createDevice.Amount, convertedCreateDevice.Amount)
	require.Equal(t, createDevice.Description, convertedCreateDevice.Description)
	require.Equal(t, createDevice.Manufacturer, convertedCreateDevice.Manufacturer)
	require.Equal(t, createDevice.Price, convertedCreateDevice.Price)

	// UPDATE DEVICE
	updateDevice := &device.UpdateDeviceReq{
		UUID:        "uuid",
		Title:       "title",
		Description: "desc",
		Price:       100,
	}
	convertedUpdateDevice := c.Admin.UpdateDeviceToService(updateDevice)

	require.Equal(t, updateDevice.UUID, convertedUpdateDevice.UUID)
	require.Equal(t, updateDevice.Title, convertedUpdateDevice.Title)
	require.Equal(t, updateDevice.Description, convertedUpdateDevice.Description)
	require.Equal(t, updateDevice.Price, convertedUpdateDevice.Price)

}
