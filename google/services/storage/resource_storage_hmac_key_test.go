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
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/storage/resource_storage_hmac_key_test.go
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package storage_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
)

func TestAccStorageHmacKey_update(t *testing.T) {
	t.Parallel()

	saName := fmt.Sprintf("%v%v", "service-account", acctest.RandString(t, 10))
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckStorageHmacKeyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGoogleStorageHmacKeyBasic(saName, "ACTIVE"),
			},
			{
				ResourceName:            "google_storage_hmac_key.key",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"secret"},
			},
			{
				Config: testAccGoogleStorageHmacKeyBasic(saName, "INACTIVE"),
			},
			{
				ResourceName:            "google_storage_hmac_key.key",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"secret"},
			},
		},
	})
}

func testAccGoogleStorageHmacKeyBasic(saName, state string) string {
	return fmt.Sprintf(`
resource "google_service_account" "service_account" {
  account_id = "%s"
}

resource "google_storage_hmac_key" "key" {
	service_account_email = google_service_account.service_account.email
	state = "%s"
}
`, saName, state)
}
