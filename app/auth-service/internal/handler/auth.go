package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	authv1 "github.com/matsumo_and/dapr-work/app/auth-service/proto/authv1"
	"github.com/matsumo_and/dapr-work/app/auth-service/proto/authv1/authv1connect"
)

type AuthHandler struct {
	authv1connect.UnimplementedAuthServiceHandler
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Login(
	ctx context.Context,
	req *connect.Request[authv1.LoginRequest],
) (*connect.Response[authv1.LoginResponse], error) {
	// 簡易実装: 実際にはデータベースでユーザー認証を行う
	userID := uuid.New().String()
	token := uuid.New().String()

	res := connect.NewResponse(&authv1.LoginResponse{
		Token:  token,
		UserId: userID,
	})

	return res, nil
}

func (h *AuthHandler) ValidateToken(
	ctx context.Context,
	req *connect.Request[authv1.ValidateTokenRequest],
) (*connect.Response[authv1.ValidateTokenResponse], error) {
	// 簡易実装: 実際にはトークンをRedisなどで検証
	res := connect.NewResponse(&authv1.ValidateTokenResponse{
		Valid:  true,
		UserId: "dummy-user-id",
	})

	return res, nil
}

func (h *AuthHandler) Logout(
	ctx context.Context,
	req *connect.Request[authv1.LogoutRequest],
) (*connect.Response[authv1.LogoutResponse], error) {
	// 簡易実装: 実際にはトークンを無効化
	res := connect.NewResponse(&authv1.LogoutResponse{
		Success: true,
	})

	return res, nil
}
