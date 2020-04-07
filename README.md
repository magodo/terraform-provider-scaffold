# Terraform Provider Scaffold

This tool generates terraform provider scaffold, which aligns with hashicorp's [scaffolding layout](https://github.com/hashicorp/terraform-provider-scaffolding).

# Install

```go
GO111MODULE=off go get github.com/magodo/terraform-provider-scaffold
```

# Usage

```bash
> go run . -h
Create terraform provider scaffold

Usage:
  scaffold [flags]

Flags:
  -h, --help                   help for scaffold
  -o, --output_dir string      outputdir of project scaffold (default "scaffold")
  -p, --pkg_path string        go package path
  -n, --provider_name string   name of terraform provider
```

Example:

```bash
> terraform-provider-scaffold -p github.com/magodo/terraform-provider-foobar -n foobar

> tree
.
└── scaffold
    ├── CHANGELOG.md
    ├── foobar
    │   ├── internal
    │   │   └── provider
    │   │       └── provider.go
    │   └── provider.go
    ├── GNUmakefile
    ├── go.mod
    ├── LICENSE
    ├── main.go
    ├── scripts
    │   ├── changelog-links.sh
    │   ├── gofmtcheck.sh
    │   ├── terrafmt-acctests.sh
    │   └── terrafmt-website.sh
    └── website
        ├── allowed-subcategories
        └── docs
            ├── d
            └── r

9 directories, 12 files
```
