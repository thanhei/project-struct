package userbiz

import (
	"context"
	"go-training/common"
	"go-training/component/app_context"
	hasher "go-training/component/hasher"
	"go-training/component/tokenprovider"
	usermodel "go-training/modules/user/model"
)

type LoginStore interface {
	FindUser(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type loginBusiness struct {
	appCtx        app_context.AppContext
	storeUser     LoginStore
	tokenProvider tokenprovider.Provider
	hasher        hasher.Hasher
	expiry        int
}

func NewLoginBusiness(storeUser LoginStore, tokenProvider tokenprovider.Provider, hasher hasher.Hasher, expiry int) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

func (biz *loginBusiness) Login(ctx context.Context, data *usermodel.UserLogin) (*usermodel.Account, error) {
	user, err := biz.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, usermodel.ErUsernameOrPasswordInvalid
	}
	passHashed := biz.hasher.Hash(data.Password + user.Salt)
	if user.Password != passHashed {
		return nil, usermodel.ErUsernameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}
	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	refreshToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	account := usermodel.NewAccount(accessToken, refreshToken)

	return account, nil
}
