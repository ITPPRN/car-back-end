package databases

import (
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"testBackend/configs"
	"testBackend/logs"
	"testBackend/pkg/utils"
)

func NewPostgresConnection(cfg *configs.Config) (*gorm.DB, error) {

	dsn, err := utils.UrlBuilder("postgres", cfg)
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(postgres.New(postgres.Config{
		// DriverName:           cfg.Postgres.DriverName,
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		logs.Error("Failed to connect to database: ", zap.Error(err))
	}
	logs.Info("postgreSQL database has been connected 🐘")
	return db, nil
}
