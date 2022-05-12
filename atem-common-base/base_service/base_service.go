package base_service

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

type BaseService struct {
	Orm                 *gorm.DB
	Redis               *redis.Client
	Logger              *zap.Logger
	Concurrency_Control *singleflight.Group
}
