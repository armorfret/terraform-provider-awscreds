package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-aws/aws"
)

// Serve up the plugin
func Serve() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider,
	})
}

func provider() terraform.ResourceProvider {
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
	provider := aws.Provider()
	return resolveProvider(provider)
}

type wrapper struct {
	provider *schema.Provider
	config   interface{}
}

func (w *wrapper) init(p *schema.Provider, d *schema.ResourceData) error {
	w.provider = p
	config, err := p.ConfigureFunc(d)
	if err != nil {
		return err
	}
	w.config = config
	return nil
}

func (w wrapper) resource(name string) *schema.Resource {
	return w.provider.ResourcesMap[name]
}

func resolveProvider(provider terraform.ResourceProvider) *schema.Provider {
	return provider.(*schema.Provider)
}
