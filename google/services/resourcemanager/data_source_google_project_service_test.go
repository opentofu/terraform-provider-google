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
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/resourcemanager/data_source_google_project_service_test.go
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package resourcemanager_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccDataSourceGoogleProjectService_basic(t *testing.T) {
	t.Parallel()

	org := envvar.GetTestOrgFromEnv(t)
	pid := fmt.Sprintf("tf-test-%d", acctest.RandInt(t))
	services := []string{"iam.googleapis.com", "cloudresourcemanager.googleapis.com"}
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGoogleProjectService_basic(services, pid, org),
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceGoogleProjectServiceCheck("data.google_project_service.foo"),
				),
			},
		},
	})
}

func testAccDataSourceGoogleProjectService_basic(services []string, pid, org string) string {
	return fmt.Sprintf(`
resource "google_project" "acceptance" {
  project_id = "%s"
  name       = "%s"
  org_id     = "%s"
  deletion_policy = "DELETE"
}

resource "google_project_service" "foo" {
  project = google_project.acceptance.project_id
  service = "%s"
}

data "google_project_service" "foo" {
  project = google_project.acceptance.project_id
  service = google_project_service.foo.service
}
`, pid, pid, org, services[0])
}

func testAccDataSourceGoogleProjectServiceCheck(datasourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ds, ok := s.RootModule().Resources[datasourceName]
		if !ok {
			return fmt.Errorf("root module has no resource called %s", datasourceName)
		}

		if ds.Primary.Attributes["id"] == "" {
			return fmt.Errorf("specified API service is not enabled")
		}

		return nil
	}
}
