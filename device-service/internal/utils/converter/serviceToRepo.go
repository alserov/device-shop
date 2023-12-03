package converter

import (
	repo "github.com/alserov/device-shop/device-service/internal/db/models"
	"github.com/alserov/device-shop/device-service/internal/service/models"
)

func GetDevicesByPriceToRepo(req models.GetByPrice) repo.GetByPrice {
	return repo.GetByPrice{
		Min: req.Min,
		Max: req.Max,
	}
}
