package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// /partitions, generated from RouterOS 7.25 by tools/ros-coverage.
func ResourcePartitions() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/partitions"),
		MetaId:           PropId(Id),

		"active": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "",
		},
		KeyComment: PropCommentRw,
		"fallback_to": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyName: PropName("Name of the item to configure."),
		"running": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "",
		},
		"size": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "",
		},
		"version": {
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
