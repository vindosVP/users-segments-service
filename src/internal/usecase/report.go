package usecase

import (
	"strconv"
	"users-segments-service/internal/entity"
)

type ReportUseCase struct {
	reportsRepo ReportsRepo
}

func NewReportsUseCase(rr ReportsRepo) *ReportUseCase {
	return &ReportUseCase{
		reportsRepo: rr,
	}
}

func (r ReportUseCase) SaveReport(operations []entity.UsersSegmentOperation) (string, error) {

	csvData := [][]string{{"UserID", "Segment", "Operation", "Time"}}

	for _, line := range operations {
		sliceLine := []string{strconv.FormatInt(line.UserID, 10), line.Segment, line.Operation, line.Time.Format("2006-01-02 15:04:05")}
		csvData = append(csvData, sliceLine)
	}

	return r.reportsRepo.CreateCSVReport(csvData)
}
