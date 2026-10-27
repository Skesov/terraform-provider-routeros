package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// /system/package/update, generated from RouterOS 7.25 by tools/ros-coverage.
func ResourceSystemPackageUpdate() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/package/update"),
		MetaId:           PropId(Id),

		"channel": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Possible values: `development`, `long-term`, `stable`, `testing`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"check_certificate": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Https certificate validation. Possible values: `no`, `yes`, `yes-without-crl`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"installed_version": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},
		"ip_version": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Possible values: `auto`, `ipv4`, `ipv6`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"latest_version": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},
		"mode": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Possible values: `http`, `https`.",
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
