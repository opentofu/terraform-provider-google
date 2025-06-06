// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
// ----------------------------------------------------------------------------
//
//	***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
//
// ----------------------------------------------------------------------------
//
//	This code is generated by Magic Modules using the following:
//
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/cloudids/resource_cloudids_endpoint_test.go
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package cloudids_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-provider-google/google/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccCloudIdsEndpoint_basic(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"network_name":  acctest.BootstrapSharedServiceNetworkingConnection(t, "cloud-ids-endpoint-1"),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudIdsEndpointDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testCloudIds_basic(context),
			},
			{
				ResourceName:      "google_cloud_ids_endpoint.endpoint",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testCloudIds_basicUpdate(context),
			},
			{
				ResourceName:      "google_cloud_ids_endpoint.endpoint",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testCloudIds_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_compute_network" "default" {
  name = "%{network_name}"
}
  
resource "google_cloud_ids_endpoint" "endpoint" {
  name              = "cloud-ids-test-%{random_suffix}"
  location          = "us-central1-f"
  network           = data.google_compute_network.default.id
  severity          = "INFORMATIONAL"
  threat_exceptions = ["12", "67"]
}
`, context)
}

func testCloudIds_basicUpdate(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_compute_network" "default" {
  name = "%{network_name}"
}
  
resource "google_cloud_ids_endpoint" "endpoint" {
  name              = "cloud-ids-test-%{random_suffix}"
  location          = "us-central1-f"
  network           = data.google_compute_network.default.id
  severity          = "INFORMATIONAL"
  threat_exceptions = ["33"]
}
`, context)
}

func testAccCheckCloudIdsEndpointDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_cloud_ids_endpoint" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{CloudIdsBasePath}}projects/{{project}}/locations/{{location}}/endpoints/{{name}}")
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
				return fmt.Errorf("CloudIdsEndpoint still exists at %s", url)
			}
		}

		return nil
	}
}
