---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/website/docs/d/container_registry_repository.html.markdown
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Container Registry"
description: |-
  Get URLs for a given project's container registry repository.
---

# google_container_registry_repository

-> **Warning**: Container Registry is deprecated. Effective March 18, 2025, Container Registry is shut down and writing images to Container Registry is unavailable. Resource will be removed in future major release.

This data source fetches the project name, and provides the appropriate URLs to use for container registry for this project.

The URLs are computed entirely offline - as long as the project exists, they will be valid, but this data source does not contact Google Container Registry (GCR) at any point.

## Example Usage

```hcl
data "google_container_registry_repository" "foo" {
}

output "gcr_location" {
  value = data.google_container_registry_repository.foo.repository_url
}
```

## Argument Reference
* `project`: (Optional) The project ID that this repository is attached to.  If not provided, provider project will be used instead.
* `region`: (Optional) The GCR region to use.  As of this writing, one of `asia`, `eu`, and `us`.  See [the documentation](https://cloud.google.com/container-registry/docs/pushing-and-pulling) for additional information.

## Attributes Reference
In addition to the arguments listed above, this data source exports:

* `repository_url`: The URL at which the repository can be accessed.
