package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/alserov/device-shop/user-service/internal/db"
	"github.com/alserov/device-shop/user-service/internal/db/models"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func NewRepo(db *sql.DB, log *slog.Logger) db.UserRepo {
	return &repo{
		db:  db,
		log: log,
	}
}

type repo struct {
	log *slog.Logger
	db  *sql.DB
}

const (
	notFound      = "user not found"
	internalError = "internal error"
	balanceError  = "insufficient funds"

	balanceConstraintErrorCode = "23514"
)

func (r *repo) GetInfo(_ context.Context, userUUID string) (models.GetUserInfoRes, error) {
	op := "repo.GetInfo"

	query := `SELECT username,email,uuid,cash FROM users WHERE uuid = $1`

	var info models.GetUserInfoRes
	err := r.db.QueryRow(query, userUUID).Scan(&info.Username, &info.Email, &info.UUID, &info.Cash)
	if errors.Is(sql.ErrNoRows, err) {
		return models.GetUserInfoRes{}, status.Error(codes.NotFound, notFound)
	}
	if err != nil {
		r.log.Error("failed to query row", slog.String("error", err.Error()), slog.String("op", op))
		return models.GetUserInfoRes{}, status.Error(codes.Internal, internalError)
	}

	return info, nil
}

func (r *repo) TopUpBalance(_ context.Context, req models.BalanceReq) (float32, error) {
	op := "repo.TopUpBalance"

	query := `UPDATE users SET cash = cash + $1 WHERE uuid = $2 RETURNING cash`

	var cash float32
	err := r.db.QueryRow(query, req.Cash, req.UserUUID).Scan(&cash)
	if errors.Is(sql.ErrNoRows, err) {
		return 0, status.Error(codes.NotFound, notFound)
	}
	if err != nil {
		r.log.Error("failed to query row", slog.String("error", err.Error()), slog.String("op", op))
		return 0, status.Error(codes.Internal, internalError)
	}

	return cash, nil
}

func (r *repo) DebitBalance(_ context.Context, req models.BalanceReq) (float32, error) {
	op := "repo.DebitBalance"

	query := `UPDATE users SET cash = cash - $1 WHERE uuid = $2 RETURNING cash`

	var cash float32
	err := r.db.QueryRow(query, req.Cash, req.UserUUID).Scan(&cash)
	if errors.Is(sql.ErrNoRows, err) {
		return 0, status.Error(codes.NotFound, notFound)
	}
	if err != nil {
		r.log.Error("failed to query row", slog.String("error", err.Error()), slog.String("op", op))
		return 0, status.Error(codes.Internal, internalError)
	}

	return cash, nil
}

func (r *repo) DebitBalanceTx(_ context.Context, req models.BalanceReq) (*sql.Tx, error) {
	op := "repo.DebitBalanceTx"
	query := `UPDATE users SET cash = cash - $1 WHERE uuid = $2`

	tx, err := r.db.Begin()
	if err != nil {
		r.log.Error("failed to begin tx", slog.String("error", err.Error()), slog.String("op", op))
		tx.Rollback()
		return tx, status.Error(codes.Internal, internalError)
	}

	_, err = tx.Exec(query, req.Cash, req.UserUUID)
	if errors.Is(sql.ErrNoRows, err) {
		return tx, status.Error(codes.NotFound, notFound)
	}
	if err, ok := err.(*pq.Error); ok {
		switch err.Code {
		case balanceConstraintErrorCode:
			return tx, status.Error(codes.Canceled, balanceError)
		default:
			r.log.Error("failed to execute tx", slog.String("error", err.Error()), slog.String("op", op))
			return tx, status.Error(codes.Internal, internalError)
		}
	}

	return tx, nil
}

func (r *repo) Signup(_ context.Context, req models.SignupReq, info models.SignupInfo) error {
	op := "repo.Signup"

	query := `INSERT INTO users (uuid,username,password,email,cash,refresh_token,role,created_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`

	_, err := r.db.Exec(query, info.UUID, req.Username, req.Password, req.Email, info.Cash, info.RefreshToken, info.Role, info.CreatedAt)
	if errors.Is(sql.ErrNoRows, err) {
		return status.Error(codes.NotFound, notFound)
	}
	if err != nil {
		r.log.Error("failed to exec query", slog.String("error", err.Error()), slog.String("op", op))
		return status.Error(codes.Internal, internalError)
	}

	return nil
}

func (r *repo) Login(_ context.Context, req models.LoginReq, rToken string) (string, error) {
	op := "repo.Login"

	query := `UPDATE users SET refresh_token = $1 WHERE username = $2 RETURNING uuid`

	var uuid string
	err := r.db.QueryRow(query, rToken, req.Username).Scan(&uuid)
	if errors.Is(sql.ErrNoRows, err) {
		return "", status.Error(codes.NotFound, notFound)
	}
	if err != nil {
		r.log.Error("failed to query row", slog.String("error", err.Error()), slog.String("op", op))
		return "", status.Error(codes.Internal, notFound)
	}

	return uuid, nil
}

func (r *repo) GetPasswordAndRoleByUsername(_ context.Context, uname string) (string, string, error) {
	op := "repo.GetPasswordAndRoleByUsername"

	query := `SELECT password, role FROM users WHERE username = $1 LIMIT 1`

	var (
		password string
		role     string
	)

	err := r.db.QueryRow(query, uname).Scan(&password, &role)
	if errors.Is(err, sql.ErrNoRows) {
		return "", "", status.Error(codes.NotFound, notFound)
	}
	if err != nil {
		r.log.Error("failed to query row", slog.String("error", err.Error()), slog.String("op", op))
		return "", "", status.Error(codes.Internal, internalError)
	}

	return password, role, nil
}

func (r *repo) GetUserInfo(_ context.Context, uuid string) (models.GetUserInfoRes, error) {
	op := "repo.GetUserInfo"

	query := `SELECT username,email,uuid,cash FROM users WHERE uuid=$1`

	var info models.GetUserInfoRes
	err := r.db.QueryRow(query, uuid).Scan(&info.Username, &info.Email, &info.UUID, &info.Cash)
	if errors.Is(sql.ErrNoRows, err) {
		return models.GetUserInfoRes{}, status.Error(codes.NotFound, notFound)
	}
	if err != nil {
		r.log.Error("failed to query row", slog.String("error", err.Error()), slog.String("op", op))
		return models.GetUserInfoRes{}, status.Error(codes.Internal, internalError)
	}

	return info, nil
}

func (r *repo) CheckIfAdmin(_ context.Context, uuid string) (bool, error) {
	op := "repo.CheckIfAdmin"

	query := `SELECT EXISTS(SELECT 1 FROM users WHERE uuid=$1 AND role=admin)`

	var isAdmin bool
	err := r.db.QueryRow(query, uuid).Scan(&isAdmin)
	if errors.Is(sql.ErrNoRows, err) {
		return false, status.Error(codes.NotFound, notFound)
	}
	if err != nil {
		r.log.Error("failed to query row", slog.String("error", err.Error()), slog.String("op", op))
		return false, status.Error(codes.Internal, internalError)
	}

	return isAdmin, nil
}
