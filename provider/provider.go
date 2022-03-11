package provider

import (
	"github.com/armorfret/terraform-provider-aws/v4/winternal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

// Serve up the plugin
func Serve() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: shimProvider,
	})
}

func shimProvider() *schema.Provider {
	return &schema.Provider{
		Schema: awsProvider().Schema,
		ResourcesMap: map[string]*schema.Resource{
			"awscreds_iam_access_key": resourceIamAccessKey(),
		},
		ConfigureFunc: configure,
	}
}

func configure(d *schema.ResourceData) (interface{}, error) {
	var w wrapper
	if err := w.init(awsProvider(), d); err != nil {
		return nil, err
	}
	return w, nil
}

func awsProvider() *schema.Provider {
	return provider.Provider()
}

type wrapper struct {
	shimProvider *schema.Provider
	config       interface{}
}

func (w *wrapper) init(p *schema.Provider, d *schema.ResourceData) error {
	w.shimProvider = p
	config, err := p.ConfigureFunc(d)
	if err != nil {
		return err
	}
	w.config = config
	return nil
}

func (w wrapper) resource(name string) *schema.Resource {
	return w.shimProvider.ResourcesMap[name]
}
