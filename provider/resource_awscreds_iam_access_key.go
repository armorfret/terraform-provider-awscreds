package provider

import (
	"context"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIamAccessKey() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIamAccessKeyCreate,
		ReadContext:   resourceIamAccessKeyRead,
		DeleteContext: resourceIamAccessKeyDelete,

		Schema: map[string]*schema.Schema{
			"user": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"file": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

//revive:disable:line-length-limit
func resourceIamAccessKeyCreate(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*iam.IAM)

	user := d.Get("user").(string)
	input := &iam.CreateAccessKeyInput{
		UserName: &user,
	}

	result, err := client.CreateAccessKey(input)
	if err != nil {
		return diag.FromErr(err)
	}
	access := *result.AccessKey.AccessKeyId
	secret := *result.AccessKey.SecretAccessKey

	d.SetId(access)

	contents := access + "\n" + secret + "\n"
	file := d.Get("file").(string)
	err = ioutil.WriteFile(file, []byte(contents), 0600)
	if err != nil {
		return diag.Errorf("failed to write secret file")
	}
	return nil
}

func resourceIamAccessKeyRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*iam.IAM)

	user := d.Get("user").(string)
	input := &iam.ListAccessKeysInput{
		UserName: &user,
	}

	result, err := client.ListAccessKeys(input)
	if err != nil {
		return diag.FromErr(err)
	}

	for _, key := range result.AccessKeyMetadata {
		if *key.AccessKeyId == d.Id() {
			return nil
		}
	}

	d.SetId("")
	return nil
}

func resourceIamAccessKeyDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*iam.IAM)

	user := d.Get("user").(string)
	access := d.Id()
	input := &iam.DeleteAccessKeyInput{
		UserName:    &user,
		AccessKeyId: &access,
	}

	_, err := client.DeleteAccessKey(input)
	return diag.FromErr(err)
}
