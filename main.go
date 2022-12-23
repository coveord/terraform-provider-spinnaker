package main

import (
	"log"

	"github.com/hashicorp/terraform/plugin"
	"github.com/jgramoll/terraform-provider-spinnaker/provider"
)

func main() {
	// Terraform expects every line to start with "[LEVEL] " but default value of logger outputs timestamps
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider})
}
