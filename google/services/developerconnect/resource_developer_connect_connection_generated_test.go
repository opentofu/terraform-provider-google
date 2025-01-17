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

package developerconnect_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccDeveloperConnectConnection_developerConnectConnectionExistingCredentialsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"secret_name":   "projects/devconnect-terraform-creds/secrets/tf-test-do-not-change-github-oauthtoken-e0b9e7/versions/1",
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDeveloperConnectConnectionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDeveloperConnectConnection_developerConnectConnectionExistingCredentialsExample(context),
			},
			{
				ResourceName:            "google_developer_connect_connection.my-connection",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations", "connection_id", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccDeveloperConnectConnection_developerConnectConnectionExistingCredentialsExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_developer_connect_connection" "my-connection" {
  location = "us-central1"
  connection_id = "tf-test-tf-test-connection-cred%{random_suffix}"

  github_config {
    github_app = "DEVELOPER_CONNECT"

    authorizer_credential {
      oauth_token_secret_version = "%{secret_name}"
    }
  }
}

output "next_steps" {
  description = "Follow the action_uri if present to continue setup"
  value = google_developer_connect_connection.my-connection.installation_state
}
`, context)
}

func TestAccDeveloperConnectConnection_developerConnectConnectionGithubExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDeveloperConnectConnectionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDeveloperConnectConnection_developerConnectConnectionGithubExample(context),
			},
			{
				ResourceName:            "google_developer_connect_connection.my-connection",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations", "connection_id", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccDeveloperConnectConnection_developerConnectConnectionGithubExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_developer_connect_connection" "my-connection" {
  location = "us-central1"
  connection_id = "tf-test-tf-test-connection%{random_suffix}"

  github_config {
    github_app = "DEVELOPER_CONNECT"

    authorizer_credential {
      oauth_token_secret_version = "projects/devconnect-terraform-creds/secrets/tf-test-do-not-change-github-oauthtoken-e0b9e7/versions/1"
    }
  }
}
`, context)
}

func TestAccDeveloperConnectConnection_developerConnectConnectionGithubEnterpriseExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDeveloperConnectConnectionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDeveloperConnectConnection_developerConnectConnectionGithubEnterpriseExample(context),
			},
			{
				ResourceName:            "google_developer_connect_connection.my-connection",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations", "connection_id", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccDeveloperConnectConnection_developerConnectConnectionGithubEnterpriseExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_developer_connect_connection" "my-connection" {
  location = "us-central1"
  connection_id = "tf-test-tf-test-connection%{random_suffix}"

  github_enterprise_config {
    host_uri = "https://ghe.proctor-staging-test.com"
    app_id = 864434
    private_key_secret_version = "projects/devconnect-terraform-creds/secrets/tf-test-ghe-do-not-change-ghe-private-key-f522d2/versions/latest"
    webhook_secret_secret_version = "projects/devconnect-terraform-creds/secrets/tf-test-ghe-do-not-change-ghe-webhook-secret-3c806f/versions/latest"
    app_installation_id = 837537
  }
}
`, context)
}

func TestAccDeveloperConnectConnection_developerConnectConnectionGitlabExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDeveloperConnectConnectionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDeveloperConnectConnection_developerConnectConnectionGitlabExample(context),
			},
			{
				ResourceName:            "google_developer_connect_connection.my-connection",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations", "connection_id", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccDeveloperConnectConnection_developerConnectConnectionGitlabExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_developer_connect_connection" "my-connection" {
  location = "us-central1"
  connection_id = "tf-test-tf-test-connection%{random_suffix}"

  gitlab_config {
    webhook_secret_secret_version = "projects/devconnect-terraform-creds/secrets/gitlab-webhook/versions/latest"
    
    read_authorizer_credential {
      user_token_secret_version = "projects/devconnect-terraform-creds/secrets/gitlab-read-cred/versions/latest"
    }
    
    authorizer_credential {
      user_token_secret_version = "projects/devconnect-terraform-creds/secrets/gitlab-auth-cred/versions/latest"
    }
  }
}
`, context)
}

func TestAccDeveloperConnectConnection_developerConnectConnectionGitlabEnterpriseExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDeveloperConnectConnectionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDeveloperConnectConnection_developerConnectConnectionGitlabEnterpriseExample(context),
			},
			{
				ResourceName:            "google_developer_connect_connection.my-connection",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations", "connection_id", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccDeveloperConnectConnection_developerConnectConnectionGitlabEnterpriseExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_developer_connect_connection" "my-connection" {
  location = "us-central1"
  connection_id = "tf-test-tf-test-connection%{random_suffix}"

  gitlab_enterprise_config {
    host_uri = "https://gle-us-central1.gcb-test.com"
    
    webhook_secret_secret_version = "projects/devconnect-terraform-creds/secrets/gitlab-enterprise-webhook/versions/latest"

    read_authorizer_credential {
      user_token_secret_version = "projects/devconnect-terraform-creds/secrets/gitlab-enterprise-read-cred/versions/latest"
    }

    authorizer_credential {
      user_token_secret_version = "projects/devconnect-terraform-creds/secrets/gitlab-enterprise-auth-cred/versions/latest"
    }
  }
}
`, context)
}

func testAccCheckDeveloperConnectConnectionDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_developer_connect_connection" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{DeveloperConnectBasePath}}projects/{{project}}/locations/{{location}}/connections/{{connection_id}}")
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
				return fmt.Errorf("DeveloperConnectConnection still exists at %s", url)
			}
		}

		return nil
	}
}