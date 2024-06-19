package services

import (
	pb "SSO/gen/go"
	"SSO/internal/contracts/usecases"
	"SSO/pkg/jwt"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

var _ pb.AuthServiceServer = AuthServiceServer{}

// AuthServiceServer is used to implement pb.AuthServiceServer.
type AuthServiceServer struct {
	logger     *slog.Logger
	uuc        usecases.UserUseCases
	jwtManager *jwt.Manager
	pb.UnimplementedAuthServiceServer
}

func NewAuthServiceServer(logger *slog.Logger, uuc usecases.UserUseCases, jwtManager *jwt.Manager) AuthServiceServer {
	return AuthServiceServer{
		logger:     logger,
		uuc:        uuc,
		jwtManager: jwtManager,
	}
}

func (a AuthServiceServer) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := a.uuc.Login(ctx, request.Email, request.Password)
	if err != nil {
		return nil, err
	}

	tokens, err := a.jwtManager.GeneratePair(user.UUID.String())
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

func (a AuthServiceServer) RefreshToken(ctx context.Context, request *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	userIdStr, err := a.jwtManager.ValidatePair(request.AccessToken, request.RefreshToken)
	if err != nil {
		return nil, err
	}

	tokens, err := a.jwtManager.GeneratePair(userIdStr)
	if err != nil {
		return nil, err
	}

	return &pb.RefreshTokenResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

func (a AuthServiceServer) Logout(ctx context.Context, request *pb.LogoutRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}
