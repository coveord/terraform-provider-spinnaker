package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

// Coveo specific

type registerAmisStage struct {
	baseStage `mapstructure:",squash"`

	HardeningCheckRefId         string `mapstructure:"hardening_check_ref_id"`
	VirusFreeCheckRefId         string `mapstructure:"virus_free_check_ref_id"`
	VulnerabilityFreeCheckRefId string `mapstructure:"vulnerability_free_check_ref_id"`
	PackageName                 string `mapstructure:"package_name"`
	AmiName                     string `mapstructure:"ami_name"`
	SourceAmi                   string `mapstructure:"source_ami"`
	IsGolden                    string `mapstructure:"is_golden"`
	FailOnNoAmiToRegister       bool   `mapstructure:"fail_on_no_ami_to_register"`

	PipelineMetadata map[string]string `mapstructure:"pipeline_metadata"`
}

func newRegisterAmisStage() *registerAmisStage {
	return &registerAmisStage{
		baseStage: *newBaseStage(),
	}
}

func (s *registerAmisStage) toClientStage(_ *client.Config, refID string) (client.Stage, error) {
	cs := client.NewRegisterAmisStage()
	err := s.baseToClientStage(&cs.BaseStage, refID, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	cs.HardeningCheckRefId = s.HardeningCheckRefId
	cs.VirusFreeCheckRefId = s.VirusFreeCheckRefId
	cs.VulnerabilityFreeCheckRefId = s.VulnerabilityFreeCheckRefId
	cs.PackageName = s.PackageName
	cs.AmiName = s.AmiName
	cs.SourceAmi = s.SourceAmi
	cs.IsGolden = s.IsGolden
	cs.FailOnNoAmiToRegister = s.FailOnNoAmiToRegister
	cs.PipelineMetadata = s.PipelineMetadata

	return cs, nil
}

func (*registerAmisStage) fromClientStage(cs client.Stage) (stage, error) {
	clientStage := cs.(*client.RegisterAmisStage)
	newStage := newRegisterAmisStage()
	err := newStage.baseFromClientStage(&clientStage.BaseStage, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	newStage.HardeningCheckRefId = clientStage.HardeningCheckRefId
	newStage.VirusFreeCheckRefId = clientStage.VirusFreeCheckRefId
	newStage.VulnerabilityFreeCheckRefId = clientStage.VulnerabilityFreeCheckRefId
	newStage.PackageName = clientStage.PackageName
	newStage.AmiName = clientStage.AmiName
	newStage.SourceAmi = clientStage.SourceAmi
	newStage.IsGolden = clientStage.IsGolden
	newStage.FailOnNoAmiToRegister = clientStage.FailOnNoAmiToRegister
	newStage.PipelineMetadata = clientStage.PipelineMetadata

	return newStage, nil
}

func (s *registerAmisStage) SetResourceData(d *schema.ResourceData) error {
	err := s.baseSetResourceData(d)
	if err != nil {
		return err
	}

	err = d.Set("hardening_check_ref_id", s.HardeningCheckRefId)
	if err != nil {
		return err
	}
	err = d.Set("virus_free_check_ref_id", s.VirusFreeCheckRefId)
	if err != nil {
		return err
	}
	err = d.Set("vulnerability_free_check_ref_id", s.VulnerabilityFreeCheckRefId)
	if err != nil {
		return err
	}
	err = d.Set("package_name", s.PackageName)
	if err != nil {
		return err
	}
	err = d.Set("ami_name", s.AmiName)
	if err != nil {
		return err
	}
	err = d.Set("source_ami", s.SourceAmi)
	if err != nil {
		return err
	}
	err = d.Set("is_golden", s.IsGolden)
	if err != nil {
		return err
	}
	err = d.Set("fail_on_no_ami_to_register", s.FailOnNoAmiToRegister)
	if err != nil {
		return err
	}
	return d.Set("pipeline_metadata", s.PipelineMetadata)
}
