---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/firebaseappcheck/PlayIntegrityConfig.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Firebase App Check"
description: |-
  An app's Play Integrity configuration object.
---

# google_firebase_app_check_play_integrity_config

An app's Play Integrity configuration object. Note that your registered SHA-256 certificate fingerprints are used to validate tokens issued by the Play Integrity API.
Make sure your `google_firebase_android_app` has at least one `sha256_hashes` present.


To get more information about PlayIntegrityConfig, see:

* [API documentation](https://firebase.google.com/docs/reference/appcheck/rest/v1/projects.apps.playIntegrityConfig)
* How-to Guides
    * [Official Documentation](https://firebase.google.com/docs/app-check)

## Example Usage - Firebase App Check Play Integrity Config Minimal


```hcl
# Enables the Play Integrity API
resource "google_project_service" "play_integrity" {
  provider = google-beta

  project = "my-project-name"
  service = "playintegrity.googleapis.com"

  # Don't disable the service if the resource block is removed by accident.
  disable_on_destroy = false
}

resource "google_firebase_android_app" "default" {
  provider = google-beta

  project       = "my-project-name"
  display_name  = "Play Integrity app"
  package_name  = "package.name.playintegrity"
  sha1_hashes   = ["2145bdf698b8715039bd0e83f2069bed435ac21c"]
  sha256_hashes = ["2145bdf698b8715039bd0e83f2069bed435ac21ca1b2c3d4e5f6123456789abc"]
}

# It takes a while for App Check to recognize the new app
# If your app already exists, you don't have to wait 30 seconds.
resource "time_sleep" "wait_30s" {
  depends_on      = [google_firebase_android_app.default]
  create_duration = "30s"
}

resource "google_firebase_app_check_play_integrity_config" "default" {
  provider = google-beta

  project = "my-project-name"
  app_id  = google_firebase_android_app.default.app_id

  depends_on = [time_sleep.wait_30s]

  lifecycle {
    precondition {
      condition     = length(google_firebase_android_app.default.sha256_hashes) > 0
      error_message = "Provide a SHA-256 certificate on the Android App to use App Check"
    }
  }
}
```
## Example Usage - Firebase App Check Play Integrity Config Full


```hcl
# Enables the Play Integrity API
resource "google_project_service" "play_integrity" {
  provider = google-beta

  project = "my-project-name"
  service = "playintegrity.googleapis.com"

  # Don't disable the service if the resource block is removed by accident.
  disable_on_destroy = false
}

resource "google_firebase_android_app" "default" {
  provider = google-beta

  project       = "my-project-name"
  display_name  = "Play Integrity app"
  package_name  = "package.name.playintegrity"
  sha1_hashes   = ["2145bdf698b8715039bd0e83f2069bed435ac21c"]
  sha256_hashes = ["2145bdf698b8715039bd0e83f2069bed435ac21ca1b2c3d4e5f6123456789abc"]
}

# It takes a while for App Check to recognize the new app
# If your app already exists, you don't have to wait 30 seconds.
resource "time_sleep" "wait_30s" {
  depends_on      = [google_firebase_android_app.default]
  create_duration = "30s"
}

resource "google_firebase_app_check_play_integrity_config" "default" {
  provider = google-beta

  project   = "my-project-name"
  app_id    = google_firebase_android_app.default.app_id
  token_ttl = "7200s"

  depends_on = [time_sleep.wait_30s]

  lifecycle {
    precondition {
      condition     = length(google_firebase_android_app.default.sha256_hashes) > 0
      error_message = "Provide a SHA-256 certificate on the Android App to use App Check"
    }
  }
}
```

## Argument Reference

The following arguments are supported:


* `app_id` -
  (Required)
  The ID of an
  [Android App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.androidApps#AndroidApp.FIELDS.app_id).


* `token_ttl` -
  (Optional)
  Specifies the duration for which App Check tokens exchanged from Play Integrity artifacts will be valid.
  If unset, a default value of 1 hour is assumed. Must be between 30 minutes and 7 days, inclusive.
  A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/apps/{{app_id}}/playIntegrityConfig`

* `name` -
  The relative resource name of the Play Integrity configuration object


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


PlayIntegrityConfig can be imported using any of these accepted formats:

* `projects/{{project}}/apps/{{app_id}}/playIntegrityConfig`
* `{{project}}/{{app_id}}`
* `{{app_id}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import PlayIntegrityConfig using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/apps/{{app_id}}/playIntegrityConfig"
  to = google_firebase_app_check_play_integrity_config.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), PlayIntegrityConfig can be imported using one of the formats above. For example:

```
$ terraform import google_firebase_app_check_play_integrity_config.default projects/{{project}}/apps/{{app_id}}/playIntegrityConfig
$ terraform import google_firebase_app_check_play_integrity_config.default {{project}}/{{app_id}}
$ terraform import google_firebase_app_check_play_integrity_config.default {{app_id}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
