package usecase

import (
	"gorm.io/gorm"
	"testing"
	"users-segments-service/config"
	"users-segments-service/internal/usecase/segment_repo"
	"users-segments-service/internal/usecase/user_repo"
	"users-segments-service/pkg/database"
	"users-segments-service/pkg/validations"
)

type TestData struct {
	db             *gorm.DB
	userUseCase    *UserUseCase
	segmentUseCase *SegmentUseCase
}

var TestsContext *TestData

func Prepare(t *testing.T) {
	t.Helper()
	cfg, err := config.NewConfig()
	cfg.DB.Name = "users-segments-test"
	if err != nil {
		t.Fatal(err)
	}

	db, err := database.NewGorm(cfg.DB)
	if err != nil {
		t.Fatal(err)
	}
	if err := validations.InitValidations(); err != nil {
		t.Fatal(err)
	}
	truncateTables(db)
	userUseCase := NewUserUseCase(user_repo.New(db))
	segmentUseCase := NewSegmentUseCase(segment_repo.New(db))

	testData := &TestData{
		db:             db,
		userUseCase:    userUseCase,
		segmentUseCase: segmentUseCase,
	}

	TestsContext = testData
}

func truncateTables(db *gorm.DB) {
	truncateUsers(db)
	truncateSegments(db)
}

func truncateUsers(db *gorm.DB) {
	db.Exec("TRUNCATE TABLE users CASCADE")
}

func truncateSegments(db *gorm.DB) {
	db.Exec("TRUNCATE TABLE segments CASCADE")
}
