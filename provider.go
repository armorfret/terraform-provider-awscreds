package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-aws/aws"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: aws_provider().Schema,
		ResourcesMap: map[string]*schema.Resource{
			"awscreds_iam_access_key": resourceIamAccessKey(),
		},
		ConfigureFunc: configure,
	}
}

func configure(d *schema.ResourceData) (interface{}, error) {
	var wrapper Wrapper
	if err := wrapper.init(aws_provider(), d); err != nil {
		return nil, err
	}
	return wrapper, nil
}

func aws_provider() *schema.Provider {
	provider := aws.Provider()
	return resolve_provider(provider)
}
