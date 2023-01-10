package client

import (
	"fmt"
	"log"
	"net/url"
)

// PipelineService used to validate CRON expression
type CronService struct {
	*Client
}

// ValidateCronExpression response for validation quartz CRON expression
type ValidateCronExpression struct {
	Valid   bool   `json:"valid"`
	Message string `json:"message"`
}

// Validate an CRON expression against a Quartz Scheduler
func (service *CronService) Validate(expression string) (*ValidateCronExpression, error) {
	log.Printf("[INFO] Validate CRON expression \"%s\"\n", expression)

	path := fmt.Sprintf("/cron/validate?expression=%s", url.QueryEscape(expression))
	req, err := service.NewRequest("GET", path)
	if err != nil {
		return nil, err
	}

	var validateResponse ValidateCronExpression
	if _, err := service.DoWithResponse(req, &validateResponse); err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] Response from validate: %v\n", validateResponse)

	return &validateResponse, nil
}
