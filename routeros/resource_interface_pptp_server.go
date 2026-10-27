package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// /interface/pptp-server/server, generated from RouterOS 7.25 by tools/ros-coverage.
func ResourceInterfacePptpServer() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/pptp-server/server"),
		MetaId:           PropId(Id),

		"authentication": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Authentication algorithm.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"default_profile": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Activate the default profile.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"enabled": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Defines whether PPTP server is enabled or not.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"keepalive_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Time period after which not responding client is proclaimed. Possible values: `disabled`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"max_mru": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Maximum Receive Unit.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"max_mtu": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Maximum Transmission Unit.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"mrru": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Possible values: `disabled`.",
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
