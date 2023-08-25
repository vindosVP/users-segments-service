package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"users-segments-service/config"
	"users-segments-service/internal/entity"
)

func NewGorm(cfg config.DB) (*gorm.DB, error) {
	dns := generateGormDNS(cfg)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("database - NewGorm - gorm.Open: %w", err)
	}

	err = autoMigrate(db)
	if err != nil {
		return nil, err
	}

	return db, err
}

func autoMigrate(db *gorm.DB) error {

	if err := db.AutoMigrate(&entity.User{}); err != nil {
		return fmt.Errorf("database - autoMigrate - db.AutoMigrate: %w", err)
	}
	if err := db.AutoMigrate(&entity.Segment{}); err != nil {
		return fmt.Errorf("database - autoMigrate - db.AutoMigrate: %w", err)
	}
	if err := db.AutoMigrate(&entity.SegmentOperation{}); err != nil {
		return fmt.Errorf("database - autoMigrate - db.AutoMigrate: %w", err)
	}
	if err := db.AutoMigrate(&entity.SegmentUser{}); err != nil {
		return fmt.Errorf("database - autoMigrate - db.AutoMigrate: %w", err)
	}

	return nil
}

func generateGormDNS(cfg config.DB) string {
	if cfg.DNS != "" {
		return cfg.DNS
	}

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Pwd, cfg.Name, cfg.SSLMode)
}
