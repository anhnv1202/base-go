package global

import (
	"github.com/anhnv1202/base-go/pkg/logger"
	"github.com/anhnv1202/base-go/pkg/setting"
	"gorm.io/gorm"
)

var (
    Config *setting.Config
    Logger *logger.Logger
    DB *gorm.DB
)
