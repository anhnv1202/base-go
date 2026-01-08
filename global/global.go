package global

import (
	"github.com/anhnv1202/base-go/pkg/logger"
	"github.com/anhnv1202/base-go/pkg/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config *setting.Config
	Logger *logger.Logger
	DB     *gorm.DB
	Redis  redis.UniversalClient
)
