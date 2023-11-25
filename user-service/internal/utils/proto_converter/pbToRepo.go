package proto_converter

import (
	pb "github.com/alserov/device-shop/proto/gen"
	"github.com/alserov/device-shop/user-service/internal/db"
)

func DeviceToRepoStruct(s *pb.Device) db.Device {
	return db.Device{
		UUID:         s.UUID,
		Title:        s.Title,
		Description:  s.Description,
		Price:        s.Price,
		Manufacturer: s.Manufacturer,
		Amount:       s.Amount,
	}
}
