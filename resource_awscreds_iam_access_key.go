package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

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
