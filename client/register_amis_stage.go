package client

import (
	"github.com/mitchellh/mapstructure"
)

// RegisterAmisStageType bake stage
var RegisterAmisStageType StageType = "registerAmis"

func init() {
	stageFactories[RegisterAmisStageType] = parseRegisterAmisStage
}

// RegisterAmisStage for pipeline
type RegisterAmisStage struct {
	BaseStage `mapstructure:",squash"`

	HardeningCheckRefId         string `json:"hardeningCheckRefId"`
	VirusFreeCheckRefId         string `json:"virusFreeCheckRefId"`
	VulnerabilityFreeCheckRefId string `json:"vulnerabilityFreeCheckRefId"`
	PackageName                 string `json:"packageName"`
	AmiName                     string `json:"amiName"`
	SourceAmi                   string `json:"sourceAmi"`
	IsGolden                    string `json:"isGolden"`
	FailOnNoAmiToRegister       bool   `json:"failOnNoAmiToRegister"`

	PipelineMetadata map[string]string `json:"pipeline_metadata" mapstructure:"pipeline_metadata"`
}

// NewRegisterAmisStage for pipeline
func NewRegisterAmisStage() *RegisterAmisStage {
	return &RegisterAmisStage{
		BaseStage: *newBaseStage(RegisterAmisStageType),
	}
}

func parseRegisterAmisStage(stageMap map[string]interface{}) (Stage, error) {
	stage := NewRegisterAmisStage()
	if err := stage.parseBaseStage(stageMap); err != nil {
		return nil, err
	}

	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}
	return stage, nil
}
