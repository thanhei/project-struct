package business

import (
	"context"
	"go-training/internal/common"
	"go-training/internal/modules/user/entity"
)

func (biz *business) Register(ctx context.Context, data *entity.UserCreate) error {
	user, err := biz.userRepo.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		if err != common.RecordNotFound {
			return common.ErrDB(err)
		}
	}

	if user.Id != 0 {
		return entity.ErrEmailExisted
	}

	salt := common.GenSalt(50)

	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user" // hard code
	data.Status = 1

	if err := biz.userRepo.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(entity.EntityName, err)
	}

	return nil
}
