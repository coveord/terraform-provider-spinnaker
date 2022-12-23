package client

import "github.com/mitchellh/mapstructure"

// CRONTriggerType cron trigger
var CRONTriggerType TriggerType = "cron"

func init() {
	triggerFactories[CRONTriggerType] = parseCRONTrigger
}

// CRONTrigger for Pipeline
type CRONTrigger struct {
	baseTrigger `mapstructure:",squash"`

	CronExpression string `json:"cronExpression"`
}

// NewCRONTrigger new trigger
func NewCRONTrigger() *CRONTrigger {
	return &CRONTrigger{
		baseTrigger: *newBaseTrigger(CRONTriggerType),
	}
}

func parseCRONTrigger(triggerMap map[string]interface{}) (Trigger, error) {
	trigger := NewCRONTrigger()
	if err := mapstructure.Decode(triggerMap, trigger); err != nil {
		return nil, err
	}
	return trigger, nil
}
