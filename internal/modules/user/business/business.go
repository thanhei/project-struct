package business

import (
	"context"
	"go-training/internal/common"
	hasher "go-training/internal/component/hasher"
	"go-training/internal/component/tokenprovider"
	"go-training/internal/modules/user/entity"
	userrepo "go-training/internal/modules/user/repository"
)

type UserBusiness interface {
	Login(ctx context.Context, data *entity.UserLogin) (*entity.Account, error)
	Register(ctx context.Context, data *entity.UserCreate) error
}

type business struct {
	config        *common.Config
	userRepo      userrepo.UserRepository
	tokenProvider tokenprovider.Provider
	hasher        hasher.Hasher
}

var _ UserBusiness = (*business)(nil)

func NewBusiness(config *common.Config, userRepo userrepo.UserRepository, tokenProvider tokenprovider.Provider, hasher hasher.Hasher) *business {
	return &business{
		config:        config,
		userRepo:      userRepo,
		tokenProvider: tokenProvider,
		hasher:        hasher,
	}
}
