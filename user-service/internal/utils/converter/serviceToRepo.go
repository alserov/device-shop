package converter

func ServiceSignupReqToRepo(req *models.SignupReq) *models.SignupReq {
	return &models.SignupReq{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
}

func ServiceLoginReqToRepo(req *models.LoginReq) *models.LoginReq {
	return &models.LoginReq{
		Username: req.Username,
		Password: req.Password,
	}
}

func RepoUserInfoResToService(res *models.GetUserInfoRes) *models.GetUserInfoRes {
	return &models.GetUserInfoRes{
		UUID:     res.UUID,
		Username: res.Username,
		Email:    res.Email,
		Cash:     res.Cash,
	}
}
