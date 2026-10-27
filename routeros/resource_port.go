package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// /port, generated from RouterOS 7.25 by tools/ros-coverage.
func ResourcePort() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/port"),
		MetaId:           PropId(Id),

		"baud_rate": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Baud rate (bits/s).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"channels": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "",
		},
		"data_bits": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Number of bits for data. Possible values: `7`, `8`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"device": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},
		"dtr": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Possible values: `off`, `on`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"flow_control": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Type of flow control. Possible values: `hardware`, `xon-xoff`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyInactive: PropInactiveRo,
		"latency_timer": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"line_state": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},
		KeyName: PropName("Name of the item to configure."),
		"parity": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Parity. Possible values: `even`, `odd`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"rts": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Possible values: `off`, `on`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"stop_bits": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Number of bits for control. Possible values: `1`, `2`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"used_by": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},
	}

	return &schema.Resource{
		CreateContext: DefaultCreateUpdate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultCreateUpdate(resSchema),
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: ImportStateCustomContext(resSchema),
		},

		Schema: resSchema,
	}
}
