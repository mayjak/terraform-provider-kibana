package kibana

import (
	"os"
	"github.com/ewilde/go-kibana"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"kibana_uri": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: envDefaultFuncWithDefault(kibana.EnvKibanaUri, kibana.DefaultKibanaUri),
				Description: "The address of the kibana admin url e.g. " + kibana.DefaultKibanaUri,
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"kibana_search": resourceDir(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func envDefaultFuncWithDefault(key string, defaultValue string) schema.SchemaDefaultFunc {
	return func() (interface{}, error) {
		if v := os.Getenv(key); v != "" {
			if v == "true" {
				return true, nil
			} else if v == "false" {
				return false, nil
			}
			return v, nil
		}
		return defaultValue, nil
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := &kibana.Config{
		HostAddress: d.Get("kibana_uri").(string),
	}

	return kibana.NewClient(config), nil
}