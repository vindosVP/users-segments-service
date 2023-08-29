package v1

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"users-segments-service/pkg/logger"
)

type ReportRoutes struct {
	rd string
	l  logger.Interface
}

func SetReportsRoutes(handler fiber.Router, reportsDirectory string, l logger.Interface) {
	r := &ReportRoutes{
		rd: reportsDirectory,
		l:  l,
	}
	h := handler.Group("/reports")
	h.Get("/:reportID", r.downloadReport)
}

// @Summary     Download
// @Description Download a report
// @ID          download
// @Tags  	    report
// @Accept      json
// @Produce     application/pdf
// @Param 		reportID path string true "user ID" example(80ef1ba7-1045-41aa-a8a2-4c0aba407baf)
// @Success     200 {file} file
// @Failure     404 {file} file
// @Router      /reports/{reportID} [get]
func (r *ReportRoutes) downloadReport(c *fiber.Ctx) error {

	id := c.Params("reportID")
	if id == "" {
		return ErrorResponse(c, fiber.StatusBadRequest, "No reportId provided", nil, nil)
	}

	filePath := fmt.Sprintf("%s\\%s.csv", r.rd, id)

	return c.Download(filePath)
}
