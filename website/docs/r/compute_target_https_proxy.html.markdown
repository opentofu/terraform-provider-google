---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/TargetHttpsProxy.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Compute Engine"
description: |-
  Represents a TargetHttpsProxy resource, which is used by one or more
  global forwarding rule to route incoming HTTPS requests to a URL map.
---

# google_compute_target_https_proxy

Represents a TargetHttpsProxy resource, which is used by one or more
global forwarding rule to route incoming HTTPS requests to a URL map.


To get more information about TargetHttpsProxy, see:

* [API documentation](https://cloud.google.com/compute/docs/reference/v1/targetHttpsProxies)
* How-to Guides
    * [Official Documentation](https://cloud.google.com/compute/docs/load-balancing/http/target-proxies)

<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=target_https_proxy_basic&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Target Https Proxy Basic


```hcl
resource "google_compute_target_https_proxy" "default" {
  name             = "test-proxy"
  url_map          = google_compute_url_map.default.id
  ssl_certificates = [google_compute_ssl_certificate.default.id]
}

resource "google_compute_ssl_certificate" "default" {
  name        = "my-certificate"
  private_key = file("path/to/private.key")
  certificate = file("path/to/certificate.crt")
}

resource "google_compute_url_map" "default" {
  name        = "url-map"
  description = "a description"

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
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=target_https_proxy_http_keep_alive_timeout&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Target Https Proxy Http Keep Alive Timeout


```hcl
resource "google_compute_target_https_proxy" "default" {
  name                        = "test-http-keep-alive-timeout-proxy"
  http_keep_alive_timeout_sec = 610
  url_map                     = google_compute_url_map.default.id
  ssl_certificates            = [google_compute_ssl_certificate.default.id]
}

resource "google_compute_ssl_certificate" "default" {
  name        = "my-certificate"
  private_key = file("path/to/private.key")
  certificate = file("path/to/certificate.crt")
}

resource "google_compute_url_map" "default" {
  name        = "url-map"
  description = "a description"

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
  name                  = "backend-service"
  port_name             = "http"
  protocol              = "HTTP"
  timeout_sec           = 10
  load_balancing_scheme = "EXTERNAL_MANAGED"

  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  name               = "http-health-check"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=target_https_proxy_mtls&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Target Https Proxy Mtls


```hcl
data "google_project" "project" {
  provider = google-beta
}

resource "google_compute_target_https_proxy" "default" {
  provider          = google-beta
  name              = "test-mtls-proxy"
  url_map           = google_compute_url_map.default.id
  ssl_certificates  = [google_compute_ssl_certificate.default.id]
  server_tls_policy = google_network_security_server_tls_policy.default.id
}

resource "google_certificate_manager_trust_config" "default" {
  provider    = google-beta
  name        = "my-trust-config"
  description = "sample description for the trust config"
  location    = "global"

  trust_stores {
    trust_anchors {
      pem_certificate = file("test-fixtures/ca_cert.pem")
    }
    intermediate_cas {
      pem_certificate = file("test-fixtures/ca_cert.pem")
    }
  }

  labels = {
    foo = "bar"
  }
}

resource "google_network_security_server_tls_policy" "default" {
  provider               = google-beta
  name                   = "my-tls-policy"
  description            = "my description"
  location               = "global"
  allow_open             = "false"
  mtls_policy {
    client_validation_mode = "ALLOW_INVALID_OR_MISSING_CLIENT_CERT"
    client_validation_trust_config = "projects/${data.google_project.project.number}/locations/global/trustConfigs/${google_certificate_manager_trust_config.default.name}"
  }
}

resource "google_compute_ssl_certificate" "default" {
  provider    = google-beta
  name        = "my-certificate"
  private_key = file("path/to/private.key")
  certificate = file("path/to/certificate.crt")
}

resource "google_compute_url_map" "default" {
  provider    = google-beta
  name        = "url-map"
  description = "a description"

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
  provider    = google-beta
  name        = "backend-service"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  provider           = google-beta
  name               = "http-health-check"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=target_https_proxy_certificate_manager_certificate&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Target Https Proxy Certificate Manager Certificate


```hcl

resource "google_compute_target_https_proxy" "default" {
  name                             = "target-http-proxy"
  url_map                          = google_compute_url_map.default.id
  certificate_manager_certificates =  ["//certificatemanager.googleapis.com/${google_certificate_manager_certificate.default.id}"] # [google_certificate_manager_certificate.default.id] is also acceptable
}

resource "google_certificate_manager_certificate" "default" {
  name              = "my-certificate"
  scope             = "ALL_REGIONS"
  self_managed {
    pem_certificate = file("test-fixtures/cert.pem")
    pem_private_key = file("test-fixtures/private-key.pem")                                                                                                                
  }
}

resource "google_compute_url_map" "default" {
  name        = "url-map"
  description = "a description"

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
  name        = "backend-service"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = 10
  load_balancing_scheme = "INTERNAL_MANAGED"
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=target_https_proxy_fingerprint&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Target Https Proxy Fingerprint


```hcl
resource "google_compute_target_https_proxy" "default" {
  name             = "test-fingerprint-proxy"
  url_map          = google_compute_url_map.default.id
  ssl_certificates = [google_compute_ssl_certificate.default.id]
}

resource "google_compute_ssl_certificate" "default" {
  name        = "my-certificate"
  private_key = file("path/to/private.key")
  certificate = file("path/to/certificate.crt")
}

resource "google_compute_url_map" "default" {
  name        = "url-map"
  description = "a description"

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

output "target_https_proxy_fingerprint" {
  value       = google_compute_target_https_proxy.default.fingerprint
  description = "The fingerprint of the target HTTPS proxy for optimistic locking"
}
```

## Argument Reference

The following arguments are supported:


* `name` -
  (Required)
  Name of the resource. Provided by the client when the resource is
  created. The name must be 1-63 characters long, and comply with
  RFC1035. Specifically, the name must be 1-63 characters long and match
  the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the
  first character must be a lowercase letter, and all following
  characters must be a dash, lowercase letter, or digit, except the last
  character, which cannot be a dash.

* `url_map` -
  (Required)
  A reference to the UrlMap resource that defines the mapping from URL
  to the BackendService.


* `description` -
  (Optional)
  An optional description of this resource.

* `quic_override` -
  (Optional)
  Specifies the QUIC override policy for this resource. This determines
  whether the load balancer will attempt to negotiate QUIC with clients
  or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is
  specified, Google manages whether QUIC is used.
  Default value is `NONE`.
  Possible values are: `NONE`, `ENABLE`, `DISABLE`.

* `tls_early_data` -
  (Optional)
  Specifies whether TLS 1.3 0-RTT Data (“Early Data”) should be accepted for this service.
  Early Data allows a TLS resumption handshake to include the initial application payload
  (a HTTP request) alongside the handshake, reducing the effective round trips to “zero”.
  This applies to TLS 1.3 connections over TCP (HTTP/2) as well as over UDP (QUIC/h3).
  Possible values are: `STRICT`, `PERMISSIVE`, `UNRESTRICTED`, `DISABLED`.

* `certificate_manager_certificates` -
  (Optional)
  URLs to certificate manager certificate resources that are used to authenticate connections between users and the load balancer.
  Certificate manager certificates only apply when the load balancing scheme is set to INTERNAL_MANAGED.
  For EXTERNAL and EXTERNAL_MANAGED, use certificate_map instead.
  sslCertificates and certificateManagerCertificates fields can not be defined together.
  Accepted format is `//certificatemanager.googleapis.com/projects/{project}/locations/{location}/certificates/{resourceName}` or just the self_link `projects/{project}/locations/{location}/certificates/{resourceName}`

* `ssl_certificates` -
  (Optional)
  URLs to SslCertificate resources that are used to authenticate connections between users and the load balancer.
  Currently, you may specify up to 15 SSL certificates. sslCertificates do not apply when the load balancing scheme is set to INTERNAL_SELF_MANAGED.
  sslCertificates and certificateManagerCertificates can not be defined together.

* `certificate_map` -
  (Optional)
  A reference to the CertificateMap resource uri that identifies a certificate map
  associated with the given target proxy. This field is only supported for EXTERNAL and EXTERNAL_MANAGED load balancing schemes.
  For INTERNAL_MANAGED, use certificate_manager_certificates instead.
  Accepted format is `//certificatemanager.googleapis.com/projects/{project}/locations/{location}/certificateMaps/{resourceName}`.

* `ssl_policy` -
  (Optional)
  A reference to the SslPolicy resource that will be associated with
  the TargetHttpsProxy resource. If not set, the TargetHttpsProxy
  resource will not have any SSL policy configured.

* `proxy_bind` -
  (Optional)
  This field only applies when the forwarding rule that references
  this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.

* `http_keep_alive_timeout_sec` -
  (Optional)
  Specifies how long to keep a connection open, after completing a response,
  while there is no matching traffic (in seconds). If an HTTP keepalive is
  not specified, a default value will be used. For Global
  external HTTP(S) load balancer, the default value is 610 seconds, the
  minimum allowed value is 5 seconds and the maximum allowed value is 1200
  seconds. For cross-region internal HTTP(S) load balancer, the default
  value is 600 seconds, the minimum allowed value is 5 seconds, and the
  maximum allowed value is 600 seconds. For Global external HTTP(S) load
  balancer (classic), this option is not available publicly.

* `server_tls_policy` -
  (Optional)
  A URL referring to a networksecurity.ServerTlsPolicy
  resource that describes how the proxy should authenticate inbound
  traffic. serverTlsPolicy only applies to a global TargetHttpsProxy
  attached to globalForwardingRules with the loadBalancingScheme
  set to INTERNAL_SELF_MANAGED or EXTERNAL or EXTERNAL_MANAGED.
  For details which ServerTlsPolicy resources are accepted with
  INTERNAL_SELF_MANAGED and which with EXTERNAL, EXTERNAL_MANAGED
  loadBalancingScheme consult ServerTlsPolicy documentation.
  If left blank, communications are not encrypted.
  If you remove this field from your configuration at the same time as
  deleting or recreating a referenced ServerTlsPolicy resource, you will
  receive a resourceInUseByAnotherResource error. Use lifecycle.create_before_destroy
  within the ServerTlsPolicy resource to avoid this.

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/global/targetHttpsProxies/{{name}}`

* `creation_timestamp` -
  Creation timestamp in RFC3339 text format.

* `proxy_id` -
  The unique identifier for the resource.

* `fingerprint` -
  Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
  This field will be ignored when inserting a TargetHttpsProxy. An up-to-date fingerprint must be provided in order to
  patch the TargetHttpsProxy; otherwise, the request will fail with error 412 conditionNotMet.
  To see the latest fingerprint, make a get() request to retrieve the TargetHttpsProxy.
  A base64-encoded string.
* `self_link` - The URI of the created resource.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


TargetHttpsProxy can be imported using any of these accepted formats:

* `projects/{{project}}/global/targetHttpsProxies/{{name}}`
* `{{project}}/{{name}}`
* `{{name}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import TargetHttpsProxy using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/global/targetHttpsProxies/{{name}}"
  to = google_compute_target_https_proxy.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), TargetHttpsProxy can be imported using one of the formats above. For example:

```
$ terraform import google_compute_target_https_proxy.default projects/{{project}}/global/targetHttpsProxies/{{name}}
$ terraform import google_compute_target_https_proxy.default {{project}}/{{name}}
$ terraform import google_compute_target_https_proxy.default {{name}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
