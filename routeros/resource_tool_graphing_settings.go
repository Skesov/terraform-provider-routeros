package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// /tool/graphing, generated from RouterOS 7.25 by tools/ros-coverage.
func ResourceToolGraphingSettings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/tool/graphing"),
		MetaId:           PropId(Id),

		"page_refresh": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "How often to refresh HTML pages (in seconds). Possible values: `never`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"store_every": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "How often to store information on system drive. Possible values: `24hours`, `5min`, `hour`.",
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
