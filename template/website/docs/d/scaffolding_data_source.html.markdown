---
subcategory: "Scaffold"
layout: "{{.ProviderName}}"
page_title: "{{title .ProviderName}}: {{.ProviderName}}_scaffolding_data_source"
description: |-
  Sample data source in the Terraform provider scaffolding.
---

# {{.ProviderName}}_scaffolding_data_source

Sample data source in the Terraform provider scaffolding.

## Example Usage

```hcl
data "{{.ProviderName}}_scaffolding_data_source" "example" {
  sample_attribute = "foo"
}
```

## Attributes Reference

* `sample_attribute` - Sample attribute.
