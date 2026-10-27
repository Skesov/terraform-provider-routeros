package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// /tool/sms, generated from RouterOS 7.25 by tools/ros-coverage.
func ResourceToolSms() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/tool/sms"),
		MetaId:           PropId(Id),

		"allowed_number": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"channel": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"last_ussd": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},
		"polling": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Poll new SMS every 5s. It's recommended to enable polling only when there are problems with receiving new SMS otherwise.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"port": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"receive_enabled": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"remove_sent_sms_after_send": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Useful when modem automatically stores sent SMS.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"secret": {
			Type:             schema.TypeString,
			Optional:         true,
			Sensitive:        true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"sim_pin": {
			Type:             schema.TypeString,
			Optional:         true,
			Sensitive:        true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"sms_storage": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Memory for reading, writing, sending and receiving operations (<mem1>, <mem2> and <mem3> in ETSI TS 127 005), will be used in further +CPMS commands (setting effective when using AT commands only).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"status": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
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
