package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

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

func (w Wrapper) resource(name string) *schema.Resource {
	return w.provider.ResourcesMap[name]
}

func resolve_provider(provider terraform.ResourceProvider) *schema.Provider {
	return provider.(*schema.Provider)
}
