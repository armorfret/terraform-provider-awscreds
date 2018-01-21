package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"io/ioutil"
)

var keysToSuppress = []string{"secret", "ses_smtp_password"}

func resourceIamAccessKey() *schema.Resource {
	return &schema.Resource{
		Create: createIamAccessKey,
		Read:   readIamAccessKey,
		Delete: deleteIamAccessKey,
		Schema: schemaIamAccessKey(),
	}
}

func schemaIamAccessKey() map[string]*schema.Schema {
	resource := awsProvider().ResourcesMap["aws_iam_access_key"]
	iamschema := resource.Schema
	iamschema["file"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	}
	return iamschema
}

func createIamAccessKey(d *schema.ResourceData, m interface{}) error {
	if err := realResource(m).Create(d, realConfig(m)); err != nil {
		return err
	}

	access := d.Id()
	secret := d.Get("secret").(string)

	for _, key := range keysToSuppress {
		if err := d.Set(key, ""); err != nil {
			return err
		}
	}

	contents := access + "\n" + secret + "\n"
	file := d.Get("file").(string)
	return ioutil.WriteFile(file, []byte(contents), 0600)
}

func readIamAccessKey(d *schema.ResourceData, m interface{}) error {
	return realResource(m).Read(d, realConfig(m))
}

func deleteIamAccessKey(d *schema.ResourceData, m interface{}) error {
	return realResource(m).Delete(d, realConfig(m))
}

func realResource(m interface{}) *schema.Resource {
	return m.(wrapper).resource("aws_iam_access_key")
}

func realConfig(m interface{}) interface{} {
	return m.(wrapper).config
}
