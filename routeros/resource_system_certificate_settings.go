package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// /certificate/settings, generated from RouterOS 7.25 by tools/ros-coverage.
func ResourceSystemCertificateSettings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/certificate/settings"),
		MetaId:           PropId(Id),

		"builtin_trust_store": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "RouterOS provided CA certificates.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"crl_download": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Auto CRL download and update.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"crl_store": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "CRL storage location. Possible values: `ram`, `system`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"crl_use": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Perform CRL checking when validating trust chain.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"current_defaults": {
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
