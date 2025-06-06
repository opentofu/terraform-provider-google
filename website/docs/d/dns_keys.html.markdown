---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/website/docs/d/dns_keys.html.markdown
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Cloud DNS"
description: |-
  Get DNSKEY and DS records of DNSSEC-signed managed zones.
---

# google_dns_keys

Get the DNSKEY and DS records of DNSSEC-signed managed zones.

For more information see the
[official documentation](https://cloud.google.com/dns/docs/dnskeys/)
and [API](https://cloud.google.com/dns/docs/reference/v1/dnsKeys).

~> A google_dns_managed_zone resource must have DNSSEC enabled in order
to contain any DNSKEYs. Queries to managed zones without this setting
enabled will result in a 404 error as the collection of DNSKEYs does
not exist in the DNS API.

## Example Usage

```hcl
resource "google_dns_managed_zone" "foo" {
  name     = "foobar"
  dns_name = "foo.bar."

  dnssec_config {
    state         = "on"
    non_existence = "nsec3"
  }
}

data "google_dns_keys" "foo_dns_keys" {
  managed_zone = google_dns_managed_zone.foo.id
}

output "foo_dns_ds_record" {
  description = "DS record of the foo subdomain."
  value       = data.google_dns_keys.foo_dns_keys.key_signing_keys[0].ds_record
}
```

## Argument Reference

The following arguments are supported:

* `managed_zone` - (Required) The name or id of the Cloud DNS managed zone.

* `project` - (Optional) The ID of the project in which the resource belongs. If `project` is not provided, the provider project is used.

## Attributes Reference

The following attributes are exported:

* `key_signing_keys` - A list of Key-signing key (KSK) records. Structure is [documented below](#nested_key_signing_keys). Additionally, the DS record is provided:

  * `ds_record` - The DS record based on the KSK record. This is used when [delegating](https://cloud.google.com/dns/docs/dnssec-advanced#subdelegation) DNSSEC-signed subdomains.

* `zone_signing_keys` - A list of Zone-signing key (ZSK) records. Structure is documented below.

---

<a name="nested_key_signing_keys"></a>The `key_signing_keys` and `zone_signing_keys` block supports:

  * `algorithm` - String mnemonic specifying the DNSSEC algorithm of this key. Immutable after creation time. Possible values are `ecdsap256sha256`, `ecdsap384sha384`, `rsasha1`, `rsasha256`, and `rsasha512`.

  * `creation_time` - The time that this resource was created in the control plane. This is in RFC3339 text format.

  * `description` - A mutable string of at most 1024 characters associated with this resource for the user's convenience.

  * `digests` - A list of cryptographic hashes of the DNSKEY resource record associated with this DnsKey. These digests are needed to construct a DS record that points at this DNS key. Each contains:
    - `digest` - The base-16 encoded bytes of this digest. Suitable for use in a DS resource record.
    - `type` - Specifies the algorithm used to calculate this digest. Possible values are `sha1`, `sha256` and `sha384`

  * `id` - Unique identifier for the resource; defined by the server.

  * `is_active` - Active keys will be used to sign subsequent changes to the ManagedZone. Inactive keys will still be present as DNSKEY Resource Records for the use of resolvers validating existing signatures.

  * `key_length` - Length of the key in bits. Specified at creation time then immutable.

  * `key_tag` - The key tag is a non-cryptographic hash of the a DNSKEY resource record associated with this DnsKey. The key tag can be used to identify a DNSKEY more quickly (but it is not a unique identifier). In particular, the key tag is used in a parent zone's DS record to point at the DNSKEY in this child ManagedZone. The key tag is a number in the range [0, 65535] and the algorithm to calculate it is specified in RFC4034 Appendix B.

  * `public_key` - Base64 encoded public half of this key.
