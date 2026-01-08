package initialize

import (
	"context"
	"regexp"
	"time"

	"github.com/anhnv1202/base-go/global"
	"github.com/anhnv1202/base-go/internal/po"
	"github.com/anhnv1202/base-go/pkg/setting"
	"github.com/anhnv1202/base-go/pkg/utils"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var validDBName = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)

func ensureDBExists(m setting.DatabaseConfig) {
	if !validDBName.MatchString(m.Name) {
		global.Logger.Warn("Invalid database name, skipping auto-create", zap.String("dbname", m.Name))
		return
	}

	db, err := gorm.Open(postgres.Open(m.DSNWithoutDB()), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		global.Logger.Warn("Cannot connect to create database", zap.String("error", err.Error()))
		return
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var exists bool
	db.Raw("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = ?)", m.Name).Scan(&exists)
	if !exists {
		if err := db.Exec("CREATE DATABASE " + m.Name).Error; err != nil {
			global.Logger.Warn("Failed to create database (may lack CREATEDB permission)", zap.String("dbname", m.Name), zap.String("error", err.Error()))
			return
		}
		global.Logger.Info("Database created", zap.String("dbname", m.Name))
	}
}

func InitDatabase() {
	m := global.Config.Database
	ensureDBExists(m)
    gormLogger := logger.Default.LogMode(logger.Silent)
	if m.Host == "localhost" {
		gormLogger = logger.Default.LogMode(logger.Info)
	}
	db, err := gorm.Open(postgres.Open(m.DSN()), &gorm.Config{
		Logger: gormLogger,
        SkipDefaultTransaction: true,
        PrepareStmt: true,
	})
	utils.CheckErrorPanic(err, "failed to connect database")

	global.DB = db
	SetPool(m)
	migrateTables()

	OnShutdown(func() {
		if sqlDB, err := db.DB(); err == nil {
			sqlDB.Close()
			global.Logger.Info("Database connection closed")
		}
	})

	global.Logger.With(
		zap.String("host", m.Host),
		zap.Int("port", m.Port),
		zap.String("user", m.User),
		zap.String("dbname", m.Name),
	).Info("Database connected")
}

func SetPool(m setting.DatabaseConfig) {
	sqlDB, err := global.DB.DB()
	utils.CheckErrorPanic(err, "failed to get underlying sql.DB")
	// Configure connection pool
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(m.ConnMaxLifetime)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	utils.CheckErrorPanic(sqlDB.PingContext(ctx), "failed to ping database")
}

func migrateTables() {
	err := global.DB.AutoMigrate(po.All()...)
	utils.CheckErrorPanic(err, "failed to migrate tables")
}
