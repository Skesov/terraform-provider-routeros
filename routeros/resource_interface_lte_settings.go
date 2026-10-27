package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// /interface/lte/settings, generated from RouterOS 7.25 by tools/ros-coverage.
func ResourceInterfaceLteSettings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/lte/settings"),
		MetaId:           PropId(Id),

		"esim_channel": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Possible values: `at`, `auto`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"firmware_path": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"link_recovery_timer": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "In seconds.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"mode": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Possible values: `auto`.",
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
