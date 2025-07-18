---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/ManagedSslCertificate.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Compute Engine"
description: |-
  An SslCertificate resource, used for HTTPS load balancing.
---

# google_compute_managed_ssl_certificate

An SslCertificate resource, used for HTTPS load balancing.  This resource
represents a certificate for which the certificate secrets are created and
managed by Google.

For a resource where you provide the key, see the
SSL Certificate resource.


To get more information about ManagedSslCertificate, see:

* [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/sslCertificates)
* How-to Guides
    * [Official Documentation](https://cloud.google.com/load-balancing/docs/ssl-certificates)

~> **Warning:** This resource should be used with extreme caution!  Provisioning an SSL
certificate is complex.  Ensure that you understand the lifecycle of a
certificate before attempting complex tasks like cert rotation automatically.
This resource will "return" as soon as the certificate object is created,
but post-creation the certificate object will go through a "provisioning"
process.  The provisioning process can complete only when the domain name
for which the certificate is created points to a target pool which, itself,
points at the certificate.  Depending on your DNS provider, this may take
some time, and migrating from self-managed certificates to Google-managed
certificates may entail some downtime while the certificate provisions.

In conclusion: Be extremely cautious.

<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=managed_ssl_certificate_basic&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Managed Ssl Certificate Basic


```hcl
resource "google_compute_managed_ssl_certificate" "default" {
  name = "test-cert"

  managed {
    domains = ["sslcert.tf-test.club."]
  }
}

resource "google_compute_target_https_proxy" "default" {
  name             = "test-proxy"
  url_map          = google_compute_url_map.default.id
  ssl_certificates = [google_compute_managed_ssl_certificate.default.id]
}

resource "google_compute_url_map" "default" {
  name        = "url-map"
  description = "a description"

  default_service = google_compute_backend_service.default.id

  host_rule {
    hosts        = ["sslcert.tf-test.club"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name            = "allpaths"
    default_service = google_compute_backend_service.default.id

    path_rule {
      paths   = ["/*"]
      service = google_compute_backend_service.default.id
    }
  }
}

resource "google_compute_backend_service" "default" {
  name        = "backend-service"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  name               = "http-health-check"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_compute_global_forwarding_rule" "default" {
  name       = "forwarding-rule"
  target     = google_compute_target_https_proxy.default.id
  port_range = 443
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=managed_ssl_certificate_recreation&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Managed Ssl Certificate Recreation


```hcl
// This example allows the list of managed domains to be modified and will
// recreate the ssl certificate and update the target https proxy correctly

resource "google_compute_target_https_proxy" "default" {
  name             = "test-proxy"
  url_map          = google_compute_url_map.default.id
  ssl_certificates = [google_compute_managed_ssl_certificate.cert.id]
}

locals {
  managed_domains = tolist(["test.example.com"])
}

resource "random_id" "certificate" {
  byte_length = 4
  prefix      = "issue6147-cert-"

  keepers = {
    domains = join(",", local.managed_domains)
  }
}

resource "google_compute_managed_ssl_certificate" "cert" {
  name     = random_id.certificate.hex

  lifecycle {
    create_before_destroy = true
  }

  managed {
    domains = local.managed_domains
  }
}

resource "google_compute_url_map" "default" {
  name            = "url-map"
  description     = "a description"
  default_service = google_compute_backend_service.default.id
  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }
  path_matcher {
    name            = "allpaths"
    default_service = google_compute_backend_service.default.id
    path_rule {
      paths   = ["/*"]
      service = google_compute_backend_service.default.id
    }
  }
}

resource "google_compute_backend_service" "default" {
  name          = "backend-service"
  port_name     = "http"
  protocol      = "HTTP"
  timeout_sec   = 10
  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  name               = "http-health-check"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
```

## Argument Reference

The following arguments are supported:



* `description` -
  (Optional)
  An optional description of this resource.

* `name` -
  (Optional)
  Name of the resource. Provided by the client when the resource is
  created. The name must be 1-63 characters long, and comply with
  RFC1035. Specifically, the name must be 1-63 characters long and match
  the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the
  first character must be a lowercase letter, and all following
  characters must be a dash, lowercase letter, or digit, except the last
  character, which cannot be a dash.
  These are in the same namespace as the managed SSL certificates.

* `managed` -
  (Optional)
  Properties relevant to a managed certificate.  These will be used if the
  certificate is managed (as indicated by a value of `MANAGED` in `type`).
  Structure is [documented below](#nested_managed).

* `type` -
  (Optional)
  Enum field whose value is always `MANAGED` - used to signal to the API
  which type this is.
  Default value is `MANAGED`.
  Possible values are: `MANAGED`.

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



<a name="nested_managed"></a>The `managed` block supports:

* `domains` -
  (Required)
  Domains for which a managed SSL certificate will be valid.  Currently,
  there can be up to 100 domains in this list.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/global/sslCertificates/{{name}}`

* `creation_timestamp` -
  Creation timestamp in RFC3339 text format.

* `certificate_id` -
  The unique identifier for the resource.

* `subject_alternative_names` -
  Domains associated with the certificate via Subject Alternative Name.

* `expire_time` -
  Expire time of the certificate in RFC3339 text format.
* `self_link` - The URI of the created resource.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 30 minutes.
- `delete` - Default is 30 minutes.

## Import


ManagedSslCertificate can be imported using any of these accepted formats:

* `projects/{{project}}/global/sslCertificates/{{name}}`
* `{{project}}/{{name}}`
* `{{name}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import ManagedSslCertificate using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/global/sslCertificates/{{name}}"
  to = google_compute_managed_ssl_certificate.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), ManagedSslCertificate can be imported using one of the formats above. For example:

```
$ terraform import google_compute_managed_ssl_certificate.default projects/{{project}}/global/sslCertificates/{{name}}
$ terraform import google_compute_managed_ssl_certificate.default {{project}}/{{name}}
$ terraform import google_compute_managed_ssl_certificate.default {{name}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
