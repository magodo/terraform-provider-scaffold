---
subcategory: "Scaffold"
layout: "{{.ProviderName}}"
page_title: "{{title .ProviderName}}: {{.ProviderName}}_scaffolding_resource"
description: |-
  Sample resource in the Terraform provider scaffolding.
---

# {{.ProviderName}}_scaffolding_resource

Sample resource in the Terraform provider scaffolding.

## Example Usage

```hcl
resource "{{.ProviderName}}_scaffolding_resource" "example" {
  sample_attribute = "foo"
}
```

## Argument Reference

The following arguments are supported:

* `sample_attribute` - Sample attribute.
