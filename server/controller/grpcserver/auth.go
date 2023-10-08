package grpcserver

import (
	"context"
	"database/sql"
	"strings"

	"github.com/gogapopp/gophkeeper/models"
	pb "github.com/gogapopp/gophkeeper/proto"
	"github.com/gogapopp/gophkeeper/server/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (gs *grpcServer) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var response pb.RegisterResponse
	user := models.User{
		Login:      in.Login,
		Password:   in.Password,
		UserPhrase: in.UserPhrase,
	}
	err := gs.auth.Register(ctx, user)
	if err != nil {
		gs.log.Error(err)
		if strings.Contains(err.Error(), repository.ErrUserAlreadyExists.Error()) {
			return nil, status.Error(codes.AlreadyExists, "user already exists")
		}
		return nil, status.Error(codes.InvalidArgument, "bad request")
	}
	response.Result = "OK"
	return &response, nil
}

func (gs *grpcServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	var response pb.LoginResponse
	user := models.User{
		Login:      in.Login,
		Password:   in.Password,
		UserPhrase: in.UserPhrase,
	}
	token, err := gs.auth.Login(ctx, user)
	if err != nil {
		gs.log.Error(err)
		if strings.Contains(err.Error(), sql.ErrNoRows.Error()) {
			return nil, status.Error(codes.Unauthenticated, "incorrect password or login")
		}
		return nil, status.Error(codes.InvalidArgument, "bad request")
	}
	response.Result = "OK"
	response.Jwt = &token
	return &response, nil
}
