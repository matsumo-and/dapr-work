package user

import (
	"context"
	"time"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	userv1 "github.com/matsumo_and/dapr-work/app/proto/user/v1"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) GetUser(
	ctx context.Context,
	req *connect.Request[userv1.GetUserRequest],
) (*connect.Response[userv1.GetUserResponse], error) {
	// 簡易実装: 実際にはデータベースから取得
	user := &userv1.User{
		Id:        req.Msg.Id,
		Username:  "john_doe",
		Email:     "john@example.com",
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	res := connect.NewResponse(&userv1.GetUserResponse{
		User: user,
	})

	return res, nil
}

func (h *Handler) CreateUser(
	ctx context.Context,
	req *connect.Request[userv1.CreateUserRequest],
) (*connect.Response[userv1.CreateUserResponse], error) {
	// 簡易実装: 実際にはデータベースに保存
	user := &userv1.User{
		Id:        uuid.New().String(),
		Username:  req.Msg.Username,
		Email:     req.Msg.Email,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	res := connect.NewResponse(&userv1.CreateUserResponse{
		User: user,
	})

	return res, nil
}

func (h *Handler) UpdateUser(
	ctx context.Context,
	req *connect.Request[userv1.UpdateUserRequest],
) (*connect.Response[userv1.UpdateUserResponse], error) {
	// 簡易実装: 実際にはデータベースを更新
	user := &userv1.User{
		Id:        req.Msg.Id,
		Username:  req.Msg.Username,
		Email:     req.Msg.Email,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	res := connect.NewResponse(&userv1.UpdateUserResponse{
		User: user,
	})

	return res, nil
}

func (h *Handler) DeleteUser(
	ctx context.Context,
	req *connect.Request[userv1.DeleteUserRequest],
) (*connect.Response[userv1.DeleteUserResponse], error) {
	// 簡易実装: 実際にはデータベースから削除
	res := connect.NewResponse(&userv1.DeleteUserResponse{
		Success: true,
	})

	return res, nil
}
