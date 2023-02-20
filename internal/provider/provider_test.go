package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

var (
	// testAccProtoV6ProviderFactories are used to instantiate a provider during
	// acceptance testing. The factory function will be invoked for every Terraform
	// CLI command executed to create a provider server to which the CLI can
	// reattach.
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"sdwan": providerserver.NewProtocol6WithError(New()),
	}
)

func testAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.
	if v := os.Getenv("SDWAN_USERNAME"); v == "" {
		t.Fatal("SDWAN_USERNAME env variable must be set for acceptance tests")
	}
	if v := os.Getenv("SDWAN_PASSWORD"); v == "" {
		t.Fatal("SDWAN_PASSWORD env variable must be set for acceptance tests")
	}
	if v := os.Getenv("SDWAN_URL"); v == "" {
		t.Fatal("SDWAN_URL env variable must be set for acceptance tests")
	}
}
