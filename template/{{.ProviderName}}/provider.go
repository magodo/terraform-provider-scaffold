package {{.ProviderName}}

import (
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"{{.PkgPath}}/{{.ProviderName}}/internal/provider"
)

func Provider() terraform.ResourceProvider {
	return provider.Provider()
}
