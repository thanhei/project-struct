package app_context

import (
	"go-training/common"
	"go-training/pubsub"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	SecretKey() string
	GetPubsub() pubsub.Pubsub
}

type appCtx struct {
	db        *gorm.DB
	secretKey string
	pb        pubsub.Pubsub
}

func NewAppContext(db *gorm.DB, config *common.Config, pb pubsub.Pubsub) *appCtx {
	return &appCtx{db, config.SecretKey, pb}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) SecretKey() string {
	return ctx.secretKey
}

func (ctx *appCtx) GetPubsub() pubsub.Pubsub {
	return ctx.pb
}
