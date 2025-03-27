// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package parametermanager_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
)

func TestAccParameterManagerParameter_labelsUpdate(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckParameterManagerParameterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterManagerParameter_withoutLabels(context),
			},
			{
				ResourceName:            "google_parameter_manager_parameter.parameter-with-labels",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "parameter_id", "terraform_labels"},
			},
			{
				Config: testAccParameterManagerParameter_labelsUpdate(context),
			},
			{
				ResourceName:            "google_parameter_manager_parameter.parameter-with-labels",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "parameter_id", "terraform_labels"},
			},
			{
				Config: testAccParameterManagerParameter_labelsUpdateOther(context),
			},
			{
				ResourceName:            "google_parameter_manager_parameter.parameter-with-labels",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "parameter_id", "terraform_labels"},
			},
			{
				Config: testAccParameterManagerParameter_withoutLabels(context),
			},
			{
				ResourceName:            "google_parameter_manager_parameter.parameter-with-labels",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "parameter_id", "terraform_labels"},
			},
		},
	})
}

func testAccParameterManagerParameter_withoutLabels(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_parameter_manager_parameter" "parameter-with-labels" {
  parameter_id = "tf_test_parameter%{random_suffix}"
  format = "JSON"
}
`, context)
}

func testAccParameterManagerParameter_labelsUpdate(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_parameter_manager_parameter" "parameter-with-labels" {
  parameter_id = "tf_test_parameter%{random_suffix}"
  format = "JSON"

  labels = {
    key1 = "val1"
    key2 = "val2"
    key3 = "val3"
    key4 = "val4"
    key5 = "val5"
  }
}
`, context)
}

func testAccParameterManagerParameter_labelsUpdateOther(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_parameter_manager_parameter" "parameter-with-labels" {
  parameter_id = "tf_test_parameter%{random_suffix}"
  format = "JSON"

  labels = {
    key1 = "val1"
    key2 = "updateval2"
    updatekey3 = "val3"
    updatekey4 = "updateval4"
    key6 = "val6"
  }
}
`, context)
}

func TestAccParameterManagerParameter_kmsKeyUpdate(t *testing.T) {
	t.Parallel()
	acctest.BootstrapIamMembers(t, []acctest.IamMember{
		{
			Member: "serviceAccount:service-{project_number}@gcp-sa-pm.iam.gserviceaccount.com",
			Role:   "roles/cloudkms.cryptoKeyEncrypterDecrypter",
		},
	})

	context := map[string]interface{}{
		"kms_key":       acctest.BootstrapKMSKeyWithPurposeInLocationAndName(t, "ENCRYPT_DECRYPT", "global", "tf-parameter-manager-managed-1").CryptoKey.Name,
		"kms_key_other": acctest.BootstrapKMSKeyWithPurposeInLocationAndName(t, "ENCRYPT_DECRYPT", "global", "tf-parameter-manager-managed-2").CryptoKey.Name,
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckParameterManagerParameterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterManagerParameter_withoutkmsKey(context),
			},
			{
				ResourceName:            "google_parameter_manager_parameter.parameter-with-kms-key",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "parameter_id", "terraform_labels"},
			},
			{
				Config: testAccParameterManagerParameter_kmsKeyUpdate(context),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("google_parameter_manager_parameter.parameter-with-kms-key", plancheck.ResourceActionUpdate),
					},
				},
			},
			{
				ResourceName:            "google_parameter_manager_parameter.parameter-with-kms-key",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "parameter_id", "terraform_labels"},
			},
			{
				Config: testAccParameterManagerParameter_kmsKeyUpdateOther(context),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("google_parameter_manager_parameter.parameter-with-kms-key", plancheck.ResourceActionUpdate),
					},
				},
			},
			{
				ResourceName:            "google_parameter_manager_parameter.parameter-with-kms-key",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "parameter_id", "terraform_labels"},
			},
			{
				Config: testAccParameterManagerParameter_withoutkmsKey(context),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("google_parameter_manager_parameter.parameter-with-kms-key", plancheck.ResourceActionUpdate),
					},
				},
			},
			{
				ResourceName:            "google_parameter_manager_parameter.parameter-with-kms-key",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "parameter_id", "terraform_labels"},
			},
		},
	})
}

func testAccParameterManagerParameter_withoutkmsKey(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {}

resource "google_parameter_manager_parameter" "parameter-with-kms-key" {
  parameter_id = "tf_test_parameter%{random_suffix}"
  format = "JSON"
}
`, context)
}

func testAccParameterManagerParameter_kmsKeyUpdate(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {}

resource "google_parameter_manager_parameter" "parameter-with-kms-key" {
  parameter_id = "tf_test_parameter%{random_suffix}"
  format = "JSON"

  kms_key = "%{kms_key}"
}
`, context)
}

func testAccParameterManagerParameter_kmsKeyUpdateOther(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {}

resource "google_parameter_manager_parameter" "parameter-with-kms-key" {
  parameter_id = "tf_test_parameter%{random_suffix}"
  format = "JSON"

  kms_key = "%{kms_key_other}"
}
`, context)
}
