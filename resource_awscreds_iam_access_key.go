package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"io/ioutil"
)

func resourceIamAccessKey() *schema.Resource {
	return &schema.Resource{
		Create: createIamAccessKey,
		Read:   readIamAccessKey,
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
				ForceNew: true,
			},
		},
	}
}

func createIamAccessKey(d *schema.ResourceData, m interface{}) error {
	if err := real_resource(m).Create(d, real_config(m)); err != nil {
		return err
	}
	secret := d.Get("secret").(string)
	file := d.Get("file").(string)
	if err := ioutil.WriteFile(file, []byte(secret), 0600); err != nil {
		return err
	}
	if err := d.Set("secret", ""); err != nil {
		return err
	}
	return nil
}

func readIamAccessKey(d *schema.ResourceData, m interface{}) error {
	return real_resource(m).Read(d, real_config(m))
}

func deleteIamAccessKey(d *schema.ResourceData, m interface{}) error {
	return real_resource(m).Delete(d, real_config(m))
}

func real_resource(m interface{}) *schema.Resource {
	return m.(*Wrapper).resource("aws_iam_access_key")
}

func real_config(m interface{}) interface{} {
	return m.(*Wrapper).config
}
