package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// /tool/traffic-generator, generated from RouterOS 7.25 by tools/ros-coverage.
func ResourceToolTrafficGenerator() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/tool/traffic-generator"),
		MetaId:           PropId(Id),

		"latency_distribution_max": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"latency_distribution_measure_interval": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},
		"latency_distribution_samples": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "",
		},
		"measure_out_of_order": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"running": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "",
		},
		"stats_samples_to_keep": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"test_id": {
			Type:             schema.TypeInt,
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
