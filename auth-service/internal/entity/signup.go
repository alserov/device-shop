package entity

import "time"

type SignupAdditional struct {
	UUID         string
	Cash         float32
	RefreshToken string
	Role         string
	CreatedAt    *time.Time
}
