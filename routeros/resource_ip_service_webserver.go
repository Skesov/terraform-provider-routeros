package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  "acme-plain": "false",
  "crl-plain": "true",
  "graphs-plain": "true",
  "graphs-secure": "true",
  "index-plain": "true",
  "index-secure": "true",
  "rest-plain": "true",
  "rest-secure": "true",
  "scep-plain": "true",
  "webfig-plain": "true",
  "webfig-secure": "true"
}
*/

// https://manual.mikrotik.com/docs/cli-reference/ip/service/webserver/
func ResourceIpServiceWebserver() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/service/webserver"),
		MetaId:           PropId(Id),

		"acme_plain": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Answer ACME HTTP-01 challenge requests at `/.well-known/acme-challenge/` over HTTP.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"crl_plain": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Serve certificate revocation list downloads at `/crl/` over HTTP.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"graphs_plain": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Serve Graphs at `/graphs/` over HTTP.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"graphs_secure": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Serve Graphs at `/graphs/` over HTTPS.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"index_plain": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Serve the home (login) page and static files such as `/favicon.png` over HTTP. Takes effect only when `webfig_plain` and `graphs_plain` are also `false`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"index_secure": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Serve the home (login) page and static files such as `/favicon.png` over HTTPS. Takes effect only when `webfig_secure` and `graphs_secure` are also `false`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"rest_plain": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Serve the REST API at `/rest` over HTTP.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"rest_secure": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Serve the REST API at `/rest` over HTTPS.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"scep_plain": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Serve the SCEP server at `/scep/` over HTTP.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"webfig_plain": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Serve WebFig at `/webfig/` over HTTP.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"webfig_secure": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Serve WebFig at `/webfig/` over HTTPS.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}

	return &schema.Resource{
		Description: "Enables or disables the individual parts of the built-in web server (WebFig, Graphs, REST API, " +
			"CRL, SCEP, ACME) separately for HTTP (`www`) and HTTPS (`www-ssl`). A disabled part answers `403 Forbidden`. " +
			"The port, allowed addresses and certificate are set with `routeros_ip_service`.",
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
