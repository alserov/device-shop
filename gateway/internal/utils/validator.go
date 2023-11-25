package utils

import (
	"errors"
	pb "github.com/alserov/device-shop/proto/gen"
	"net/mail"
)

// ADMIN VALIDATORS
func CheckCreateDevice(r *pb.CreateDeviceReq) error {
	if len(r.Title) < 5 {
		return errors.New("title too short")
	}

	if len(r.Manufacturer) < 3 {
		return errors.New("manufacturer too short")
	}

	if len(r.Description) < 5 {
		return errors.New("description too short")
	}

	if r.Amount < 1 {
		return errors.New("amount should be more than 0")
	}

	if r.Price <= 0 {
		return errors.New("price should be more than 0")
	}

	return nil
}
func CheckUpdateDevice(r *pb.UpdateDeviceReq) error {
	if len(r.Title) < 5 {
		return errors.New("title too short")
	}
	if len(r.Description) < 5 {
		return errors.New("description too short")
	}
	if r.Price <= 0 {
		return errors.New("price should be more than 0")
	}
	if r.UUID == "" {
		return errors.New("uuid can't be empty")
	}

	return nil
}

// AUTH VALIDATORS
func CheckSignup(r *pb.SignupReq) error {
	if len(r.Password) < 5 {
		return errors.New("password is too short")
	}

	if len(r.Username) < 3 {
		return errors.New("username is too short")
	}

	if _, err := mail.ParseAddress(r.Email); err != nil {
		return errors.New("invalid email")
	}

	return nil
}
func CheckLogin(r *pb.LoginReq) error {
	if len(r.Password) < 5 {
		return errors.New("password is too short")
	}

	if len(r.Username) < 3 {
		return errors.New("username is too short")
	}

	return nil
}

// COLLECTION VALIDATOR
func CheckCollection(r *pb.ChangeCollectionReq) error {
	if r.DeviceUUID == "" {
		return errors.New("device uuid can't be empty")
	}

	if r.UserUUID == "" {
		return errors.New("user uuid can't be empty")
	}

	return nil
}

// DEVICE VALIDATORS
func CheckGetAll(r *pb.GetAllDevicesReq) error {
	if r.Index < 0 {
		return errors.New("index should be >= 0")
	}

	if r.Amount > 25 || r.Amount < 2 {
		return errors.New("amount should be more than 1 and less than 26")
	}

	return nil
}

// ORDER VALIDATORS
func CheckCreateOrder(r *pb.CreateOrderReq) error {
	if len(r.Devices) < 1 {
		return errors.New("length of cart should be more than 0 to create an order")
	}

	if r.UserUUID == "" {
		return errors.New("user uuid can't be empty")
	}

	return nil
}

func CheckUpdateOrder(r *pb.UpdateOrderReq) error {
	status := map[string]struct{}{
		"canceled":   {},
		"pending":    {},
		"delivering": {},
		"ready":      {},
		"creating":   {},
	}

	if _, ok := status[r.Status]; !ok {
		return errors.New("invalid status")
	}

	return nil
}

func CheckTopUpBalance(r *pb.BalanceReq) error {
	if r.Cash <= 0 {
		return errors.New("deposit should be more than 0")
	}

	if r.UserUUID == "" {
		return errors.New("user uuid can't be empty")
	}

	return nil
}
