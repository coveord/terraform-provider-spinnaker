package provider

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

const cron_description = `Quartz CRON expression describing when to trigger this pipeline.
See https://www.quartz-scheduler.org/documentation/2.3.1-SNAPSHOT/tutorials/tutorial-lesson-06.html
and https://www.freeformatter.com/cron-expression-generator-quartz.html for more information.`

const expression_field_name = "cron_expression"

func pipelineCRONTriggerResource() *schema.Resource {
	triggerInterface := func() trigger {
		return newCRONTrigger()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			// Add custom validation on field cron_expression
			if _, err := validateExpression(d.Get(expression_field_name), m); err != nil {
				return err
			}
			return resourcePipelineTriggerCreate(d, m, triggerInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineTriggerRead(d, m, triggerInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			// Add custom validation on field cron_expression
			if _, err := validateExpression(d.Get(expression_field_name), m); err != nil {
				return err
			}
			return resourcePipelineTriggerUpdate(d, m, triggerInterface)
		},
		Delete: resourcePipelineTriggerDelete,
		Importer: &schema.ResourceImporter{
			State: resourceTriggerImporter,
		},

		Schema: triggerResource(map[string]*schema.Schema{
			expression_field_name: {
				Type:        schema.TypeString,
				Description: cron_description,
				Required:    true,
			},
		}),
	}
}

func validateExpression(expr interface{}, m interface{}) (bool, error) {
	cronExpr := expr.(string)
	cronService := m.(*Services).CronService

	validation, err := cronService.Validate(cronExpr)
	if err != nil {
		return false, err
	}

	if !validation.Valid {
		return validation.Valid, fmt.Errorf("while validating expression: %s", validation.Message)
	}

	return true, nil
}
