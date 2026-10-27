package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// /system/health/settings, generated from RouterOS 7.25 by tools/ros-coverage.
func ResourceSystemHealthSettings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/health/settings"),
		MetaId:           PropId(Id),

		"cpu_overtemp_check": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"cpu_overtemp_startup_delay": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"cpu_overtemp_threshold": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"fan_control_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"fan_full_speed_temp": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"fan_min_speed_percent": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"fan_mode": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Possible values: `auto`, `manual`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"fan_on_threshold": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"fan_switch": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Possible values: `auto`, `off`, `on`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"fan_target_temp": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"use_fan": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Possible values: `auxiliary`.",
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
