package proto_converter

import (
	"github.com/alserov/device-shop/auth-service/internal/db"
	pb "github.com/alserov/device-shop/proto/gen"
)

func SignupReqToRepoStruct(s *pb.SignupReq) db.SignupReq {
	return db.SignupReq{
		Username: s.Username,
		Email:    s.Email,
		Password: s.Password,
	}
}

func LoginReqToRepoStruct(s *pb.LoginReq) db.LoginReq {
	return db.LoginReq{
		Username: s.Username,
		Password: s.Password,
	}
}
