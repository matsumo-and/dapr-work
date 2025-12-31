package auth

import (
	"context"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	authv1 "github.com/matsumo_and/dapr-work/app/proto/auth/v1"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Login(
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

func (h *Handler) ValidateToken(
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

func (h *Handler) Logout(
	ctx context.Context,
	req *connect.Request[authv1.LogoutRequest],
) (*connect.Response[authv1.LogoutResponse], error) {
	// 簡易実装: 実際にはトークンを無効化
	res := connect.NewResponse(&authv1.LogoutResponse{
		Success: true,
	})

	return res, nil
}
