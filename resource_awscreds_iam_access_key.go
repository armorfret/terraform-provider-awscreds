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
	if err := proxy_method(m).Create(d, m); err != nil {
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
	return proxy_method(m).Read(d, m)
}

func deleteIamAccessKey(d *schema.ResourceData, m interface{}) error {
	return proxy_method(m).Delete(d, m)
}

func proxy_method(m interface{}) *schema.Resource {
	return m.(*schema.Provider).ResourcesMap["aws_iam_access_key"]
}
