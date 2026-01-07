package initialize

import (
	"context"
	"time"

	"github.com/anhnv1202/base-go/global"
	"github.com/anhnv1202/base-go/pkg/setting"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
        global.Logger.Err(errString, err)
		panic(errString)
	}
}

func InitDatabase() {
	m := global.Config.Database
    gormLogger := logger.Default.LogMode(logger.Silent)
	if m.Host == "localhost" {
		gormLogger = logger.Default.LogMode(logger.Info)
	}
	db, err := gorm.Open(postgres.Open(m.DSN()), &gorm.Config{
		Logger: gormLogger,
        SkipDefaultTransaction: true,
        PrepareStmt: true,
	})
	checkErrorPanic(err, "failed to connect database")

	SetPool(m)

	global.DB = db
    global.Logger.With(
        zap.String("host", m.Host),
        zap.Int("port", m.Port),
        zap.String("user", m.User),
        zap.String("dbname", m.Name),
    ).Info("Database connected")
}

func SetPool(m setting.DatabaseConfig) {
	sqlDB, err := global.DB.DB()
	checkErrorPanic(err, "failed to get underlying sql.DB")
	// Configure connection pool
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(m.ConnMaxLifetime)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	checkErrorPanic(sqlDB.PingContext(ctx), "failed to ping database")
}

func migrateTables() {
    
}
