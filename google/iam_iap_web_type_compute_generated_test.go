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

	"github.com/hashicorp/terraform-provider-google/google/acctest"
)

func TestAccIapWebTypeComputeIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/iap.httpsResourceAccessor",
		"org_id":        acctest.GetTestOrgFromEnv(t),

		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebTypeComputeIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_compute_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute roles/iap.httpsResourceAccessor", fmt.Sprintf("tf-test%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccIapWebTypeComputeIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_compute_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute roles/iap.httpsResourceAccessor", fmt.Sprintf("tf-test%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebTypeComputeIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/iap.httpsResourceAccessor",
		"org_id":        acctest.GetTestOrgFromEnv(t),

		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccIapWebTypeComputeIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_compute_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute roles/iap.httpsResourceAccessor user:admin@hashicorptest.com", fmt.Sprintf("tf-test%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebTypeComputeIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/iap.httpsResourceAccessor",
		"org_id":        acctest.GetTestOrgFromEnv(t),

		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebTypeComputeIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_iap_web_type_compute_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_iap_web_type_compute_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute", fmt.Sprintf("tf-test%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccIapWebTypeComputeIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_iap_web_type_compute_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute", fmt.Sprintf("tf-test%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebTypeComputeIamBindingGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/iap.httpsResourceAccessor",
		"org_id":        acctest.GetTestOrgFromEnv(t),

		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebTypeComputeIamBinding_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_compute_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute roles/iap.httpsResourceAccessor %s", fmt.Sprintf("tf-test%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebTypeComputeIamBindingGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/iap.httpsResourceAccessor",
		"org_id":        acctest.GetTestOrgFromEnv(t),

		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebTypeComputeIamBinding_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_compute_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute roles/iap.httpsResourceAccessor", fmt.Sprintf("tf-test%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_iap_web_type_compute_iam_binding.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute roles/iap.httpsResourceAccessor %s", fmt.Sprintf("tf-test%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_iap_web_type_compute_iam_binding.foo3",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute roles/iap.httpsResourceAccessor %s", fmt.Sprintf("tf-test%s", context["random_suffix"]), context["condition_title_no_desc"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebTypeComputeIamMemberGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/iap.httpsResourceAccessor",
		"org_id":        acctest.GetTestOrgFromEnv(t),

		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebTypeComputeIamMember_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_compute_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute roles/iap.httpsResourceAccessor user:admin@hashicorptest.com %s", fmt.Sprintf("tf-test%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebTypeComputeIamMemberGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/iap.httpsResourceAccessor",
		"org_id":        acctest.GetTestOrgFromEnv(t),

		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebTypeComputeIamMember_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_compute_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute roles/iap.httpsResourceAccessor user:admin@hashicorptest.com", fmt.Sprintf("tf-test%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_iap_web_type_compute_iam_member.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute roles/iap.httpsResourceAccessor user:admin@hashicorptest.com %s", fmt.Sprintf("tf-test%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_iap_web_type_compute_iam_member.foo3",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute roles/iap.httpsResourceAccessor user:admin@hashicorptest.com %s", fmt.Sprintf("tf-test%s", context["random_suffix"]), context["condition_title_no_desc"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebTypeComputeIamPolicyGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/iap.httpsResourceAccessor",
		"org_id":        acctest.GetTestOrgFromEnv(t),

		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	// Test should have 2 bindings: one with a description and one without. Any < chars are converted to a unicode character by the API.
	expectedPolicyData := Nprintf(`{"bindings":[{"condition":{"description":"%{condition_desc}","expression":"%{condition_expr}","title":"%{condition_title}"},"members":["user:admin@hashicorptest.com"],"role":"%{role}"},{"condition":{"expression":"%{condition_expr}","title":"%{condition_title}-no-description"},"members":["user:admin@hashicorptest.com"],"role":"%{role}"}]}`, context)
	expectedPolicyData = strings.Replace(expectedPolicyData, "<", "\\u003c", -1)

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebTypeComputeIamPolicy_withConditionGenerated(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					// TODO(SarahFrench) - uncomment once https://github.com/GoogleCloudPlatform/magic-modules/pull/6466 merged
					// resource.TestCheckResourceAttr("data.google_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttr("google_iap_web_type_compute_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttrWith("data.google_iam_policy.foo", "policy_data", checkGoogleIamPolicy),
				),
			},
			{
				ResourceName:      "google_iap_web_type_compute_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute", fmt.Sprintf("tf-test%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccIapWebTypeComputeIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  project_id = "tf-test%{random_suffix}"
  name       = "tf-test%{random_suffix}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}

resource "google_iap_web_type_compute_iam_member" "foo" {
  project = google_project_service.project_service.project
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccIapWebTypeComputeIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  project_id = "tf-test%{random_suffix}"
  name       = "tf-test%{random_suffix}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_iap_web_type_compute_iam_policy" "foo" {
  project = google_project_service.project_service.project
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_iap_web_type_compute_iam_policy" "foo" {
  project = google_project_service.project_service.project
  depends_on = [
    google_iap_web_type_compute_iam_policy.foo
  ]
}
`, context)
}

func testAccIapWebTypeComputeIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  project_id = "tf-test%{random_suffix}"
  name       = "tf-test%{random_suffix}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}

data "google_iam_policy" "foo" {
}

resource "google_iap_web_type_compute_iam_policy" "foo" {
  project = google_project_service.project_service.project
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccIapWebTypeComputeIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  project_id = "tf-test%{random_suffix}"
  name       = "tf-test%{random_suffix}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}

resource "google_iap_web_type_compute_iam_binding" "foo" {
  project = google_project_service.project_service.project
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccIapWebTypeComputeIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  project_id = "tf-test%{random_suffix}"
  name       = "tf-test%{random_suffix}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}

resource "google_iap_web_type_compute_iam_binding" "foo" {
  project = google_project_service.project_service.project
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}

func testAccIapWebTypeComputeIamBinding_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  project_id = "tf-test%{random_suffix}"
  name       = "tf-test%{random_suffix}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}

resource "google_iap_web_type_compute_iam_binding" "foo" {
  project = google_project_service.project_service.project
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccIapWebTypeComputeIamBinding_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  project_id = "tf-test%{random_suffix}"
  name       = "tf-test%{random_suffix}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}

resource "google_iap_web_type_compute_iam_binding" "foo" {
  project = google_project_service.project_service.project
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}

resource "google_iap_web_type_compute_iam_binding" "foo2" {
  project = google_project_service.project_service.project
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}

resource "google_iap_web_type_compute_iam_binding" "foo3" {
  project = google_project_service.project_service.project
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    # Check that lack of description doesn't cause any issues
    # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
    title       = "%{condition_title_no_desc}"
    expression  = "%{condition_expr_no_desc}"
  }
}
`, context)
}

func testAccIapWebTypeComputeIamMember_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  project_id = "tf-test%{random_suffix}"
  name       = "tf-test%{random_suffix}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}

resource "google_iap_web_type_compute_iam_member" "foo" {
  project = google_project_service.project_service.project
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccIapWebTypeComputeIamMember_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  project_id = "tf-test%{random_suffix}"
  name       = "tf-test%{random_suffix}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}

resource "google_iap_web_type_compute_iam_member" "foo" {
  project = google_project_service.project_service.project
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}

resource "google_iap_web_type_compute_iam_member" "foo2" {
  project = google_project_service.project_service.project
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}

resource "google_iap_web_type_compute_iam_member" "foo3" {
  project = google_project_service.project_service.project
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    # Check that lack of description doesn't cause any issues
    # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
    title       = "%{condition_title_no_desc}"
    expression  = "%{condition_expr_no_desc}"
  }
}
`, context)
}

func testAccIapWebTypeComputeIamPolicy_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  project_id = "tf-test%{random_suffix}"
  name       = "tf-test%{random_suffix}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
    condition {
      # Check that lack of description doesn't cause any issues
      # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
      title       = "%{condition_title_no_desc}"
      expression  = "%{condition_expr_no_desc}"
    }
  }
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
    condition {
      title       = "%{condition_title}"
      description = "%{condition_desc}"
      expression  = "%{condition_expr}"
    }
  }
}

resource "google_iap_web_type_compute_iam_policy" "foo" {
  project = google_project_service.project_service.project
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}
