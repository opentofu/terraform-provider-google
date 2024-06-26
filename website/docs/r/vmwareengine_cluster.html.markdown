---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This file is automatically generated by Magic Modules and manual
#     changes will be clobbered when the file is regenerated.
#
#     Please read more about how to change this file in
#     .github/CONTRIBUTING.md.
#
# ----------------------------------------------------------------------------
subcategory: "Cloud VMware Engine"
description: |-
  A cluster in a private cloud.
---

# google_vmwareengine_cluster

A cluster in a private cloud.


To get more information about Cluster, see:

* [API documentation](https://cloud.google.com/vmware-engine/docs/reference/rest/v1/projects.locations.privateClouds.clusters)

## Example Usage - Vmware Engine Cluster Basic


```hcl
resource "google_vmwareengine_cluster" "vmw-engine-ext-cluster" {
  name     = "ext-cluster"
  parent   = google_vmwareengine_private_cloud.cluster-pc.id
  node_type_configs {
    node_type_id = "standard-72"
    node_count   = 3
  }
}

resource "google_vmwareengine_private_cloud" "cluster-pc" {
  location    = "us-west1-a"
  name        = "sample-pc"
  description = "Sample test PC."
  network_config {
    management_cidr       = "192.168.30.0/24"
    vmware_engine_network = google_vmwareengine_network.cluster-nw.id
  }

  management_cluster {
    cluster_id = "sample-mgmt-cluster"
    node_type_configs {
      node_type_id = "standard-72"
      node_count   = 3
    }
  }
}

resource "google_vmwareengine_network" "cluster-nw" {
  name        = "pc-nw"
  type        = "STANDARD"
  location    = "global"
  description = "PC network description."
}
```
## Example Usage - Vmware Engine Cluster Full


```hcl
resource "google_vmwareengine_cluster" "vmw-ext-cluster" {
  name     = "ext-cluster"
  parent   = google_vmwareengine_private_cloud.cluster-pc.id
  node_type_configs {
    node_type_id = "standard-72"
    node_count   = 3
    custom_core_count = 32
  }
}

resource "google_vmwareengine_private_cloud" "cluster-pc" {
  location    = "us-west1-a"
  name        = "sample-pc"
  description = "Sample test PC."
  network_config {
    management_cidr       = "192.168.30.0/24"
    vmware_engine_network = google_vmwareengine_network.cluster-nw.id
  }

  management_cluster {
    cluster_id = "sample-mgmt-cluster"
    node_type_configs {
      node_type_id = "standard-72"
      node_count   = 3
      custom_core_count = 32
    }
  }
}

resource "google_vmwareengine_network" "cluster-nw" {
  name        = "pc-nw"
  type        = "STANDARD"
  location    = "global"
  description = "PC network description."
}
```

## Argument Reference

The following arguments are supported:


* `parent` -
  (Required)
  The resource name of the private cloud to create a new cluster in.
  Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.
  For example: projects/my-project/locations/us-west1-a/privateClouds/my-cloud

* `name` -
  (Required)
  The ID of the Cluster.


- - -


* `node_type_configs` -
  (Optional)
  The map of cluster node types in this cluster,
  where the key is canonical identifier of the node type (corresponds to the NodeType).
  Structure is [documented below](#nested_node_type_configs).


<a name="nested_node_type_configs"></a>The `node_type_configs` block supports:

* `node_type_id` - (Required) The identifier for this object. Format specified above.

* `node_count` -
  (Required)
  The number of nodes of this type in the cluster.

* `custom_core_count` -
  (Optional)
  Customized number of cores available to each node of the type.
  This number must always be one of `nodeType.availableCustomCoreCounts`.
  If zero is provided max value from `nodeType.availableCustomCoreCounts` will be used.
  Once the customer is created then corecount cannot be changed.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `{{parent}}/clusters/{{name}}`

* `management` -
  True if the cluster is a management cluster; false otherwise.
  There can only be one management cluster in a private cloud and it has to be the first one.

* `uid` -
  System-generated unique identifier for the resource.

* `state` -
  State of the Cluster.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 210 minutes.
- `update` - Default is 190 minutes.
- `delete` - Default is 150 minutes.

## Import


Cluster can be imported using any of these accepted formats:

* `{{parent}}/clusters/{{name}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Cluster using one of the formats above. For example:

```tf
import {
  id = "{{parent}}/clusters/{{name}}"
  to = google_vmwareengine_cluster.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), Cluster can be imported using one of the formats above. For example:

```
$ terraform import google_vmwareengine_cluster.default {{parent}}/clusters/{{name}}
```
