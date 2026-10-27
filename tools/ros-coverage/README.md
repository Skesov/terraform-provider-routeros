Compares the menu tree of live RouterOS devices with the provider schema.

`dump` walks the tree with `/console/inspect` (read-only) and `GET`, and writes
each configuration menu's settable arguments (with RouterOS help text, syntax
and enum values) and printed properties (type only) to
`schema-<version>-<board>.json`. Values are never stored; names of objects seen
on the router (interfaces, scripts, comments) are replaced with `<name>`.
Runtime menus such as `/log`, `/file` and the connection table are skipped,
see `skipPrefixes` in `dump.go`.

`compare` merges dumps and prints, per menu and resource:

- **Missing RW**: settable on the router, absent from the schema.
- **Missing RO**: printed by the router, absent from the schema (the
  `Field '...' not found in the schema` warning).
- **RO in provider**: settable on the router, computed-only in the schema.
- **Absent on router**: settable in the schema, unknown to the router (removed
  in this version, or present on other hardware only).

`gen` drafts schema entries for a resource's missing properties, or a whole
resource, example and test for a menu without one. Review its types and
descriptions before committing.

After a RouterOS upgrade, one command refreshes the dumps and the table:

```
ROS_USERNAME=... ROS_PASSWORD=... \
make coverage ROS_COVERAGE_HOSTS='https://router1 https://router2'
```

Other commands:

```
cd tools/ros-coverage
go run . dump -url https://router
go run . compare [-f /interface/wifi] [-configured] schema-7.25-*.json
go run . gen -resource routeros_ip_ssh_server schema-7.25-*.json
go run . gen -repo ../.. -path /ip/proxy -name routeros_ip_proxy schema-7.25-*.json
```
