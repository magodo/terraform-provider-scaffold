package pkg

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			// TODO
		},
		ConfigureFunc: func(d *schema.ResourceData) (interface{}, error) {
			return &client{
				// TODO
			}, nil
		},
	}
}

type client struct {
	// TODO
}
