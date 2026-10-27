package routeros

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  ".id": "*1",
  "active-queue": "only-hardware-queue",
  "default-queue": "only-hardware-queue",
  "interface": "ether1",
  "queue": "only-hardware-queue"
}
*/

// ResourceQueueInterface manages the queue type of an interface. Every
// interface has exactly one row, keyed by `interface`; rows cannot be added
// or removed.
// https://help.mikrotik.com/docs/spaces/ROS/pages/328088/Queues#Queues-InterfaceQueue
func ResourceQueueInterface() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/queue/interface"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("interface"),

		"active_queue": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Queue type in use.",
		},
		"default_queue": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Queue type the interface uses by default.",
		},
		"interface": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "Name of the interface.",
		},
		"queue": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Queue type to use on the interface.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}

	resCreateUpdate := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		item, metadata := TerraformResourceDataToMikrotik(resSchema, d)

		res, err := ReadItemsFiltered([]string{"interface=" + d.Get("interface").(string)}, metadata.Path, m.(Client))
		if err != nil {
			return diag.FromErr(err)
		}
		if len(*res) != 1 {
			return diag.FromErr(fmt.Errorf("queue interface %q: found %d rows", d.Get("interface").(string), len(*res)))
		}

		d.SetId((*res)[0].GetID(Id))
		item[".id"] = d.Id()

		var resUrl string
		if m.(Client).GetTransport() == TransportREST {
			resUrl = "/set"
		}
		if err := m.(Client).SendRequest(crudPost, &URL{Path: metadata.Path + resUrl}, item, nil); err != nil {
			return diag.FromErr(err)
		}

		return ResourceRead(ctx, resSchema, d, m)
	}

	return &schema.Resource{
		CreateContext: resCreateUpdate,
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: resCreateUpdate,
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: ImportStateCustomContext(resSchema),
		},

		Schema: resSchema,
	}
}
