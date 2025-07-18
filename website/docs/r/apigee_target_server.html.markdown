---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/apigee/TargetServer.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Apigee"
description: |-
  TargetServer configuration.
---

# google_apigee_target_server

TargetServer configuration. TargetServers are used to decouple a proxy TargetEndpoint HTTPTargetConnections from concrete URLs for backend services.


To get more information about TargetServer, see:

* [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.environments.targetservers/create)
* How-to Guides
    * [Load balancing across backend servers](https://cloud.google.com/apigee/docs/api-platform/deploy/load-balancing-across-backend-servers)

## Example Usage - Apigee Target Server Test Basic


```hcl
resource "google_project" "project" {
  project_id      = "my-project"
  name            = "my-project"
  org_id          = "123456789"
  billing_account = "000000-0000000-0000000-000000"
  deletion_policy = "DELETE"
}

resource "google_project_service" "apigee" {
  project = google_project.project.project_id
  service = "apigee.googleapis.com"
}

resource "google_project_service" "servicenetworking" {
  project    = google_project.project.project_id
  service    = "servicenetworking.googleapis.com"
  depends_on = [google_project_service.apigee]
}

resource "google_project_service" "compute" {
  project    = google_project.project.project_id
  service    = "compute.googleapis.com"
  depends_on = [google_project_service.servicenetworking]
}

resource "google_compute_network" "apigee_network" {
  name       = "apigee-network"
  project    = google_project.project.project_id
  depends_on = [google_project_service.compute]
}

resource "google_compute_global_address" "apigee_range" {
  name          = "apigee-range"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = google_compute_network.apigee_network.id
  project       = google_project.project.project_id
}

resource "google_service_networking_connection" "apigee_vpc_connection" {
  network                 = google_compute_network.apigee_network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.apigee_range.name]
  depends_on              = [google_project_service.servicenetworking]
}

resource "google_apigee_organization" "apigee_org" {
  analytics_region   = "us-central1"
  project_id         = google_project.project.project_id
  authorized_network = google_compute_network.apigee_network.id
  depends_on         = [
    google_service_networking_connection.apigee_vpc_connection,
    google_project_service.apigee,
  ]
}

resource "google_apigee_environment" "apigee_environment" {
  org_id       = google_apigee_organization.apigee_org.id
  name         = "my-environment-name"
  description  = "Apigee Environment"
  display_name = "environment-1"
}

resource "google_apigee_target_server" "apigee_target_server" {
  name        = "my-target-server"
  description = "Apigee Target Server"
  protocol    = "HTTP"
  host        = "abc.foo.com"
  port        = 8080
  env_id      = google_apigee_environment.apigee_environment.id
}
```

## Argument Reference

The following arguments are supported:


* `name` -
  (Required)
  The resource id of this reference. Values must match the regular expression [\w\s-.]+.

* `host` -
  (Required)
  The host name this target connects to. Value must be a valid hostname as described by RFC-1123.

* `port` -
  (Required)
  The port number this target connects to on the given host. Value must be between 1 and 65535, inclusive.

* `env_id` -
  (Required)
  The Apigee environment group associated with the Apigee environment,
  in the format `organizations/{{org_name}}/environments/{{env_name}}`.


* `description` -
  (Optional)
  A human-readable description of this TargetServer.

* `is_enabled` -
  (Optional)
  Enabling/disabling a TargetServer is useful when TargetServers are used in load balancing configurations, and one or more TargetServers need to taken out of rotation periodically. Defaults to true.

* `s_sl_info` -
  (Optional)
  Specifies TLS configuration info for this TargetServer. The JSON name is sSLInfo for legacy/backwards compatibility reasons -- Edge originally supported SSL, and the name is still used for TLS configuration.
  Structure is [documented below](#nested_s_sl_info).

* `protocol` -
  (Optional)
  Immutable. The protocol used by this TargetServer.
  Possible values are: `HTTP`, `HTTP2`, `GRPC_TARGET`, `GRPC`, `EXTERNAL_CALLOUT`.



<a name="nested_s_sl_info"></a>The `s_sl_info` block supports:

* `enabled` -
  (Required)
  Enables TLS. If false, neither one-way nor two-way TLS will be enabled.

* `client_auth_enabled` -
  (Optional)
  Enables two-way TLS.

* `key_store` -
  (Optional)
  Required if clientAuthEnabled is true. The resource ID of the keystore.

* `key_alias` -
  (Optional)
  Required if clientAuthEnabled is true. The resource ID for the alias containing the private key and cert.

* `trust_store` -
  (Optional)
  The resource ID of the truststore.

* `ignore_validation_errors` -
  (Optional)
  If true, Edge ignores TLS certificate errors. Valid when configuring TLS for target servers and target endpoints, and when configuring virtual hosts that use 2-way TLS. When used with a target endpoint/target server, if the backend system uses SNI and returns a cert with a subject Distinguished Name (DN) that does not match the hostname, there is no way to ignore the error and the connection fails.

* `protocols` -
  (Optional)
  The TLS versioins to be used.

* `ciphers` -
  (Optional)
  The SSL/TLS cipher suites to be used. For programmable proxies, it must be one of the cipher suite names listed in: http://docs.oracle.com/javase/8/docs/technotes/guides/security/StandardNames.html#ciphersuites. For configurable proxies, it must follow the configuration specified in: https://commondatastorage.googleapis.com/chromium-boringssl-docs/ssl.h.html#Cipher-suite-configuration. This setting has no effect for configurable proxies when negotiating TLS 1.3.

* `common_name` -
  (Optional)
  The TLS Common Name of the certificate.
  Structure is [documented below](#nested_s_sl_info_common_name).

* `enforce` -
  (Optional)
  If true, TLS is strictly enforced.


<a name="nested_s_sl_info_common_name"></a>The `common_name` block supports:

* `value` -
  (Optional)
  The TLS Common Name string of the certificate.

* `wildcard_match` -
  (Optional)
  Indicates whether the cert should be matched against as a wildcard cert.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `{{env_id}}/targetservers/{{name}}`


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 1 minutes.
- `update` - Default is 1 minutes.
- `delete` - Default is 1 minutes.

## Import


TargetServer can be imported using any of these accepted formats:

* `{{env_id}}/targetservers/{{name}}`
* `{{env_id}}/{{name}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import TargetServer using one of the formats above. For example:

```tf
import {
  id = "{{env_id}}/targetservers/{{name}}"
  to = google_apigee_target_server.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), TargetServer can be imported using one of the formats above. For example:

```
$ terraform import google_apigee_target_server.default {{env_id}}/targetservers/{{name}}
$ terraform import google_apigee_target_server.default {{env_id}}/{{name}}
```
