package migration

import (
	"context"
	"fmt"
	"github.com/vertinofff/blog-api/data/models"
	"gorm.io/gorm"
)

// Up is deliberately explicit and returns failures so the process cannot start
// on a partially migrated schema. In multi-replica production it should be run
// by a single migration job, not every application replica.
func Up(ctx context.Context, database *gorm.DB) error {
	return database.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if e := tx.AutoMigrate(&models.User{}, &models.Post{}); e != nil {
			return fmt.Errorf("migrate schema: %w", e)
		}
		return nil
	})
}
