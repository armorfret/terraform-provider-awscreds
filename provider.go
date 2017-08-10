package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-aws/aws"
)

// Provider exposes the custom Terraform provider
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: awsProvider().Schema,
		ResourcesMap: map[string]*schema.Resource{
			"awscreds_iam_access_key": resourceIamAccessKey(),
		},
		ConfigureFunc: configure,
	}
}

func configure(d *schema.ResourceData) (interface{}, error) {
	var wrapper Wrapper
	if err := wrapper.init(awsProvider(), d); err != nil {
		return nil, err
	}
	return wrapper, nil
}

func awsProvider() *schema.Provider {
	provider := aws.Provider()
	return resolveProvider(provider)
}
