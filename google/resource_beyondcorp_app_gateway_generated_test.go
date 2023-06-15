// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccBeyondcorpAppGateway_beyondcorpAppGatewayBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBeyondcorpAppGatewayDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBeyondcorpAppGateway_beyondcorpAppGatewayBasicExample(context),
			},
			{
				ResourceName:            "google_beyondcorp_app_gateway.app_gateway",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "region"},
			},
		},
	})
}

func testAccBeyondcorpAppGateway_beyondcorpAppGatewayBasicExample(context map[string]interface{}) string {
	return tpgresource.Nprintf(`
resource "google_beyondcorp_app_gateway" "app_gateway" {
  name = "tf-test-my-app-gateway%{random_suffix}"
  type = "TCP_PROXY"
  region = "us-central1"
  host_type = "GCP_REGIONAL_MIG"
}
`, context)
}

func TestAccBeyondcorpAppGateway_beyondcorpAppGatewayFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBeyondcorpAppGatewayDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBeyondcorpAppGateway_beyondcorpAppGatewayFullExample(context),
			},
			{
				ResourceName:            "google_beyondcorp_app_gateway.app_gateway",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "region"},
			},
		},
	})
}

func testAccBeyondcorpAppGateway_beyondcorpAppGatewayFullExample(context map[string]interface{}) string {
	return tpgresource.Nprintf(`
resource "google_beyondcorp_app_gateway" "app_gateway" {
  name = "tf-test-my-app-gateway%{random_suffix}"
  type = "TCP_PROXY"
  region = "us-central1"
  display_name = "some display name%{random_suffix}"
  labels = {
    foo = "bar"
    bar = "baz"
  }
  host_type = "GCP_REGIONAL_MIG"
}
`, context)
}

func testAccCheckBeyondcorpAppGatewayDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_beyondcorp_app_gateway" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{BeyondcorpBasePath}}projects/{{project}}/locations/{{region}}/appGateways/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("BeyondcorpAppGateway still exists at %s", url)
			}
		}

		return nil
	}
}
