# Fork notes

> AI-maintained personal fork of [terraform-routeros/terraform-provider-routeros](https://github.com/terraform-routeros/terraform-provider-routeros)
> for the owner's home routers. Changes are written by an AI agent. Not
> supported; use at your own risk.

## Why

Upstream has had no active maintainer since spring 2026: v1.99.1
(2026-03-08) is the last release and community pull requests stay unmerged
([upstream #1020](https://github.com/terraform-routeros/terraform-provider-routeros/issues/1020)). This fork carries the open pull requests
the owner's network needs and the attributes it is missing. Branch `skesov`
is based on upstream `main`, which has unreleased commits after v1.99.1.

## Merged upstream pull requests

| PR                                                                                   | Change                                                                                                                               |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| [#983](https://github.com/terraform-routeros/terraform-provider-routeros/pull/983)   | `routeros_ip_service`: key rows by `.id`, ignore dynamic per-connection rows (RouterOS 7.20+)                                        |
| [#992](https://github.com/terraform-routeros/terraform-provider-routeros/pull/992)   | RouterOS 7.22 schema: bridge `mlag_*`/`ra_guard`, detect-internet `request_interval`, wifi-cap `mld_datapath`, ipv6 pool `from_pool` |
| [#1003](https://github.com/terraform-routeros/terraform-provider-routeros/pull/1003) | `routeros_ipv6_firewall_filter`: `address_list`                                                                                      |
| [#1009](https://github.com/terraform-routeros/terraform-provider-routeros/pull/1009) | Skip fields RouterOS 7.24 returns but the schema lacks; `/ip/service` `address` → `available-from` drift                             |
| [#1013](https://github.com/terraform-routeros/terraform-provider-routeros/pull/1013) | Diff suppressors log instead of panicking on unparseable values                                                                      |
| [#1014](https://github.com/terraform-routeros/terraform-provider-routeros/pull/1014) | Accept lowercase `k` in bit/byte values                                                                                              |
| [#1024](https://github.com/terraform-routeros/terraform-provider-routeros/pull/1024) | `routeros_ipv6_firewall_filter`: do not send `place_before` on update                                                                |
| [#1025](https://github.com/terraform-routeros/terraform-provider-routeros/pull/1025) | State migration keeps the value when converting a scalar to a list                                                                   |
| [#1026](https://github.com/terraform-routeros/terraform-provider-routeros/pull/1026) | Configurable fields RouterOS 7.x reports (incl. `system_user_sshkeys.info`)                                                          |
| [#1027](https://github.com/terraform-routeros/terraform-provider-routeros/pull/1027) | Read-only fields RouterOS 7.x reports (incl. `wifi.cap`)                                                                             |

Conflict resolutions are recorded in the merge commits for #1009 and #1027.

## Added in this fork

- `routeros_wifi_channel.preamble_puncturing` (`preamble-puncturing`, 802.11be).
- `routeros_wifi_provisioning.multi_link_mode` (`multi-link-mode`: `auto`, `disabled`, `master`, `all`).
- `routeros_ip_service`: set the static row by `.id`, fill `numbers` on read so imported services plan clean.
- `routeros_ip_service_webserver`: new singleton for `/ip/service/webserver`.

### RouterOS 7.25 coverage (1.99.2-skesov.3)

[`tools/ros-coverage`](tools/ros-coverage/README.md) dumps the menu tree of a
live router and compares it with the schema. The dumps of a hAP ax² and a hAP
be lite on 7.25 and the resulting table,
[`coverage-7.25.md`](tools/ros-coverage/coverage-7.25.md), are committed;
`make coverage` refreshes them after an upgrade.

- Existing resources: every property RouterOS 7.25 reports and the schema
  lacked is added (configurable ones as optional, read-only ones as computed;
  switch-port counters skipped). `ip_ssh_server.allow_none_crypto` is removed;
  three misspelled fields are deprecated in favour of the 7.25 names; wifi band
  and width validation accept 802.11be, 6 GHz and 320 MHz values.
- New resources for settings menus and configured lists: `certificate/settings`,
  `ip/proxy`, `ip/socks`, `ip/firewall/service-port`, `queue/interface`,
  `system/ntp/client/servers`, `system/watchdog`, `system/package/update`,
  `tool/romon`, `routing/settings` and 18 more (see the table). They were
  drafted from the router's own syntax with `ros-coverage gen`; many
  descriptions are empty.

Before: 102 resources without gaps, 106 with gaps, 150 menus without a
resource. After: 192, 44, 122.

### Backlog

- Menus without a resource, none of them configured on the owner's routers:
  MPLS/LDP/TE, ISIS, PIM-SM, RIP, OSPF interface/static-neighbor, routing
  filter lists, RPKI, MACsec, mesh, VPLS, PPP/PPTP/L2TP-ether interfaces,
  `wifi/network`, kid-control, proxy/socks access lists, SMB, reverse proxy,
  `ip/arp`, `ipv6/firewall/raw`, `ipv6/dhcp-relay`, traffic-generator streams.
  Full list: rows marked **none** in the coverage table.
- Remaining gaps in existing resources (44): mostly fields absent on these
  boards (other switch chips, LTE, PoE, removed BGP `add_path_out`/
  `cluster_id`/`router_id`). Read-only in the schema but settable on the
  router: `interface_eoip.mac_address`, `interface_ethernet.orig_mac_address`,
  `interface_vxlan.ttl`, `wifi.radio_mac`, address-list `dynamic`,
  `ip_hotspot_service_port.ports`, `system_script.owner`.
- Not verified against a router: writes of the new properties and resources.
  Acceptance tests for the new settings resources exist but have not run.

## Build and install

The binary is installed into a Terraform filesystem mirror under a prerelease
version above the last upstream release, so it never collides with one:

```sh
V=1.99.2-skesov.3
OS_ARCH=$(go env GOOS)_$(go env GOARCH)
D=~/.terraform.d/plugins/registry.terraform.io/terraform-routeros/routeros/$V/$OS_ARCH
mkdir -p "$D"
CGO_ENABLED=0 go build -trimpath -o "$D/terraform-provider-routeros_v$V" .
```

Pin `version = "1.99.2-skesov.3"` in `required_providers` and run
`terraform init -upgrade`. Without a `provider_installation` block Terraform
reads `~/.terraform.d/plugins` as an implicit local mirror.

Docs regenerate with `make docs` (needs GNU sed).
