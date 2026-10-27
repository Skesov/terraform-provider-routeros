package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// /ip/socks, generated from RouterOS 7.25 by tools/ros-coverage.
func ResourceIpSocks() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/socks"),
		MetaId:           PropId(Id),

		"auth_method": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Possible values: `password`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"connection_idle_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Connection is discarded if no data is passed for some time.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"enabled": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Defines whether SOCKS are enabled or not.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"max_connections": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Max number of simultaneous connections.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"port": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Client connections are accepted on this port.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"version": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Possible values: `4`, `5`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"vrf": {
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
