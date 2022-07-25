package xray

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jfrog/terraform-provider-shared/util"
	"github.com/jfrog/terraform-provider-shared/validator"
)

func resourceXrayLicensePolicyV2() *schema.Resource {
	var criteriaSchema = map[string]*schema.Schema{
		"banned_licenses": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "A list of OSS license names that may not be attached to a component.",
			Elem: &schema.Schema{
				Type:             schema.TypeString,
				ValidateDiagFunc: validator.LicenseType,
			},
		},
		"allowed_licenses": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "A list of OSS license names that may be attached to a component.",
			Elem: &schema.Schema{
				Type:             schema.TypeString,
				ValidateDiagFunc: validator.LicenseType,
			},
		},
		"allow_unknown": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "A violation will be generated for artifacts with unknown licenses (`true` or `false`).",
		},
		"multi_license_permissive": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Do not generate a violation if at least one license is valid in cases whereby multiple licenses were detected on the component",
		},
	}

	var actionsSchema = util.MergeMaps(
		commonActionsSchema,
		map[string]*schema.Schema{
			"custom_severity": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "High",
				Description:      "The severity of violation to be triggered if the `criteria` are met.",
				ValidateDiagFunc: validator.StringInSlice(true, "Critical", "High", "Medium", "Low"),
			},
		},
	)

	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: resourceXrayPolicyCreate,
		ReadContext:   resourceXrayPolicyRead,
		UpdateContext: resourceXrayPolicyUpdate,
		DeleteContext: resourceXrayPolicyDelete,
		Description: "Creates an Xray policy using V2 of the underlying APIs. Please note: " +
			"It's only compatible with Bearer token auth method (Identity and Access => Access Tokens)",

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: getPolicySchema(criteriaSchema, actionsSchema),
	}
}
