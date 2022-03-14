package provider

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// New creates a provider instance
func New() func() *schema.Provider {
	return func() *schema.Provider {
		return &schema.Provider{
			Schema: map[string]*schema.Schema{
				"region": {
					Type:     schema.TypeString,
					Optional: true,
					Default:  "us-east-1",
				},
			},
			ResourcesMap: map[string]*schema.Resource{
				"awscreds_iam_access_key": resourceIamAccessKey(),
			},
			ConfigureContextFunc: configureIamClient,
		}
	}
}

func configureIamClient(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	region := d.Get("region").(string)
	sess, err := session.NewSession(&aws.Config{
		Region: &region,
	})
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return iam.New(sess), nil
}
