package memcache

import (
	"context"
	"fmt"
	"go-training/internal/modules/user/entity"
)

type RealStore interface {
	FindUser(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*entity.User, error)
}

type userCaching struct {
	store     Caching
	realStore RealStore
}

func NewUserCaching(store Caching, realStore RealStore) *userCaching {
	return &userCaching{
		store:     store,
		realStore: realStore,
	}
}

func (uc *userCaching) FindUser(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*entity.User, error) {
	userId := condition["id"].(int)
	key := fmt.Sprintf("user-%d", userId)
	userInCache := uc.store.Read(key)
	if userInCache != nil {
		return userInCache.(*entity.User), nil
	}

	user, err := uc.realStore.FindUser(ctx, condition, moreInfo...)
	if err != nil {
		return nil, err
	}

	uc.store.Write(key, user)
	return user, nil
}
