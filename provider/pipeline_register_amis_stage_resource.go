package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Coveo specific

func pipelineRegisterAmisStageResource() *schema.Resource {
	newRegisterAmisStageInterface := func() stage {
		return newRegisterAmisStage()
	}

	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageCreate(d, m, newRegisterAmisStageInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageRead(d, m, newRegisterAmisStageInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageUpdate(d, m, newRegisterAmisStageInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageDelete(d, m, newRegisterAmisStageInterface)
		},
		Importer: &schema.ResourceImporter{
			State: resourcePipelineImporter,
		},

		Schema: stageResource(map[string]*schema.Schema{
			"hardening_check_ref_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"virus_free_check_ref_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vulnerability_free_check_ref_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"package_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ami_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"source_ami": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_golden": {
				Type:     schema.TypeString,
				Required: true,
			},
			"fail_on_no_ami_to_register": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"pipeline_metadata": {
				Type:     schema.TypeMap,
				Optional: true,
			},
		}),
	}
}
