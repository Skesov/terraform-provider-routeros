package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// /ip/cloud/back-to-home-file/settings, generated from RouterOS 7.25 by tools/ros-coverage.
func ResourceIpCloudBackToHomeFileSettings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/cloud/back-to-home-file/settings"),
		MetaId:           PropId(Id),

		"certificate": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},
		"enabled": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "",
		},
		"prefer_relay_code": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}

	return &schema.Resource{
		CreateContext: DefaultSystemCreate(resSchema),
		ReadContext:   DefaultSystemRead(resSchema),
		UpdateContext: DefaultSystemUpdate(resSchema),
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
