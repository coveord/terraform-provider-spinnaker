package provider

import (
	"errors"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

// Jenkins trigger for Pipeline
type CRONTrigger struct {
	baseTrigger `mapstructure:",squash"`

	CronExpression string `mapstructure:"cron_expression"`
}

func newCRONTrigger() *CRONTrigger {
	return &CRONTrigger{}
}

func (t *CRONTrigger) toClientTrigger(id string) (client.Trigger, error) {
	clientTrigger := client.NewCRONTrigger()
	clientTrigger.ID = id
	clientTrigger.Enabled = t.Enabled

	clientTrigger.CronExpression = t.CronExpression
	return clientTrigger, nil
}

func (*CRONTrigger) fromClientTrigger(clientTriggerInterface client.Trigger) (trigger, error) {
	clientTrigger, ok := clientTriggerInterface.(*client.CRONTrigger)
	if !ok {
		return nil, errors.New("expected CRON trigger")
	}
	t := newCRONTrigger()
	t.Enabled = clientTrigger.Enabled
	t.CronExpression = clientTrigger.CronExpression
	return t, nil
}

func (t *CRONTrigger) setResourceData(d *schema.ResourceData) error {
	var err error
	err = d.Set("enabled", t.Enabled)
	if err != nil {
		return err
	}
	err = d.Set("cron_expression", t.CronExpression)
	if err != nil {
		return err
	}
	return nil
}
