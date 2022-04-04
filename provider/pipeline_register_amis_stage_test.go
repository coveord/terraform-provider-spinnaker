package provider

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

func init() {
	stageTypes["spinnaker_pipeline_register_amis_stage"] = client.RegisterAmisStageType
}

type RegisterAmisStageTestCase struct {
	HardeningCheckRefId         string
	VirusFreeCheckRefId         string
	VulnerabilityFreeCheckRefId string
	PackageName                 string
	AmiName                     string
	SourceAmi                   string
	IsGolden                    string
	FailOnNoAmiToRegister       bool
	PipelineMetadataValue       string
}

func TestAccPipelineRegisterAmisStageBasic(t *testing.T) {
	var pipelineRef client.Pipeline
	var stages []client.Stage
	pipeName := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

	createTestCase := RegisterAmisStageTestCase{
		HardeningCheckRefId:         "1",
		VirusFreeCheckRefId:         "2",
		VulnerabilityFreeCheckRefId: "3",
		PackageName:                 "package",
		AmiName:                     "tonAmi",
		SourceAmi:                   "ami-trouveaujeancoutu",
		IsGolden:                    "nah",
		FailOnNoAmiToRegister:       true,
		PipelineMetadataValue:       "benOui",
	}
	modifyTestCase := RegisterAmisStageTestCase{
		HardeningCheckRefId:         "4",
		VirusFreeCheckRefId:         "5",
		VulnerabilityFreeCheckRefId: "6",
		PackageName:                 "newPackage",
		AmiName:                     "tonAmiImaginaireFinalement",
		SourceAmi:                   "ami-trouveauxobjetsperdus",
		IsGolden:                    "maybe",
		FailOnNoAmiToRegister:       false,
		PipelineMetadataValue:       "benNon",
	}

	pipelineResourceName := "spinnaker_pipeline.test"
	stage1 := "spinnaker_pipeline_register_amis_stage.s1"
	stage2 := "spinnaker_pipeline_register_amis_stage.s2"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelineRegisterAmisStageConfigBasic(pipeName, &createTestCase, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "hardening_check_ref_id", createTestCase.HardeningCheckRefId),
					resource.TestCheckResourceAttr(stage1, "virus_free_check_ref_id", createTestCase.VirusFreeCheckRefId),
					resource.TestCheckResourceAttr(stage1, "vulnerability_free_check_ref_id", createTestCase.VulnerabilityFreeCheckRefId),
					resource.TestCheckResourceAttr(stage1, "package_name", createTestCase.PackageName),
					resource.TestCheckResourceAttr(stage1, "ami_name", createTestCase.AmiName),
					resource.TestCheckResourceAttr(stage1, "source_ami", createTestCase.SourceAmi),
					resource.TestCheckResourceAttr(stage1, "is_golden", createTestCase.IsGolden),
					resource.TestCheckResourceAttr(stage1, "fail_on_no_ami_to_register", strconv.FormatBool(createTestCase.FailOnNoAmiToRegister)),
					resource.TestCheckResourceAttr(stage1, "pipeline_metadata.key", createTestCase.PipelineMetadataValue+"1"),
					resource.TestCheckResourceAttr(stage2, "name", "Stage 2"),
					resource.TestCheckResourceAttr(stage2, "hardening_check_ref_id", createTestCase.HardeningCheckRefId),
					resource.TestCheckResourceAttr(stage2, "virus_free_check_ref_id", createTestCase.VirusFreeCheckRefId),
					resource.TestCheckResourceAttr(stage2, "vulnerability_free_check_ref_id", createTestCase.VulnerabilityFreeCheckRefId),
					resource.TestCheckResourceAttr(stage2, "package_name", createTestCase.PackageName),
					resource.TestCheckResourceAttr(stage2, "ami_name", createTestCase.AmiName),
					resource.TestCheckResourceAttr(stage2, "source_ami", createTestCase.SourceAmi),
					resource.TestCheckResourceAttr(stage2, "is_golden", createTestCase.IsGolden),
					resource.TestCheckResourceAttr(stage2, "fail_on_no_ami_to_register", strconv.FormatBool(createTestCase.FailOnNoAmiToRegister)),
					resource.TestCheckResourceAttr(stage2, "pipeline_metadata.key", createTestCase.PipelineMetadataValue+"2"),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineStages(pipelineResourceName, []string{
						stage1,
						stage2,
					}, &stages),
				),
			},
			{
				ResourceName:  stage1,
				ImportStateId: "invalid",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(`Invalid import key, must be pipelineID_stageID`),
			},
			{
				ResourceName: stage1,
				ImportState:  true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					if len(stages) == 0 {
						return "", fmt.Errorf("no stages to import")
					}
					return fmt.Sprintf("%s_%s", pipelineRef.ID, stages[0].GetRefID()), nil
				},
				ImportStateVerify: true,
			},
			{
				ResourceName: stage2,
				ImportState:  true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					if len(stages) < 2 {
						return "", fmt.Errorf("no stages to import")
					}
					return fmt.Sprintf("%s_%s", pipelineRef.ID, stages[1].GetRefID()), nil
				},
				ImportStateVerify: true,
			},
			{
				Config: testAccPipelineRegisterAmisStageConfigBasic(pipeName, &modifyTestCase, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "hardening_check_ref_id", modifyTestCase.HardeningCheckRefId),
					resource.TestCheckResourceAttr(stage1, "virus_free_check_ref_id", modifyTestCase.VirusFreeCheckRefId),
					resource.TestCheckResourceAttr(stage1, "vulnerability_free_check_ref_id", modifyTestCase.VulnerabilityFreeCheckRefId),
					resource.TestCheckResourceAttr(stage1, "package_name", modifyTestCase.PackageName),
					resource.TestCheckResourceAttr(stage1, "ami_name", modifyTestCase.AmiName),
					resource.TestCheckResourceAttr(stage1, "source_ami", modifyTestCase.SourceAmi),
					resource.TestCheckResourceAttr(stage1, "is_golden", modifyTestCase.IsGolden),
					resource.TestCheckResourceAttr(stage1, "fail_on_no_ami_to_register", strconv.FormatBool(modifyTestCase.FailOnNoAmiToRegister)),
					resource.TestCheckResourceAttr(stage1, "pipeline_metadata.key", modifyTestCase.PipelineMetadataValue+"1"),
					resource.TestCheckResourceAttr(stage2, "name", "Stage 2"),
					resource.TestCheckResourceAttr(stage2, "hardening_check_ref_id", modifyTestCase.HardeningCheckRefId),
					resource.TestCheckResourceAttr(stage2, "virus_free_check_ref_id", modifyTestCase.VirusFreeCheckRefId),
					resource.TestCheckResourceAttr(stage2, "vulnerability_free_check_ref_id", modifyTestCase.VulnerabilityFreeCheckRefId),
					resource.TestCheckResourceAttr(stage2, "package_name", modifyTestCase.PackageName),
					resource.TestCheckResourceAttr(stage2, "ami_name", modifyTestCase.AmiName),
					resource.TestCheckResourceAttr(stage2, "source_ami", modifyTestCase.SourceAmi),
					resource.TestCheckResourceAttr(stage2, "is_golden", modifyTestCase.IsGolden),
					resource.TestCheckResourceAttr(stage2, "fail_on_no_ami_to_register", strconv.FormatBool(modifyTestCase.FailOnNoAmiToRegister)),
					resource.TestCheckResourceAttr(stage2, "pipeline_metadata.key", modifyTestCase.PipelineMetadataValue+"2"),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineStages(pipelineResourceName, []string{
						stage1,
						stage2,
					}, &stages),
				),
			},
			{
				Config: testAccPipelineRegisterAmisStageConfigBasic(pipeName, nil, 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineStages(pipelineResourceName, []string{}, &stages),
				),
			},
		},
	})
}

func testAccPipelineRegisterAmisStageConfigBasic(pipeName string, testCase *RegisterAmisStageTestCase, count int) string {
	stages := ""
	for i := 1; i <= count; i++ {
		stages += fmt.Sprintf(`
resource "spinnaker_pipeline_register_amis_stage" "s%v" {
    pipeline = "${spinnaker_pipeline.test.id}"
    name     = "Stage %v"
	
    hardening_check_ref_id = "%v"
    virus_free_check_ref_id = "%v"
    vulnerability_free_check_ref_id = "%v"
    package_name = "%v"
    ami_name = "%v"
    source_ami = "%v"
    is_golden = "%v"
    fail_on_no_ami_to_register = "%v"

	pipeline_metadata = {
        key = "%v%v"	
    }
}`, i, i, testCase.HardeningCheckRefId, testCase.VirusFreeCheckRefId, testCase.VulnerabilityFreeCheckRefId,
			testCase.PackageName, testCase.AmiName, testCase.SourceAmi, testCase.IsGolden,
			testCase.FailOnNoAmiToRegister, testCase.PipelineMetadataValue, i)
	}

	return testAccPipelineConfigBasic("app", pipeName) + stages
}
