package userreposql

import (
	"context"
	"go-training/internal/common"
	"go-training/internal/modules/user/entity"

	"gorm.io/gorm"
)

func (s *sqlRepo) FindUser(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*entity.User, error) {
	db := s.db.Table(entity.User{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user entity.User

	err := db.Where(condition).Find(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &user, nil
}
