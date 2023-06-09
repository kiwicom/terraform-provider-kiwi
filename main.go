package main

import (
	"context"
	"terraform-provider-kiwi/kiwi"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

// Provider documentation generation.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name kiwi

func main() {
	providerserver.Serve(context.Background(), kiwi.New, providerserver.ServeOpts{
		Address: "registry.terraform.io/kiwicom/kiwi",
	})
}
