---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/website/docs/guides/version_6_upgrade.html.markdown
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
page_title: "Terraform provider for Google Cloud 6.0.0 Upgrade Guide"
description: |-
  Terraform provider for Google Cloud 6.0.0 Upgrade Guide
---

# Terraform Google Provider 6.0.0 Upgrade Guide

The `6.0.0` release of the Google provider for Terraform is a major version and
includes some changes that you will need to consider when upgrading. This guide
is intended to help with that process and focuses only on the changes necessary
to upgrade from the final `5.X` series release to `6.0.0`.

Most of the changes outlined in this guide have been previously marked as
deprecated in the Terraform `plan`/`apply` output throughout previous provider
releases, up to and including the final `5.X` series release. These changes,
such as deprecation notices, can always be found in the CHANGELOG of the
affected providers. [google](https://github.com/hashicorp/terraform-provider-google/blob/main/CHANGELOG.md)
[google-beta](https://github.com/hashicorp/terraform-provider-google-beta/blob/main/CHANGELOG.md)

## I accidentally upgraded to 6.0.0, how do I downgrade to `5.X`?

If you've inadvertently upgraded to `6.0.0`, first see the
[Provider Version Configuration Guide](#provider-version-configuration) to lock
your provider version; if you've constrained the provider to a lower version
such as shown in the previous version example in that guide, Terraform will pull
in a `5.X` series release on `terraform init`.

If you've only ran `terraform init` or `terraform plan`, your state will not
have been modified and downgrading your provider is sufficient.

If you've ran `terraform refresh` or `terraform apply`, Terraform may have made
state changes in the meantime.

* If you're using a local state, or a remote state backend that does not support
versioning, `terraform refresh` with a downgraded provider is likely sufficient
to revert your state. The Google provider generally refreshes most state
information from the API, and the properties necessary to do so have been left
unchanged.

* If you're using a remote state backend that supports versioning such as
[Google Cloud Storage](https://developer.hashicorp.com/terraform/language/settings/backends/gcs),
you can revert the Terraform state file to a previous version. If you do
so and Terraform had created resources as part of a `terraform apply` in the
meantime, you'll need to either delete them by hand or `terraform import` them
so Terraform knows to manage them.

## Provider Version Configuration

-> Before upgrading to version 6.0.0, it is recommended to upgrade to the most
recent `5.X` series release of the provider, make the changes noted in this guide,
and ensure that your environment successfully runs
[`terraform plan`](https://developer.hashicorp.com/terraform/cli/commands/plan)
without unexpected changes or deprecation notices.

It is recommended to use [version constraints](https://developer.hashicorp.com/terraform/language/providers/requirements#requiring-providers)
when configuring Terraform providers. If you are following that recommendation,
update the version constraints in your Terraform configuration and run
[`terraform init`](https://developer.hashicorp.com/terraform/cli/commands/init) to download
the new version.

If you aren't using version constraints, you can use `terraform init -upgrade`
in order to upgrade your provider to the latest released version.

For example, given this previous configuration:

```hcl
terraform {
  required_providers {
    google = {
      version = "~> 5.30.0"
    }
  }
}
```

An updated configuration:

```hcl
terraform {
  required_providers {
    google = {
      version = "~> 6.0.0"
    }
  }
}
```

## Provider: Terraform provider attribution label is added to new resources by default

Version 5.16.0 introduced the `goog-terraform-provisioned = true` label that could
be automatically added to resources, making it easy to identify resources created
by the provider when using other tools such as `gcloud` or the GCP web console. In
5.16.0 the label needed to be enabled explicitly; in 6.0.0 the default is to add the
label to all newly created resources. This behavior can be disabled in the provider
configuration. For example:

```hcl
provider "google" {
  add_terraform_attribution_label = false
}
```

## Provider: `name_prefix` max length has been extended from 37 to 54 characters for multiple resources

Affected resources: `google_compute_instance_template`, `google_compute_region_instance_template`, `google_compute_ssl_certificate`,
and `google_compute_region_ssl_certificate`

Previously, the max length of `name_prefix` was 37 characters since the autogenerated UUID suffix was 26 characters which combined to
the total max length for names of 63 characters.
In 6.0.0, providing a `name_prefix` larger than 37 characters will prompt the provider to use a shortened suffix of only 9 characters, leading to a new max of 54 characters for `name_prefix`. This shortened suffix is inevitably more prone to collisions, so use the longer max `name_prefix` length with caution.

## Provider: Opt-out deletion protection is added to several resources, including google_project

Affected resources: `google_cloud_run_v2_job`, `google_cloud_run_v2_service`, `google_domain`, `google_folder`, 
`google_project`

Protection against Terraform deleting these resources has been added via a new field in each resource which is enabled by default. See the individual resource entries for the field names and other resource-specific details.

## Resource: `google_alloydb_cluster` 

### `network` is now removed

`network` has been removed in favor of `network_config.network`

## Resource: `google_bigquery_reservation`

### `multi_region_auxiliary` is now removed

This field is no longer supported by the BigQuery Reservation API.

## Resource: `google_bigquery_table`

### View creation now validates `schema`

A `view` can no longer be created when `schema` contains required fields

### `allow_resource_tags_on_deletion` is now removed

Resource tags are now always allowed on table deletion.

## Resource: `google_cloud_run_v2_job`

### Job deletion now prevented by default with `deletion_protection`

The field `deletion_protection` has been added with a default value of `true`. This field prevents
Terraform from destroying or recreating the Job. In 6.0.0, existing jobs will have
`deletion_protection` set to `true` during the next refresh unless otherwise set in configuration.

**`deletion_protection` does NOT prevent deletion outside of Terraform.**

To disable deletion protection, explicitly set this field to `false` in configuration
and then run `terraform apply` to apply the change.

### retyped `containers.env` to SET from ARRAY

Previously, `containers.env` was a list, making it order-dependent. It is now a set.

If you were relying on accessing an individual environment variable by index (for example, `google_cloud_run_v2_job.template.containers.0.env.0.name`), then that will now need to by hash (for example, `google_cloud_run_v2_job.template.containers.0.env.<some-hash>.name`).

## Resource: `google_cloud_run_v2_service`

### Service deletion now prevented by default with `deletion_protection`

The field `deletion_protection` has been added with a default value of `true`. This field prevents
Terraform from destroying or recreating the Service. In 6.0.0, existing services will have
`deletion_protection` set to `true` during the next refresh unless otherwise set in configuration.

**`deletion_protection` does NOT prevent deletion outside of Terraform.**

To disable deletion protection, explicitly set this field to `false` in configuration
and then run `terraform apply` to apply the change.

### `liveness_probe` no longer defaults from API

Cloud Run does not provide a default value for liveness probe. Now removing this field
will remove the liveness probe from the Cloud Run service.

### retyped `containers.env` to SET from ARRAY

Previously, `containers.env` was a list, making it order-dependent. It is now a set.

If you were relying on accessing an individual environment variable by index (for example, `google_cloud_run_v2_service.template.containers.0.env.0.name`), then that will now need to by hash (for example, `google_cloud_run_v2_service.template.containers.0.env.<some-hash>.name`).

## Resource: `google_composer_environment`

### `ip_allocation_policy = []` is no longer valid configuration

There was no functional difference between setting `ip_allocation_policy = []` and not setting `ip_allocation_policy` at all. Removing the field from configuration should not produce a diff.

## Resources: `google_compute_backend_service` and `google_compute_region_backend_service`

### `iap.enabled` is now required in the `iap` block

To apply the IAP settings to the backend service, `true` needs to be set for `enabled` field.

### `outlier_detection` subfields default values removed

Empty values mean the setting should be cleared.

### `connection_draining_timeout_sec` default value changed

An empty value now means 300.

### `balancing_mode` default value changed

An empty value now means UTILIZATION.

## Resources: `google_compute_instance`, `google_container_cluster`, and `google_container_node_pool` 

### `guest_accelerator = []` is no longer valid configuration

[Argument syntax](https://developer.hashicorp.com/terraform/language/syntax/configuration#arguments) is no longer supported for this field,
in favor of [block syntax](https://developer.hashicorp.com/terraform/language/syntax/configuration#blocks).
For configurations using argument syntax dynamically with variables, it is recommended to use
[dynamic blocks](https://developer.hashicorp.com/terraform/language/expressions/dynamic-blocks) instead.

By default, omitting `guest_accelerator` will lead to Terraform defaulting to the API's value.
To explicitly set an empty `guest_accelerator` list, define a `guest_accelerator` block with `guest_accelerator.count = 0`.
This is necessary to mirror the previous behavior of `guest_accelerator = []` in 5.X.

Argument syntax was previously enabled to maintain compatability in behavior between Terraform versions 0.11 and 0.12 using a special setting ["attributes as blocks"](https://developer.hashicorp.com/terraform/language/attr-as-blocks).
This special setting causes other breakages so it is now removed, with setting `guest_accelerator.count = 0` available as an alternative form of empty `guest_accelerator` object.

### `guest_accelerator.gpu_driver_installation_config = []` and `guest_accelerator.gpu_sharing_config = []` are no longer valid configuration

These were never intended to be set this way. Removing the fields from configuration should not produce a diff.

## Resources: `google_compute_instance_from_template` and `google_compute_instance_from_machine_image`

### `network_interface.alias_ip_range, network_interface.access_config, attached_disk, guest_accelerator, service_account, scratch_disk` can no longer be set to an empty block `[]`

`field = []` is no longer valid configuration for these fields. Removing the fields from configuration should not produce a diff.

## Resource: `google_compute_subnetwork`

### `secondary_ip_range = []` is no longer valid configuration

To explicitly set an empty list of objects, use `send_secondary_ip_range_if_empty = true` and completely remove `secondary_ip_range` from config.

Previously, to explicitly set `secondary_ip_range` as an empty list of objects, the specific configuration `secondary_ip_range = []` was necessary.
This was to maintain compatability in behavior between Terraform versions 0.11 and 0.12 using a special setting ["attributes as blocks"](https://developer.hashicorp.com/terraform/language/attr-as-blocks).
This special setting causes other breakages so it is now removed, with `send_secondary_ip_range_if_empty` available instead.

## Resource: `google_container_cluster`

### `advanced_datapath_observability_config.relay_mode` is now removed

Previously, through `relay_mode` field usage, users could both enable Dataplane V2
Flow Observability feature (that deploys Hubble relay component) and configure
managed load balancers. Due to users' needs to have better control over how
Hubble relay components shall be exposed in their clusters, managed load
balancer deployments are not supported anymore and users are expected to deploy
their own load balancers.

If `advanced_datapath_observability_config` is defined, `enable_relay` is now a
required field instead and users are expected to use this field instead.

Recommended migration from `relay_mode` to `enable_relay` depending on
`relay_mode` value:
* `DISABLED`: set `enable_relay` to `false`
* `INTERNAL_VPC_LB`: set `enable_relay` to `true` and define internal load
  balancer with VPC scope
* `EXTERNAL_LB`: set `enable_relay` to `true` and define external load balancer
  with public access

See exported endpoints for Dataplane V2 Observability feature to learn what
target you might wish to expose with load balancers:
https://cloud.google.com/kubernetes-engine/docs/concepts/about-dpv2-observability#gke-dataplane-v2-observability-endpoints

### Three label-related fields are now present

* `resource_labels` field is non-authoritative and only manages the labels defined by
the users on the resource through Terraform.
* The new output-only `terraform_labels` field merges the labels defined by the users
on the resource through Terraform and the default labels configured on the provider.
* The new output-only `effective_labels` field lists all of labels present on the resource
in GCP, including the labels configured through Terraform, the system, and other clients.

## Data source: `google_container_cluster`

### Three label-related fields are now present

All three of `resource_labels`, `effective_labels` and `terraform_labels` will now be present.
All of these three fields include all of the labels present on the resource in GCP including
the labels configured through Terraform, the system, and other clients, equivalent to
`effective_labels` on the resource.

## Resource: `google_datastore_index` is now removed

`google_datastore_index` is removed in favor of `google_firestore_index`

## Resource: `google_domain`

### Domain deletion now prevented by default with `deletion_protection`

The field `deletion_protection` has been added with a default value of `true`. This field prevents
Terraform from destroying or recreating the Domain. In 6.0.0, existing domains will have 
`deletion_protection` set to `true` during the next refresh unless otherwise set in configuration.

**`deletion_protection` does NOT prevent deletion outside of Terraform.**

To disable deletion protection, explicitly set this field to `false` in configuration
and then run `terraform apply` to apply the change.

## Resource: `google_edgenetwork_network`

### Three label-related fields are now present

* `labels` field is non-authoritative and only manages the labels defined by
the users on the resource through Terraform.
* The new output-only `terraform_labels` field merges the labels defined by the users
on the resource through Terraform and the default labels configured on the provider.
* The new output-only `effective_labels` field lists all of labels present on the resource
in GCP, including the labels configured through Terraform, the system, and other clients.

## Resource: `google_edgenetwork_subnet`

### Three label-related fields are now present

* `labels` field is non-authoritative and only manages the labels defined by
the users on the resource through Terraform.
* The new output-only `terraform_labels` field merges the labels defined by the users
on the resource through Terraform and the default labels configured on the provider.
* The new output-only `effective_labels` field lists all of labels present on the resource
in GCP, including the labels configured through Terraform, the system, and other clients.

## Resource: `google_folder`

### Folder deletion now prevented by default with `deletion_protection`

The field `deletion_protection` has been added with a default value of `true`. This field prevents
Terraform from destroying or recreating the Folder. In 6.0.0, existing folders will have
`deletion_protection` set to `true` during the next refresh unless otherwise set in configuration.

**`deletion_protection` does NOT prevent deletion outside of Terraform.**

To disable deletion protection, explicitly set this field to `false` in configuration
and then run `terraform apply` to apply the change.

## Resource: `google_identity_platform_project_default_config` is now removed

`google_identity_platform_project_default_config` is removed in favor of `google_identity_platform_project_config`

## Resource: `google_integrations_client`

### `create_sample_workflows` and `provision_gmek` is now removed

`create_sample_workflows` and `provision_gmek` is now removed in favor of `create_sample_integrations`

## Resource: `google_project`

### Project deletion now prevented by default with `deletion_policy`

The default value for `deletion_policy` is now `PREVENT` instead of `DELETE`. The `PREVENT` value for `deletion_policy` stops Terraform
from deleting or recreating your project. To remove deletion protection entirely,
explicitly set this field to `DELETE` in your configuration and run `terraform apply`.
Alternatively, setting this field to `ABANDON` allows Terraform to remove your project from state without destroying it.

### `skip_delete` is now removed

`skip_delete` has been removed in favor of `deletion_policy`. In order to get the same behavior as `skip_delete = true`, set `deletion_policy = ABANDON`. Be aware that `deletion_policy = ABANDON` will override the error-on-delete behaviour that the `PREVENT` policy applies.

## Resource: `google_pubsub_topic`

### `schema_settings` no longer has a default value

An empty value means the setting should be cleared.

## Resource: `google_redis_cluster`

### `deletion_protection_enabled` field with default value added

Support for the deletionProtectionEnabled field has been added. Redis clusters will now be created with a `deletion_protection_enabled = true` value by default. 

## Resource: `google_sql_database_instance`

### `settings.ip_configuration.require_ssl` is now removed (in 6.0.1)

Removed in favor of field `settings.ip_configuration.ssl_mode`. `settings.ip_configuration.require_ssl` was intended to be removed in 6.0.0, but is removed in 6.0.1 instead.

## Resource: `google_storage_bucket`

### `lifecycle_rule.condition.no_age` is now removed

Previously `lifecycle_rule.condition.age` attribute was being set to zero by default and `lifecycle_rule.condition.no_age` was introduced to prevent that.
Now `lifecycle_rule.condition.no_age` is no longer supported and `lifecycle_rule.condition.age` won't be set to zero by default.
Removed in favor of the field `lifecycle_rule.condition.send_age_if_zero` which can be used to set a zero value for the `lifecycle_rule.condition.age` attribute. 

For a seamless update, if your state today uses `no_age=true`, update it to remove `no_age` and set `send_age_if_zero=false`. If you do not use `no_age=true` and desire to continue creating rules with an `age=0` condition, you will need to add `send_age_if_zero=true` to your state to avoid any changes after updating to 6.0.0. 

With the 6.0.0 update, `send_age_if_zero` will be set to `false` by default unless declared explicitly `true`, and `age=0` conditions will be removed from existing buckets next time your `lifecycle_rule.condition` configuration is updated.

## Resource: `google_vpc_access_connector`

### Fields `min_throughput` and `max_throughput` no longer have default values

The fields `min_throughput` and `max_throughput` no longer have default values
set by the provider. This was necessary to add conflicting field validation, also
described in this guide.

No configuration changes are needed for existing resources as these fields' values
will default to values present in data returned from the API.

### Conflicting field validation added for `min_throughput` and `min_instances`, and `max_throughput` and `max_instances`

The provider will now enforce that `google_vpc_access_connector` resources can only
include one of `min_throughput` and `min_instances` and one of `max_throughput` and
`max_instances`. Previously if a user included all four fields in a resource block
they would experience a permadiff. This is a result of how `min_instances` and
`max_instances` fields' values take precedence in the API, and how the API calculates
values for `min_throughput` and `max_throughput` that match the number of instances.

Users will need to check their configuration for any `google_vpc_access_connector`
resource blocks that contain both fields in a conflicting pair, and remove one of those fields.
The fields that are removed from the configuration will still have Computed values,
that are derived from the API.

## Resource: `google_workstations_workstation_config`

### `host.gce_instance.disable_ssh` now defaults to true

`disable_ssh` field now defaults to true. To enable SSH, please set `disable_ssh` to false.