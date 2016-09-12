package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"io/ioutil"
)

var keys_to_suppress = []string{"secret", "ses_smtp_password"}

func resourceIamAccessKey() *schema.Resource {
	return &schema.Resource{
		Create: createIamAccessKey,
		Read:   readIamAccessKey,
		Delete: deleteIamAccessKey,
		Schema: schemaIamAccessKey(),
	}
}

func schemaIamAccessKey() map[string]*schema.Schema {
	resource := aws_provider().ResourcesMap["aws_iam_access_key"]
	iamschema := resource.Schema
	iamschema["file"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	}
	return iamschema
}

func createIamAccessKey(d *schema.ResourceData, m interface{}) error {
	if err := real_resource(m).Create(d, real_config(m)); err != nil {
		return err
	}

	access := d.Id()
	secret := d.Get("secret").(string)
	contents := access + "\n" + secret + "\n"

	file := d.Get("file").(string)
	if err := ioutil.WriteFile(file, []byte(contents), 0600); err != nil {
		return err
	}

	for _, key := range keys_to_suppress {
		if err := d.Set(key, ""); err != nil {
			return err
		}
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
	return m.(Wrapper).resource("aws_iam_access_key")
}

func real_config(m interface{}) interface{} {
	return m.(Wrapper).config
}
