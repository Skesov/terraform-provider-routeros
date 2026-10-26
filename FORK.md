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

## Build and install

The binary is installed into a Terraform filesystem mirror under a prerelease
version above the last upstream release, so it never collides with one:

```sh
V=1.99.2-skesov.1
OS_ARCH=$(go env GOOS)_$(go env GOARCH)
D=~/.terraform.d/plugins/registry.terraform.io/terraform-routeros/routeros/$V/$OS_ARCH
mkdir -p "$D"
CGO_ENABLED=0 go build -trimpath -o "$D/terraform-provider-routeros_v$V" .
```

Pin `version = "1.99.2-skesov.1"` in `required_providers` and run
`terraform init -upgrade`. A `provider_installation` block with a
`filesystem_mirror` for `terraform-routeros/routeros` makes Terraform pick it
up.

Docs regenerate with `make docs` (needs GNU sed).
