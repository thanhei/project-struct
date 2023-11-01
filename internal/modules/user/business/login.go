package business

import (
	"context"
	"go-training/internal/common"
	"go-training/internal/component/tokenprovider"
	"go-training/internal/modules/user/entity"
)

func (biz *business) Login(ctx context.Context, data *entity.UserLogin) (*entity.Account, error) {
	user, err := biz.userRepo.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, entity.ErUsernameOrPasswordInvalid
	}
	passHashed := biz.hasher.Hash(data.Password + user.Salt)
	if user.Password != passHashed {
		return nil, entity.ErUsernameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}
	accessToken, err := biz.tokenProvider.Generate(payload, biz.config.System.TokenExpire)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	refreshToken, err := biz.tokenProvider.Generate(payload, biz.config.System.TokenExpire)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	account := entity.NewAccount(accessToken, refreshToken)

	return account, nil
}
