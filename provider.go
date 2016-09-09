package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},
		ResourcesMap: map[string]*schema.Resource{
			"iam_access_key": resourceIamAccessKey(),
		},
		ConfigureFunc: configure,
	}
}

func configure(d *schema.ResourceData) (interface{}, error) {
	return nil, nil
}

func resourceIamAccessKey() *schema.Resource {
	return &schema.Resource{
		Create: createIamAccessKey,
		Read:   readIamAccessKey,
		Update: updateIamAccessKey,
		Delete: deleteIamAccessKey,
		Schema: map[string]*schema.Schema{
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"file": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func createIamAccessKey(d *schema.ResourceData, m interface{}) error {
	return nil
}

func readIamAccessKey(d *schema.ResourceData, m interface{}) error {
	return nil
}

func updateIamAccessKey(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteIamAccessKey(d *schema.ResourceData, m interface{}) error {
	return nil
}
