package sql

import (
	"context"
	"go-training/internal/common"
	"go-training/internal/modules/restaurant/entity"
)

func (s *sqlRepo) ListDataWithCondition(context context.Context, condition map[string]interface{}, filter *entity.Filter, paging *common.Paging, moreKeys ...string) ([]entity.Restaurant, error) {
	var result []entity.Restaurant

	db := s.db

	db = db.Table(entity.Restaurant{}.TableName()).Where(condition).Where("status in (1)")

	if v := filter; v != nil {
		if v.OwnerId > 0 {
			db = db.Where("owner_id = ?", v.OwnerId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if paging.FakeCursor != "" {
		if uid, err := common.FromBase58(paging.FakeCursor); err == nil {
			db = db.Where("id < ?", uid.GetLocalID())
		} else {
			db = db.Offset((paging.Page - 1) * paging.Limit)
		}
	}

	if err := db.Limit(paging.Limit).Order("id desc").Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
