package userreposql

import (
	"context"
	"go-training/internal/common"
	"go-training/internal/modules/user/entity"
)

func (s *sqlRepo) CreateUser(ctx context.Context, data *entity.UserCreate) error {
	db := s.db.Begin()

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil

}
