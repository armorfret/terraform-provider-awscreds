package main

import (
	"github.com/hashicorp/terraform/builtin/providers/aws"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},
		ResourcesMap: map[string]*schema.Resource{
			"awscreds_iam_access_key": resourceIamAccessKey(),
		},
		ConfigureFunc: configure,
	}
}

type Wrapper struct {
	provider *schema.Provider
	config   interface{}
}

func (w *Wrapper) init(p *schema.Provider, d *schema.ResourceData) error {
	w.provider = p
	config, err := p.ConfigureFunc(d)
	if err != nil {
		return err
	}
	w.config = config
	return nil
}

func (w *Wrapper) resource(name string) *schema.Resource {
	return w.provider.ResourcesMap[name]
}

func configure(d *schema.ResourceData) (interface{}, error) {
	var wrapper Wrapper
	provider := aws.Provider()
	cast := provider.(*schema.Provider)
	if err := wrapper.init(cast, d); err != nil {
		return nil, err
	}
	return wrapper, nil
}
