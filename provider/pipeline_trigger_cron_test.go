package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

func TestAccCRONTriggerBasic(t *testing.T) {
	var pipelineRef client.Pipeline
	var triggers []client.Trigger
	pipeName := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	cronExpression := "0 0 0 1/1 * ? *"
	secondCronExpression := "0 30 0 1/1 * ? *"
	trigger1 := "spinnaker_pipeline_cron_trigger.t1"
	trigger2 := "spinnaker_pipeline_cron_trigger.t2"
	pipelineResourceName := "spinnaker_pipeline.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCRONTriggerConfigBasic(pipeName, cronExpression, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(trigger1, "cron_expression", cronExpression),
					resource.TestCheckResourceAttr(trigger2, "cron_expression", cronExpression),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineTriggers(pipelineResourceName, []string{
						trigger1,
						trigger2,
					}, &triggers),
				),
			},
			{
				ResourceName:  trigger1,
				ImportStateId: "invalid",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(`Invalid import key, must be pipelineID_triggerID`),
			},
			{
				ResourceName: trigger1,
				ImportState:  true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					if len(triggers) == 0 {
						return "", fmt.Errorf("no triggers to import")
					}
					return fmt.Sprintf("%s_%s", pipelineRef.ID, triggers[0].GetID()), nil
				},
				ImportStateVerify: true,
			},
			{
				ResourceName: trigger2,
				ImportState:  true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					if len(triggers) < 2 {
						return "", fmt.Errorf("no triggers to import")
					}
					return fmt.Sprintf("%s_%s", pipelineRef.ID, triggers[1].GetID()), nil
				},
				ImportStateVerify: true,
			},
			{
				Config: testAccCRONTriggerConfigBasic(pipeName, secondCronExpression, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(trigger1, "cron_expression", secondCronExpression),
					resource.TestCheckResourceAttr(trigger2, "cron_expression", secondCronExpression),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineTriggers(pipelineResourceName, []string{
						trigger1,
						trigger2,
					}, &triggers),
				),
			},
			{
				Config: testAccCRONTriggerConfigBasic(pipeName, cronExpression, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(trigger1, "cron_expression", cronExpression),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineTriggers(pipelineResourceName, []string{
						trigger1,
					}, &triggers),
				),
			},
			{
				Config: testAccCRONTriggerConfigBasic(pipeName, cronExpression, 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineTriggers(pipelineResourceName, []string{}, &triggers),
				),
			},
		},
	})
}

func TestAccCRONTriggerBadExpression(t *testing.T) {
	pipeName := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	cronExpression := "0 0 0 1/1 * ? XXX" // Invalid expression

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCRONTriggerConfigBasic(pipeName, cronExpression, 2),
				ExpectError: regexp.MustCompile(`while validating expression:.*'XXX'.*`),
			},
		},
	})
}

func testAccCRONTriggerConfigBasic(pipeName string, cronExpression string, count int) string {
	triggers := ""
	for i := 1; i <= count; i++ {
		triggers += fmt.Sprintf(`
resource "spinnaker_pipeline_cron_trigger" "t%v" {
	pipeline        = "${spinnaker_pipeline.test.id}"
	cron_expression = "%s"
}`, i, cronExpression)
	}

	return testAccPipelineConfigBasic("app", pipeName) + triggers
}
