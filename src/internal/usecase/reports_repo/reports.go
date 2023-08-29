package reports_repo

import (
	"encoding/csv"
	"fmt"
	"github.com/google/uuid"
	"os"
)

type ReportsRepository struct {
	rd string
}

func New(reportsDirectory string) *ReportsRepository {
	return &ReportsRepository{
		rd: reportsDirectory,
	}
}

func (rr *ReportsRepository) CreateCSVReport(operations [][]string) (string, error) {

	fileID := uuid.New().String()
	f, err := os.Create(fmt.Sprintf("%s\\%s.csv", rr.rd, fileID))
	defer f.Close()
	if err != nil {
		return "", err
	}
	w := csv.NewWriter(f)
	err = w.WriteAll(operations)
	if err != nil {
		return "", err
	}

	return fileID, err
}
