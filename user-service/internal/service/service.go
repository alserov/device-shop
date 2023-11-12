package service

import (
	"context"
	"database/sql"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/gateway/pkg/models"
	pb "github.com/alserov/device-shop/proto/gen"
	"github.com/alserov/device-shop/user-service/internal/db/mongo"
	"github.com/alserov/device-shop/user-service/internal/db/postgres"
	"github.com/alserov/device-shop/user-service/internal/entity"
	"github.com/alserov/device-shop/user-service/pkg/utils"
	"github.com/google/uuid"
	mg "go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net/http"
	"os"
	"time"
)

type service struct {
	postgres *sql.DB
	mongo    *mg.Client
}

func New(pg *sql.DB, mg *mg.Client) pb.UsersServer {
	return &service{
		postgres: pg,
		mongo:    mg,
	}
}

var DEVICES_ADDR = os.Getenv("DEVICE_ADDR")

func (s *service) Signup(ctx context.Context, req *pb.SignupReq) (*pb.SignupRes, error) {
	exists, err := postgres.NewRepo(s.postgres).CheckIfExistsByUsername(ctx, req.Username)
	if err != nil {
		return &pb.SignupRes{}, err
	}
	if exists {
		return &pb.SignupRes{}, status.Error(http.StatusBadRequest, "already exists")
	}

	r := &entity.User{
		Username:  req.Username,
		Email:     req.Email,
		UUID:      uuid.New().String(),
		Role:      "user",
		Cash:      0,
		CreatedAt: time.Now().UTC(),
		Password:  req.Password,
	}

	r.Token, r.RefreshToken, err = utils.GenerateTokens("user")
	if err != nil {
		return &pb.SignupRes{}, err
	}

	if err = r.HashPassword(); err != nil {
		return &pb.SignupRes{}, err
	}

	if err = postgres.NewRepo(s.postgres).Signup(ctx, r); err != nil {
		return &pb.SignupRes{}, err
	}

	go func() {
		if err = utils.SendEmail(r.Email); err != nil {
			log.Println("FAILED TO SEND EMAIL: ", err.Error())
		}
	}()

	return &pb.SignupRes{
		Username:     r.Username,
		Email:        r.Email,
		UUID:         r.UUID,
		Cash:         int32(r.Cash),
		RefreshToken: r.RefreshToken,
		Token:        r.Token,
	}, nil
}

func (s *service) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {
	user, err := postgres.NewRepo(s.postgres).FindByUsername(ctx, req.Username)
	if err != nil {
		return &pb.LoginRes{}, err
	}

	if err = user.CheckPassword(req.Password); err != nil {
		return nil, status.Error(http.StatusBadRequest, "invalid password")
	}

	token, rToken, err := utils.GenerateTokens(user.Role)
	if err != nil {
		return &pb.LoginRes{}, err
	}

	r := &entity.RepoLoginReq{
		RefreshToken: rToken,
		Username:     user.Username,
	}
	if err = postgres.NewRepo(s.postgres).Login(ctx, r); err != nil {
		return &pb.LoginRes{}, err
	}

	return &pb.LoginRes{
		RefreshToken: rToken,
		Token:        token,
	}, nil
}

func (s *service) AddToFavourite(ctx context.Context, req *pb.AddReq) (*emptypb.Empty, error) {
	cl, cc, err := client.DialDevices(DEVICES_ADDR)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	defer cc.Close()

	getDeviceReq := &pb.UUIDReq{
		UUID: req.DeviceUUID,
	}

	device, err := cl.GetDeviceByUUID(ctx, getDeviceReq)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	r := &entity.AddReq{
		UserUUID: req.UserUUID,
		Device: &models.Device{
			UUID:         device.UUID,
			Title:        device.Title,
			Manufacturer: device.Manufacturer,
			Price:        device.Price,
			Description:  device.Description,
		},
	}

	if err = mongo.NewRepo(s.mongo).AddToFavourite(ctx, r); err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (s *service) RemoveFromFavourite(ctx context.Context, req *pb.RemoveReq) (*emptypb.Empty, error) {
	err := mongo.NewRepo(s.mongo).RemoveFromFavourite(ctx, &entity.RemoveReq{
		UserUUID:   req.UserUUID,
		DeviceUUID: req.DeviceUUID,
	})
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (s *service) GetFavourite(ctx context.Context, req *pb.GetReq) (*pb.GetRes, error) {
	coll, err := mongo.NewRepo(s.mongo).GetFavourite(ctx, req.UserUUID)
	if err != nil {
		return &pb.GetRes{}, err
	}

	var devices []*pb.Device

	for _, v := range coll {
		device := &pb.Device{
			UUID:         v.UUID,
			Title:        v.Title,
			Description:  v.Description,
			Price:        v.Price,
			Manufacturer: v.Manufacturer,
		}
		devices = append(devices, device)
	}

	return &pb.GetRes{
		Devices: devices,
	}, nil
}

func (s *service) AddToCart(ctx context.Context, req *pb.AddReq) (*emptypb.Empty, error) {
	cl, cc, err := client.DialDevices(DEVICES_ADDR)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	defer cc.Close()

	getDeviceReq := &pb.UUIDReq{
		UUID: req.DeviceUUID,
	}

	device, err := cl.GetDeviceByUUID(ctx, getDeviceReq)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	r := &entity.AddReq{
		UserUUID: req.UserUUID,
		Device: &models.Device{
			UUID:         device.UUID,
			Title:        device.Title,
			Manufacturer: device.Manufacturer,
			Price:        device.Price,
			Description:  device.Description,
		},
	}

	if err = mongo.NewRepo(s.mongo).AddToCart(ctx, r); err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (s *service) RemoveFromCart(ctx context.Context, req *pb.RemoveReq) (*emptypb.Empty, error) {
	err := mongo.NewRepo(s.mongo).RemoveFromCart(ctx, &entity.RemoveReq{
		UserUUID:   req.UserUUID,
		DeviceUUID: req.GetDeviceUUID(),
	})
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (s *service) GetCart(ctx context.Context, req *pb.GetReq) (*pb.GetRes, error) {
	coll, err := mongo.NewRepo(s.mongo).GetCart(ctx, req.UserUUID)
	if err != nil {
		return &pb.GetRes{}, err
	}

	var devices []*pb.Device

	for _, v := range coll {
		device := &pb.Device{
			UUID:         v.UUID,
			Title:        v.Title,
			Description:  v.Description,
			Price:        v.Price,
			Manufacturer: v.Manufacturer,
		}
		devices = append(devices, device)
	}

	return &pb.GetRes{
		Devices: devices,
	}, nil
}
