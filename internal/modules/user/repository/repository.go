package repository

import (
	"context"
	"go-training/internal/modules/user/entity"
)

type UserRepository interface {
	FindUser(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*entity.User, error)
	CreateUser(ctx context.Context, data *entity.UserCreate) error
}
