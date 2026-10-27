package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// /system/watchdog, generated from RouterOS 7.25 by tools/ros-coverage.
func ResourceSystemWatchdog() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/watchdog"),
		MetaId:           PropId(Id),

		"auto_send_supout": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "After the support output file is automatically generated, it can be sent by email.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"automatic_supout": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "When software failure happens, a file named \"autosupout.rif\" is generated automatically.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ping_start_after_boot": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ping_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"send_email_from": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "E-mail address to send the support output file from.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"send_email_to": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "E-mail address to send the support output file to.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"send_smtp_server": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "SMTP server address to send the support output file through.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"watch_address": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The system will reboot in case 6 sequential pings to the given IP address will fail.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"watchdog_timer": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Reboot if system is unresponsive for a minute.",
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
