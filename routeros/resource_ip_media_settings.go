package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// /ip/media/settings, generated from RouterOS 7.25 by tools/ros-coverage.
func ResourceIpMediaSettings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/media/settings"),
		MetaId:           PropId(Id),

		"thumbnails": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Comma separated list of filenames (should end with .jpg). If name of a file in a directory matches any filename in the list, that file is treated as a thumbnail to any media file in that directory.",
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
