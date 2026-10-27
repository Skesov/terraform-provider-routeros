package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// /ip/firewall/service-port, generated from RouterOS 7.25 by tools/ros-coverage.
func ResourceIpFirewallServicePort() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/firewall/service-port"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("name"),

		KeyDisabled: PropDisabledRw,
		KeyInvalid:  PropInvalidRo,
		KeyName:     PropName("Name of the item to configure."),
		"ports": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Port numbers, comma separated.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"sip_direct_media": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"sip_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
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
