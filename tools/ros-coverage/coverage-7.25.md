RouterOS 7.25 (hAP ax^2 7.25, hAP be lite 7.25)

| Path | Kind | Static rows | Resource | OK | Missing RW | Missing RO | RO in provider | Absent on router |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| /certificate | list | 6 | routeros_system_certificate | 27 | 0 | 0 | 1 | 3 |
| /certificate/crl | list | 0 | **none** | | | | | |
| /certificate/scep-server | list | 0 | routeros_certificate_scep_server | 6 | 0 | 0 | 0 | 0 |
| /certificate/scep-server | list | 0 | routeros_system_certificate_scep_server | 6 | 0 | 0 | 0 | 0 |
| /certificate/scep-server/ra | list | 0 | **none** | | | | | |
| /certificate/settings | singleton | 0 | routeros_system_certificate_settings | 5 | 0 | 0 | 0 | 0 |
| /disk | list | 0 | **none** | | | | | |
| /disk/settings | singleton | 0 | routeros_disk_settings | 5 | 0 | 0 | 0 | 0 |
| /interface | fixed-list | 9 | **none** | | | | | |
| /interface/6to4 | list | 0 | routeros_interface_6to4 | 11 | 0 | 0 | 0 | 0 |
| /interface/bonding | list | 0 | routeros_interface_bonding | 24 | 0 | 0 | 0 | 0 |
| /interface/bridge | list | 1 | routeros_bridge | 60 | 0 | 0 | 0 | 1 |
| /interface/bridge | list | 1 | routeros_interface_bridge | 60 | 0 | 0 | 0 | 1 |
| /interface/bridge/calea | list | 0 | **none** | | | | | |
| /interface/bridge/filter | list | 0 | routeros_interface_bridge_filter | 58 | 0 | 0 | 0 | 0 |
| /interface/bridge/host | list | 0 | **none** | | | | | |
| /interface/bridge/mdb | list | 0 | **none** | | | | | |
| /interface/bridge/msti | list | 0 | **none** | | | | | |
| /interface/bridge/nat | list | 0 | **none** | | | | | |
| /interface/bridge/port | list | 6 | routeros_bridge_port | 48 | 0 | 0 | 0 | 0 |
| /interface/bridge/port | list | 6 | routeros_interface_bridge_port | 48 | 0 | 0 | 0 | 0 |
| /interface/bridge/port/mst-override | list | 0 | **none** | | | | | |
| /interface/bridge/settings | singleton | 0 | routeros_interface_bridge_settings | 9 | 0 | 0 | 0 | 0 |
| /interface/bridge/vlan | list | 0 | routeros_bridge_vlan | 7 | 0 | 0 | 0 | 0 |
| /interface/bridge/vlan | list | 0 | routeros_interface_bridge_vlan | 7 | 0 | 0 | 0 | 0 |
| /interface/detect-internet | singleton | 0 | routeros_interface_detect_internet | 5 | 0 | 0 | 0 | 0 |
| /interface/dot1x/client | list | 0 | routeros_interface_dot1x_client | 8 | 0 | 0 | 0 | 0 |
| /interface/dot1x/server | list | 0 | routeros_interface_dot1x_server | 15 | 0 | 0 | 0 | 0 |
| /interface/dot1x/server/active | list | 0 | **none** | | | | | |
| /interface/eoip | list | 0 | routeros_interface_eoip | 19 | 0 | 0 | 1 | 0 |
| /interface/eoipv6 | list | 0 | **none** | | | | | |
| /interface/ethernet | fixed-list | 5 | routeros_interface_ethernet | 36 | 0 | 0 | 1 | 5 |
| /interface/ethernet/poe | fixed-list | 1 | **none** | | | | | |
| /interface/ethernet/pon | fixed-list | 0 | **none** | | | | | |
| /interface/ethernet/switch | fixed-list | 1 | routeros_interface_ethernet_switch | 7 | 0 | 0 | 0 | 7 |
| /interface/ethernet/switch/host | list | 0 | routeros_interface_ethernet_switch_host | 9 | 0 | 0 | 0 | 0 |
| /interface/ethernet/switch/port | fixed-list | 6 | routeros_interface_ethernet_switch_port | 8 | 0 | 0 | 0 | 8 |
| /interface/ethernet/switch/port-isolation | fixed-list | 6 | routeros_interface_ethernet_switch_port_isolation | 3 | 0 | 0 | 0 | 0 |
| /interface/ethernet/switch/rule | list | 0 | routeros_interface_ethernet_switch_rule | 27 | 0 | 0 | 0 | 1 |
| /interface/ethernet/switch/vlan | list | 0 | routeros_interface_ethernet_switch_vlan | 6 | 0 | 0 | 0 | 0 |
| /interface/gre | list | 0 | routeros_gre | 12 | 0 | 0 | 0 | 0 |
| /interface/gre | list | 0 | routeros_interface_gre | 12 | 0 | 0 | 0 | 0 |
| /interface/gre6 | list | 0 | routeros_interface_gre6 | 11 | 0 | 0 | 0 | 0 |
| /interface/ipip | list | 0 | routeros_interface_ipip | 12 | 0 | 0 | 0 | 0 |
| /interface/ipip | list | 0 | routeros_ipip | 12 | 0 | 0 | 0 | 0 |
| /interface/ipipv6 | list | 0 | **none** | | | | | |
| /interface/l2tp-client | list | 0 | routeros_interface_l2tp_client | 25 | 0 | 0 | 0 | 0 |
| /interface/l2tp-ether | list | 0 | **none** | | | | | |
| /interface/l2tp-server | list | 0 | **none** | | | | | |
| /interface/l2tp-server/server | singleton | 0 | routeros_interface_l2tp_server | 19 | 0 | 0 | 0 | 0 |
| /interface/list | list | 2 | routeros_interface_list | 6 | 0 | 0 | 0 | 0 |
| /interface/list/member | list | 2 | routeros_interface_list_member | 5 | 0 | 0 | 0 | 0 |
| /interface/lte | fixed-list | 0 | routeros_interface_lte | 14 | 0 | 0 | 0 | 0 |
| /interface/lte/apn | list | 0 | routeros_interface_lte_apn | 16 | 0 | 0 | 0 | 2 |
| /interface/lte/settings | singleton | 0 | routeros_interface_lte_settings | 4 | 0 | 0 | 0 | 0 |
| /interface/macsec | list | 0 | **none** | | | | | |
| /interface/macsec/profile | list | 0 | **none** | | | | | |
| /interface/macvlan | list | 0 | routeros_interface_macvlan | 13 | 0 | 0 | 0 | 1 |
| /interface/mesh | list | 0 | **none** | | | | | |
| /interface/mesh/port | list | 0 | **none** | | | | | |
| /interface/ovpn-client | list | 0 | routeros_interface_ovpn_client | 21 | 0 | 0 | 0 | 0 |
| /interface/ovpn-server | list | 0 | routeros_interface_ovpn_server | 4 | 0 | 0 | 0 | 0 |
| /interface/ovpn-server/server | list | 0 | routeros_ovpn_server | 20 | 0 | 0 | 0 | 1 |
| /interface/ppp-client | list | 0 | **none** | | | | | |
| /interface/ppp-server | list | 0 | **none** | | | | | |
| /interface/pppoe-client | list | 0 | routeros_interface_pppoe_client | 19 | 0 | 0 | 0 | 0 |
| /interface/pppoe-server | list | 0 | routeros_interface_pppoe_server | 5 | 0 | 0 | 0 | 11 |
| /interface/pppoe-server/server | list | 0 | **none** | | | | | |
| /interface/pptp-client | list | 0 | **none** | | | | | |
| /interface/pptp-server | list | 0 | **none** | | | | | |
| /interface/pptp-server/server | singleton | 0 | routeros_interface_pptp_server | 7 | 0 | 0 | 0 | 0 |
| /interface/sstp-client | list | 0 | routeros_interface_sstp_client | 25 | 0 | 0 | 0 | 0 |
| /interface/sstp-server | list | 0 | **none** | | | | | |
| /interface/sstp-server/server | singleton | 0 | routeros_interface_sstp_server | 13 | 0 | 0 | 0 | 0 |
| /interface/vlan | list | 0 | routeros_interface_vlan | 15 | 0 | 0 | 0 | 1 |
| /interface/vlan | list | 0 | routeros_vlan | 15 | 0 | 0 | 0 | 1 |
| /interface/vpls | list | 0 | **none** | | | | | |
| /interface/vrrp | list | 0 | routeros_interface_vrrp | 24 | 0 | 0 | 0 | 0 |
| /interface/vrrp | list | 0 | routeros_vrrp | 24 | 0 | 0 | 0 | 0 |
| /interface/vxlan | list | 0 | routeros_interface_vxlan | 27 | 0 | 0 | 1 | 0 |
| /interface/vxlan/vteps | list | 0 | routeros_interface_vxlan_vteps | 4 | 0 | 0 | 0 | 1 |
| /interface/wifi | list | 2 | routeros_wifi | 138 | 0 | 0 | 1 | 1 |
| /interface/wifi/aaa | list | 0 | routeros_wifi_aaa | 10 | 0 | 0 | 0 | 0 |
| /interface/wifi/access-list | list | 0 | routeros_wifi_access_list | 16 | 0 | 0 | 0 | 0 |
| /interface/wifi/cap | singleton | 0 | routeros_wifi_cap | 13 | 0 | 0 | 0 | 0 |
| /interface/wifi/capsman | singleton | 0 | routeros_wifi_capsman | 9 | 0 | 0 | 0 | 0 |
| /interface/wifi/channel | list | 4 | routeros_wifi_channel | 13 | 0 | 0 | 0 | 0 |
| /interface/wifi/configuration | list | 5 | routeros_wifi_configuration | 121 | 0 | 0 | 0 | 1 |
| /interface/wifi/datapath | list | 1 | routeros_wifi_datapath | 10 | 0 | 0 | 0 | 0 |
| /interface/wifi/interworking | list | 0 | routeros_wifi_interworking | 32 | 0 | 0 | 0 | 2 |
| /interface/wifi/network | list | 0 | **none** | | | | | |
| /interface/wifi/network/radio | list | 0 | **none** | | | | | |
| /interface/wifi/provisioning | list | 2 | routeros_wifi_provisioning | 14 | 0 | 0 | 0 | 0 |
| /interface/wifi/radio/settings | singleton | 0 | routeros_wifi_radio_settings | 2 | 0 | 0 | 0 | 0 |
| /interface/wifi/security | list | 1 | routeros_wifi_security | 35 | 0 | 0 | 0 | 0 |
| /interface/wifi/security/multi-passphrase | list | 0 | routeros_wifi_security_multi_passphrase | 7 | 0 | 0 | 0 | 0 |
| /interface/wifi/steering | list | 0 | routeros_wifi_steering | 12 | 0 | 0 | 0 | 0 |
| /interface/wifi/steering/neighbor-group | fixed-list | 1 | **none** | | | | | |
| /interface/wireguard | list | 0 | routeros_interface_wireguard | 7 | 0 | 0 | 0 | 0 |
| /interface/wireguard | list | 0 | routeros_wireguard | 7 | 0 | 0 | 0 | 0 |
| /interface/wireguard/peers | list | 0 | routeros_interface_wireguard_peer | 18 | 0 | 0 | 0 | 0 |
| /interface/wireguard/peers | list | 0 | routeros_wireguard_peer | 18 | 0 | 0 | 0 | 0 |
| /interface/xfrm | list | 0 | **none** | | | | | |
| /ip/address | list | 1 | routeros_ip_address | 12 | 0 | 0 | 0 | 0 |
| /ip/arp | list | 0 | **none** | | | | | |
| /ip/cloud | singleton | 0 | routeros_ip_cloud | 9 | 0 | 0 | 0 | 0 |
| /ip/cloud/advanced | singleton | 0 | routeros_ip_cloud_advanced | 1 | 0 | 0 | 0 | 0 |
| /ip/cloud/back-to-home-file | list | 0 | **none** | | | | | |
| /ip/cloud/back-to-home-file/settings | singleton | 0 | routeros_ip_cloud_back_to_home_file_settings | 3 | 0 | 0 | 0 | 0 |
| /ip/cloud/back-to-home-user | list | 0 | **none** | | | | | |
| /ip/dhcp-client | list | 1 | routeros_dhcp_client | 25 | 0 | 0 | 0 | 1 |
| /ip/dhcp-client | list | 1 | routeros_ip_dhcp_client | 25 | 0 | 0 | 0 | 1 |
| /ip/dhcp-client/option | list | 0 | routeros_dhcp_client_option | 6 | 0 | 0 | 0 | 0 |
| /ip/dhcp-client/option | list | 0 | routeros_ip_dhcp_client_option | 6 | 0 | 0 | 0 | 0 |
| /ip/dhcp-relay | list | 0 | routeros_ip_dhcp_relay | 10 | 0 | 0 | 0 | 0 |
| /ip/dhcp-server | list | 1 | routeros_dhcp_server | 32 | 0 | 0 | 0 | 0 |
| /ip/dhcp-server | list | 1 | routeros_ip_dhcp_server | 32 | 0 | 0 | 0 | 0 |
| /ip/dhcp-server/alert | list | 0 | **none** | | | | | |
| /ip/dhcp-server/config | singleton | 0 | routeros_ip_dhcp_server_config | 4 | 0 | 0 | 0 | 0 |
| /ip/dhcp-server/lease | list | 6 | routeros_dhcp_server_lease | 34 | 0 | 0 | 0 | 0 |
| /ip/dhcp-server/lease | list | 6 | routeros_ip_dhcp_server_lease | 34 | 0 | 0 | 0 | 0 |
| /ip/dhcp-server/matcher | list | 0 | routeros_ip_dhcp_server_option_matcher | 9 | 0 | 0 | 0 | 0 |
| /ip/dhcp-server/network | list | 1 | routeros_dhcp_server_network | 16 | 0 | 0 | 0 | 0 |
| /ip/dhcp-server/network | list | 1 | routeros_ip_dhcp_server_network | 16 | 0 | 0 | 0 | 0 |
| /ip/dhcp-server/option | list | 0 | routeros_ip_dhcp_server_option | 5 | 0 | 0 | 0 | 0 |
| /ip/dhcp-server/option/sets | list | 0 | routeros_ip_dhcp_server_option_set | 3 | 0 | 0 | 0 | 0 |
| /ip/dhcp-server/option/sets | list | 0 | routeros_ip_dhcp_server_option_sets | 3 | 0 | 0 | 0 | 0 |
| /ip/dns | singleton | 0 | routeros_dns | 19 | 0 | 0 | 0 | 0 |
| /ip/dns | singleton | 0 | routeros_ip_dns | 19 | 0 | 0 | 0 | 0 |
| /ip/dns/adlist | list | 0 | routeros_ip_dns_adlist | 5 | 0 | 0 | 0 | 0 |
| /ip/dns/forwarders | list | 0 | routeros_ip_dns_forwarders | 6 | 0 | 0 | 0 | 0 |
| /ip/dns/static | list | 4 | routeros_dns_record | 20 | 0 | 0 | 0 | 0 |
| /ip/dns/static | list | 4 | routeros_ip_dns_record | 20 | 0 | 0 | 0 | 0 |
| /ip/firewall/address-list | list | 0 | routeros_firewall_addr_list | 5 | 0 | 0 | 1 | 0 |
| /ip/firewall/address-list | list | 0 | routeros_ip_firewall_addr_list | 5 | 0 | 0 | 1 | 0 |
| /ip/firewall/calea | list | 0 | **none** | | | | | |
| /ip/firewall/filter | list | 11 | routeros_firewall_filter | 65 | 0 | 0 | 0 | 2 |
| /ip/firewall/filter | list | 11 | routeros_ip_firewall_filter | 65 | 0 | 0 | 0 | 2 |
| /ip/firewall/layer7-protocol | list | 0 | routeros_ip_firewall_layer7_protocol | 3 | 0 | 0 | 0 | 0 |
| /ip/firewall/mangle | list | 0 | routeros_firewall_mangle | 75 | 0 | 0 | 0 | 0 |
| /ip/firewall/mangle | list | 0 | routeros_ip_firewall_mangle | 75 | 0 | 0 | 0 | 0 |
| /ip/firewall/nat | list | 1 | routeros_firewall_nat | 66 | 0 | 0 | 0 | 0 |
| /ip/firewall/nat | list | 1 | routeros_ip_firewall_nat | 66 | 0 | 0 | 0 | 0 |
| /ip/firewall/raw | list | 0 | routeros_ip_firewall_raw | 52 | 0 | 0 | 0 | 0 |
| /ip/firewall/service-port | fixed-list | 10 | routeros_ip_firewall_service_port | 5 | 0 | 0 | 0 | 0 |
| /ip/hotspot | list | 0 | routeros_ip_hotspot | 8 | 0 | 0 | 0 | 0 |
| /ip/hotspot/ip-binding | list | 0 | routeros_ip_hotspot_ip_binding | 7 | 0 | 0 | 0 | 0 |
| /ip/hotspot/profile | list | 0 | routeros_ip_hotspot_profile | 27 | 0 | 0 | 0 | 1 |
| /ip/hotspot/service-port | fixed-list | 1 | routeros_ip_hotspot_service_port | 1 | 0 | 0 | 1 | 0 |
| /ip/hotspot/user | list | 0 | routeros_ip_hotspot_user | 17 | 0 | 0 | 0 | 0 |
| /ip/hotspot/user/profile | list | 0 | routeros_ip_hotspot_user_profile | 27 | 0 | 0 | 0 | 0 |
| /ip/hotspot/walled-garden | list | 0 | routeros_ip_hotspot_walled_garden | 9 | 0 | 0 | 0 | 0 |
| /ip/hotspot/walled-garden/ip | list | 0 | routeros_ip_hotspot_walled_garden_ip | 11 | 0 | 0 | 0 | 0 |
| /ip/ipsec/active-peers | list | 0 | **none** | | | | | |
| /ip/ipsec/identity | list | 0 | routeros_ip_ipsec_identity | 19 | 0 | 0 | 0 | 0 |
| /ip/ipsec/key/psk | list | 0 | **none** | | | | | |
| /ip/ipsec/key/rsa | fixed-list | 0 | **none** | | | | | |
| /ip/ipsec/mode-config | list | 0 | routeros_ip_ipsec_mode_config | 13 | 0 | 0 | 0 | 0 |
| /ip/ipsec/peer | list | 0 | routeros_ip_ipsec_peer | 11 | 0 | 0 | 0 | 0 |
| /ip/ipsec/policy | list | 0 | routeros_ip_ipsec_policy | 17 | 0 | 0 | 0 | 0 |
| /ip/ipsec/policy/group | list | 0 | routeros_ip_ipsec_policy_group | 4 | 0 | 0 | 0 | 0 |
| /ip/ipsec/profile | list | 0 | routeros_ip_ipsec_profile | 13 | 0 | 0 | 0 | 0 |
| /ip/ipsec/proposal | list | 0 | routeros_ip_ipsec_proposal | 8 | 0 | 0 | 0 | 0 |
| /ip/ipsec/settings | singleton | 0 | routeros_ip_ipsec_settings | 4 | 0 | 0 | 0 | 0 |
| /ip/kid-control | list | 0 | **none** | | | | | |
| /ip/kid-control/device | list | 0 | **none** | | | | | |
| /ip/media | list | 0 | **none** | | | | | |
| /ip/media/settings | singleton | 0 | routeros_ip_media_settings | 1 | 0 | 0 | 0 | 0 |
| /ip/nat-pmp | singleton | 0 | routeros_ip_nat_pmp | 1 | 0 | 0 | 0 | 0 |
| /ip/nat-pmp/interfaces | list | 0 | routeros_ip_nat_pmp_interfaces | 4 | 0 | 0 | 0 | 0 |
| /ip/neighbor/discovery-settings | singleton | 0 | routeros_ip_neighbor_discovery_settings | 9 | 0 | 0 | 0 | 1 |
| /ip/packing | list | 0 | **none** | | | | | |
| /ip/pool | list | 1 | routeros_ip_pool | 4 | 0 | 0 | 0 | 0 |
| /ip/pool/used | fixed-list | 12 | **none** | | | | | |
| /ip/proxy | singleton | 0 | routeros_ip_proxy | 17 | 0 | 0 | 0 | 0 |
| /ip/proxy/access | list | 0 | **none** | | | | | |
| /ip/proxy/cache | list | 0 | **none** | | | | | |
| /ip/proxy/cache-contents | list | 0 | **none** | | | | | |
| /ip/proxy/connections | fixed-list | 0 | **none** | | | | | |
| /ip/proxy/direct | list | 0 | **none** | | | | | |
| /ip/reverse-proxy | list | 0 | **none** | | | | | |
| /ip/route | list | 0 | routeros_ip_route | 19 | 0 | 0 | 0 | 1 |
| /ip/service | fixed-list | 9 | routeros_ip_service | 12 | 0 | 0 | 0 | 0 |
| /ip/service/webserver | singleton | 0 | routeros_ip_service_webserver | 11 | 0 | 0 | 0 | 0 |
| /ip/settings | singleton | 0 | routeros_ip_settings | 15 | 0 | 0 | 0 | 1 |
| /ip/smb | singleton | 0 | routeros_ip_smb | 5 | 0 | 0 | 0 | 0 |
| /ip/smb/shares | list | 0 | **none** | | | | | |
| /ip/smb/users | list | 0 | **none** | | | | | |
| /ip/socks | singleton | 0 | routeros_ip_socks | 7 | 0 | 0 | 0 | 0 |
| /ip/socks/access | list | 0 | **none** | | | | | |
| /ip/socks/connections | fixed-list | 0 | **none** | | | | | |
| /ip/socks/users | list | 0 | **none** | | | | | |
| /ip/socksify | list | 0 | **none** | | | | | |
| /ip/ssh | singleton | 0 | routeros_ip_ssh_server | 10 | 0 | 0 | 0 | 0 |
| /ip/ssh/known-hosts | list | 0 | **none** | | | | | |
| /ip/tftp | list | 0 | routeros_ip_tftp | 10 | 0 | 0 | 0 | 0 |
| /ip/tftp/settings | singleton | 0 | routeros_ip_tftp_settings | 2 | 0 | 0 | 0 | 0 |
| /ip/traffic-flow | singleton | 0 | routeros_ip_traffic_flow | 8 | 0 | 0 | 0 | 0 |
| /ip/traffic-flow/ipfix | singleton | 0 | routeros_ip_traffic_flow_ipfix | 36 | 0 | 0 | 0 | 1 |
| /ip/traffic-flow/target | list | 0 | routeros_ip_traffic_flow_target | 7 | 0 | 0 | 0 | 0 |
| /ip/upnp | singleton | 0 | routeros_ip_upnp | 3 | 0 | 0 | 0 | 0 |
| /ip/upnp/interfaces | list | 2 | routeros_ip_upnp_interfaces | 5 | 0 | 0 | 0 | 0 |
| /ip/vrf | list | 0 | routeros_ip_vrf | 5 | 0 | 0 | 0 | 0 |
| /ipv6/address | list | 1 | routeros_ipv6_address | 17 | 0 | 0 | 0 | 0 |
| /ipv6/dhcp-client | list | 1 | routeros_ipv6_dhcp_client | 28 | 0 | 0 | 0 | 0 |
| /ipv6/dhcp-client/option | list | 0 | routeros_ipv6_dhcp_client_option | 6 | 0 | 0 | 0 | 0 |
| /ipv6/dhcp-relay | list | 0 | **none** | | | | | |
| /ipv6/dhcp-relay/option | list | 0 | **none** | | | | | |
| /ipv6/dhcp-server | list | 0 | routeros_ipv6_dhcp_server | 21 | 0 | 0 | 0 | 0 |
| /ipv6/dhcp-server/binding | list | 0 | **none** | | | | | |
| /ipv6/dhcp-server/option | list | 0 | routeros_ipv6_dhcp_server_option | 4 | 0 | 0 | 0 | 0 |
| /ipv6/dhcp-server/option/sets | list | 0 | routeros_ipv6_dhcp_server_option_sets | 3 | 0 | 0 | 0 | 0 |
| /ipv6/firewall/address-list | list | 9 | routeros_ipv6_firewall_addr_list | 6 | 0 | 0 | 1 | 0 |
| /ipv6/firewall/filter | list | 23 | routeros_ipv6_firewall_filter | 59 | 0 | 0 | 0 | 0 |
| /ipv6/firewall/mangle | list | 0 | routeros_ipv6_firewall_mangle | 70 | 0 | 0 | 0 | 2 |
| /ipv6/firewall/nat | list | 0 | routeros_ipv6_firewall_nat | 56 | 0 | 0 | 0 | 2 |
| /ipv6/firewall/raw | list | 0 | **none** | | | | | |
| /ipv6/nd | list | 0 | routeros_ipv6_neighbor_discovery | 19 | 0 | 0 | 0 | 1 |
| /ipv6/nd/prefix | list | 0 | routeros_ipv6_nd_prefix | 10 | 0 | 0 | 0 | 0 |
| /ipv6/nd/prefix/default | singleton | 0 | routeros_ipv6_nd_prefix_default | 4 | 0 | 0 | 0 | 0 |
| /ipv6/nd/proxy | list | 0 | **none** | | | | | |
| /ipv6/nd/settings | singleton | 0 | routeros_ipv6_nd_settings | 2 | 0 | 0 | 0 | 0 |
| /ipv6/neighbor | list | 0 | **none** | | | | | |
| /ipv6/pool | list | 0 | routeros_ipv6_pool | 5 | 0 | 0 | 0 | 0 |
| /ipv6/route | list | 0 | routeros_ipv6_route | 17 | 0 | 0 | 0 | 1 |
| /ipv6/settings | singleton | 0 | routeros_ipv6_settings | 15 | 0 | 0 | 0 | 0 |
| /mpls/interface | list | 0 | **none** | | | | | |
| /mpls/ldp | list | 0 | **none** | | | | | |
| /mpls/ldp/accept-filter | list | 0 | **none** | | | | | |
| /mpls/ldp/advertise-filter | list | 0 | **none** | | | | | |
| /mpls/ldp/interface | list | 0 | **none** | | | | | |
| /mpls/ldp/local-mapping | list | 0 | **none** | | | | | |
| /mpls/ldp/neighbor | list | 0 | **none** | | | | | |
| /mpls/ldp/remote-mapping | list | 0 | **none** | | | | | |
| /mpls/mangle | list | 0 | **none** | | | | | |
| /mpls/settings | singleton | 0 | routeros_mpls_settings | 5 | 0 | 0 | 0 | 0 |
| /mpls/traffic-eng/interface | list | 0 | **none** | | | | | |
| /mpls/traffic-eng/path | list | 0 | **none** | | | | | |
| /mpls/traffic-eng/tunnel | list | 0 | **none** | | | | | |
| /partitions | fixed-list | 1 | routeros_partitions | 7 | 0 | 0 | 0 | 0 |
| /port | fixed-list | 2 | routeros_port | 14 | 0 | 0 | 0 | 0 |
| /port/remote-access | list | 0 | **none** | | | | | |
| /ppp/aaa | singleton | 0 | routeros_ppp_aaa | 5 | 0 | 0 | 0 | 0 |
| /ppp/l2tp-secret | list | 0 | **none** | | | | | |
| /ppp/profile | list | 0 | routeros_ppp_profile | 38 | 0 | 0 | 0 | 0 |
| /ppp/secret | list | 0 | routeros_ppp_secret | 14 | 0 | 0 | 0 | 0 |
| /queue/interface | fixed-list | 11 | routeros_queue_interface | 3 | 0 | 0 | 0 | 0 |
| /queue/simple | list | 0 | routeros_queue_simple | 24 | 0 | 0 | 0 | 0 |
| /queue/tree | list | 0 | routeros_queue_tree | 13 | 0 | 0 | 0 | 0 |
| /queue/type | list | 0 | routeros_queue_type | 51 | 0 | 0 | 0 | 0 |
| /radius | list | 0 | routeros_radius | 17 | 0 | 0 | 0 | 0 |
| /radius/incoming | singleton | 0 | routeros_radius_incoming | 3 | 0 | 0 | 0 | 0 |
| /routing/bfd/configuration | list | 0 | routeros_routing_bfd_configuration | 10 | 0 | 0 | 0 | 0 |
| /routing/bgp/connection | list | 0 | routeros_routing_bgp_connection | 57 | 0 | 0 | 0 | 3 |
| /routing/bgp/evpn | list | 0 | routeros_routing_bgp_evpn | 9 | 0 | 0 | 0 | 0 |
| /routing/bgp/instance | list | 0 | routeros_routing_bgp_instance | 11 | 0 | 0 | 0 | 0 |
| /routing/bgp/template | list | 0 | routeros_routing_bgp_template | 45 | 0 | 0 | 0 | 5 |
| /routing/bgp/vpls | list | 0 | **none** | | | | | |
| /routing/bgp/vpn | list | 0 | routeros_routing_bgp_vpn | 13 | 0 | 0 | 0 | 0 |
| /routing/fantasy | list | 0 | **none** | | | | | |
| /routing/filter/community-ext-list | list | 0 | **none** | | | | | |
| /routing/filter/community-large-list | list | 0 | **none** | | | | | |
| /routing/filter/community-list | list | 0 | **none** | | | | | |
| /routing/filter/num-list | list | 0 | **none** | | | | | |
| /routing/filter/rule | list | 0 | routeros_routing_filter_rule | 4 | 0 | 0 | 0 | 0 |
| /routing/filter/select-rule | list | 0 | **none** | | | | | |
| /routing/gmp | list | 0 | **none** | | | | | |
| /routing/id | list | 0 | routeros_routing_id | 9 | 0 | 0 | 0 | 0 |
| /routing/igmp-proxy | singleton | 0 | routeros_routing_igmp_proxy | 4 | 0 | 0 | 0 | 0 |
| /routing/igmp-proxy/interface | list | 0 | routeros_routing_igmp_proxy_interface | 10 | 0 | 0 | 0 | 0 |
| /routing/igmp-proxy/mfc | list | 0 | **none** | | | | | |
| /routing/isis/instance | list | 0 | **none** | | | | | |
| /routing/isis/interface | fixed-list | 0 | **none** | | | | | |
| /routing/isis/interface-template | list | 0 | **none** | | | | | |
| /routing/isis/lsp | fixed-list | 0 | **none** | | | | | |
| /routing/isis/neighbor | fixed-list | 0 | **none** | | | | | |
| /routing/ospf/area | list | 0 | routeros_routing_ospf_area | 9 | 0 | 0 | 0 | 0 |
| /routing/ospf/area/range | list | 0 | routeros_routing_ospf_area_range | 6 | 0 | 0 | 0 | 0 |
| /routing/ospf/instance | list | 0 | routeros_routing_ospf_instance | 17 | 0 | 0 | 0 | 0 |
| /routing/ospf/interface | list | 0 | **none** | | | | | |
| /routing/ospf/interface-template | list | 0 | routeros_routing_ospf_interface_template | 19 | 0 | 0 | 0 | 3 |
| /routing/ospf/neighbor | fixed-list | 0 | **none** | | | | | |
| /routing/ospf/static-neighbor | list | 0 | **none** | | | | | |
| /routing/pimsm/bsr/candidate | list | 0 | **none** | | | | | |
| /routing/pimsm/bsr/rp-candidate | list | 0 | **none** | | | | | |
| /routing/pimsm/igmp-interface-template | list | 0 | **none** | | | | | |
| /routing/pimsm/instance | list | 0 | **none** | | | | | |
| /routing/pimsm/interface-template | list | 0 | **none** | | | | | |
| /routing/pimsm/static-rp | list | 0 | **none** | | | | | |
| /routing/rip/instance | list | 0 | **none** | | | | | |
| /routing/rip/interface-template | list | 0 | **none** | | | | | |
| /routing/rip/keys | list | 0 | **none** | | | | | |
| /routing/rip/static-neighbor | list | 0 | **none** | | | | | |
| /routing/route | fixed-list | 14 | **none** | | | | | |
| /routing/rpki | list | 0 | **none** | | | | | |
| /routing/rule | list | 0 | routeros_routing_rule | 12 | 0 | 0 | 0 | 0 |
| /routing/settings | singleton | 0 | routeros_routing_settings | 7 | 0 | 0 | 0 | 0 |
| /routing/table | list | 0 | routeros_routing_table | 6 | 0 | 0 | 0 | 0 |
| /snmp | singleton | 0 | routeros_snmp | 12 | 0 | 0 | 0 | 0 |
| /snmp/community | list | 0 | routeros_snmp_community | 13 | 0 | 0 | 0 | 0 |
| /system/clock | singleton | 0 | routeros_system_clock | 6 | 0 | 0 | 0 | 0 |
| /system/clock/manual | singleton | 0 | routeros_system_clock_manual | 4 | 0 | 0 | 0 | 0 |
| /system/console | list | 0 | **none** | | | | | |
| /system/health/settings | singleton | 0 | routeros_system_health_settings | 11 | 0 | 0 | 0 | 0 |
| /system/identity | singleton | 0 | routeros_identity | 1 | 0 | 0 | 0 | 0 |
| /system/identity | singleton | 0 | routeros_system_identity | 1 | 0 | 0 | 0 | 0 |
| /system/keymat-provider | list | 0 | **none** | | | | | |
| /system/leds | list | 0 | routeros_system_led | 7 | 0 | 0 | 0 | 1 |
| /system/leds/settings | singleton | 0 | routeros_system_led_settings | 1 | 0 | 0 | 0 | 0 |
| /system/logging | list | 0 | routeros_system_logging | 8 | 0 | 0 | 0 | 0 |
| /system/logging/action | list | 0 | routeros_system_logging_action | 25 | 0 | 0 | 0 | 1 |
| /system/note | singleton | 0 | routeros_system_note | 3 | 0 | 0 | 0 | 0 |
| /system/ntp/client | singleton | 0 | routeros_system_ntp_client | 9 | 0 | 0 | 0 | 0 |
| /system/ntp/client/servers | list | 1 | routeros_system_ntp_client_servers | 9 | 0 | 0 | 0 | 0 |
| /system/ntp/key | list | 0 | **none** | | | | | |
| /system/ntp/server | singleton | 0 | routeros_system_ntp_server | 9 | 0 | 0 | 0 | 0 |
| /system/package/local-update/mirror | singleton | 0 | routeros_system_package_local_update_mirror | 7 | 0 | 0 | 0 | 0 |
| /system/package/local-update/update-package-source | list | 0 | **none** | | | | | |
| /system/package/update | singleton | 0 | routeros_system_package_update | 7 | 0 | 0 | 0 | 0 |
| /system/resource/hardware/usb-settings | singleton | 0 | routeros_system_resource_usb_settings | 1 | 0 | 0 | 0 | 0 |
| /system/resource/irq | fixed-list | 60 | **none** | | | | | |
| /system/resource/irq/rps | fixed-list | 5 | **none** | | | | | |
| /system/routerboard/mode-button | singleton | 0 | routeros_system_routerboard_button_mode | 3 | 0 | 0 | 0 | 0 |
| /system/routerboard/reset-button | singleton | 0 | routeros_system_routerboard_button_reset | 3 | 0 | 0 | 0 | 0 |
| /system/routerboard/settings | singleton | 0 | routeros_system_routerboard_settings | 22 | 0 | 0 | 0 | 3 |
| /system/routerboard/usb | singleton | 0 | routeros_system_routerboard_usb | 2 | 0 | 0 | 0 | 0 |
| /system/routerboard/wps-button | singleton | 0 | routeros_system_routerboard_button_wps | 3 | 0 | 0 | 0 | 0 |
| /system/scheduler | list | 1 | routeros_scheduler | 12 | 0 | 0 | 0 | 0 |
| /system/scheduler | list | 1 | routeros_system_scheduler | 12 | 0 | 0 | 0 | 0 |
| /system/script | list | 1 | routeros_system_script | 7 | 0 | 0 | 1 | 1 |
| /system/script/environment | fixed-list | 1 | **none** | | | | | |
| /system/watchdog | singleton | 0 | routeros_system_watchdog | 9 | 0 | 0 | 0 | 0 |
| /tool/bandwidth-server | singleton | 0 | routeros_tool_bandwidth_server | 6 | 0 | 0 | 0 | 0 |
| /tool/e-mail | singleton | 0 | routeros_tool_email | 8 | 0 | 0 | 0 | 0 |
| /tool/graphing | singleton | 0 | routeros_tool_graphing_settings | 2 | 0 | 0 | 0 | 0 |
| /tool/graphing/interface | list | 0 | routeros_tool_graphing_interface | 5 | 0 | 0 | 0 | 0 |
| /tool/graphing/queue | list | 0 | routeros_tool_graphing_queue | 6 | 0 | 0 | 0 | 0 |
| /tool/graphing/resource | list | 0 | routeros_tool_graphing_resource | 4 | 0 | 0 | 0 | 0 |
| /tool/mac-server | singleton | 0 | routeros_tool_mac_server | 1 | 0 | 0 | 0 | 0 |
| /tool/mac-server/mac-winbox | singleton | 0 | routeros_tool_mac_server_winbox | 1 | 0 | 0 | 0 | 0 |
| /tool/mac-server/ping | singleton | 0 | routeros_tool_mac_server_ping | 1 | 0 | 0 | 0 | 0 |
| /tool/mac-server/sessions | fixed-list | 0 | **none** | | | | | |
| /tool/netwatch | list | 0 | routeros_tool_netwatch | 35 | 0 | 0 | 0 | 2 |
| /tool/romon | singleton | 0 | routeros_tool_romon | 3 | 0 | 0 | 0 | 0 |
| /tool/romon/port | list | 0 | **none** | | | | | |
| /tool/sms | singleton | 0 | routeros_tool_sms | 11 | 0 | 0 | 0 | 0 |
| /tool/traffic-generator | singleton | 0 | routeros_tool_traffic_generator | 7 | 0 | 0 | 0 | 0 |
| /tool/traffic-generator/packet-template | list | 0 | **none** | | | | | |
| /tool/traffic-generator/port | list | 0 | **none** | | | | | |
| /tool/traffic-generator/raw-packet-template | list | 0 | **none** | | | | | |
| /tool/traffic-generator/stream | list | 0 | **none** | | | | | |
| /tool/traffic-monitor | list | 0 | **none** | | | | | |
| /user | list | 3 | routeros_system_user | 9 | 0 | 0 | 0 | 0 |
| /user/aaa | singleton | 0 | routeros_system_user_aaa | 5 | 0 | 0 | 0 | 0 |
| /user/active | list | 2 | **none** | | | | | |
| /user/group | list | 4 | routeros_system_user_group | 4 | 0 | 0 | 0 | 0 |
| /user/settings | singleton | 0 | routeros_system_user_settings | 2 | 0 | 0 | 0 | 0 |
| /user/ssh-keys | list | 3 | routeros_system_user_sshkeys | 7 | 0 | 0 | 0 | 0 |

Resources without gaps: 192, with gaps: 44. Menus without a resource: 122.

/certificate (routeros_system_certificate)
  read-only in provider: digest-algorithm
  absent on router: import sign sign_via_scep
/interface/bridge (routeros_bridge)
  absent on router: add_dhcp_option82
/interface/bridge (routeros_interface_bridge)
  absent on router: add_dhcp_option82
/interface/eoip (routeros_interface_eoip)
  read-only in provider: mac-address
/interface/ethernet (routeros_interface_ethernet)
  read-only in provider: orig-mac-address
  absent on router: cable_settings disable_running_check factory_name full_duplex poe_lldp_enabled
/interface/ethernet/switch (routeros_interface_ethernet_switch)
  absent on router: l3_hw_offloading mirror_egress_target qos_hw_offloading rspan rspan_egress_vlan_id rspan_ingress_vlan_id switch_id
/interface/ethernet/switch/port (routeros_interface_ethernet_switch_port)
  absent on router: l3_hw_offloading limit_broadcasts limit_unknown_multicasts limit_unknown_unicasts mirror_egress mirror_ingress mirror_ingress_target storm_rate
/interface/ethernet/switch/rule (routeros_interface_ethernet_switch_rule)
  absent on router: mirror_ports
/interface/lte/apn (routeros_interface_lte_apn)
  absent on router: number passthrough_subnet_selection
/interface/macvlan (routeros_interface_macvlan)
  absent on router: loop_protect_status
/interface/ovpn-server/server (routeros_ovpn_server)
  absent on router: enabled
/interface/pppoe-server (routeros_interface_pppoe_server)
  absent on router: authentication default_profile interface keepalive_timeout max_mru max_mtu max_sessions mrru one_session_per_host pppoe_over_vlan_range service_name
/interface/vlan (routeros_interface_vlan)
  absent on router: hw_offloaded
/interface/vlan (routeros_vlan)
  absent on router: hw_offloaded
/interface/vxlan (routeros_interface_vxlan)
  read-only in provider: ttl
/interface/vxlan/vteps (routeros_interface_vxlan_vteps)
  absent on router: port
/interface/wifi (routeros_wifi)
  read-only in provider: radio-mac
  absent on router: deprioritize_unii_3_4
/interface/wifi/configuration (routeros_wifi_configuration)
  absent on router: deprioritize_unii_3_4
/interface/wifi/interworking (routeros_wifi_interworking)
  absent on router: 3gpp_raw asra
/ip/dhcp-client (routeros_dhcp_client)
  absent on router: use_reconfigure
/ip/dhcp-client (routeros_ip_dhcp_client)
  absent on router: use_reconfigure
/ip/firewall/address-list (routeros_firewall_addr_list)
  read-only in provider: dynamic
/ip/firewall/address-list (routeros_ip_firewall_addr_list)
  read-only in provider: dynamic
/ip/firewall/filter (routeros_firewall_filter)
  absent on router: hw_offload routing_table
/ip/firewall/filter (routeros_ip_firewall_filter)
  absent on router: hw_offload routing_table
/ip/hotspot/profile (routeros_ip_hotspot_profile)
  absent on router: https_redirect
/ip/hotspot/service-port (routeros_ip_hotspot_service_port)
  read-only in provider: ports
/ip/neighbor/discovery-settings (routeros_ip_neighbor_discovery_settings)
  absent on router: lldp_dcbx
/ip/route (routeros_ip_route)
  absent on router: suppress_hw_offload
/ip/settings (routeros_ip_settings)
  absent on router: route_cache
/ip/traffic-flow/ipfix (routeros_ip_traffic_flow_ipfix)
  absent on router: ip_total_lenght
/ipv6/firewall/address-list (routeros_ipv6_firewall_addr_list)
  read-only in provider: dynamic
/ipv6/firewall/mangle (routeros_ipv6_firewall_mangle)
  absent on router: new_ttl ttl
/ipv6/firewall/nat (routeros_ipv6_firewall_nat)
  absent on router: randomise_ports ttl
/ipv6/nd (routeros_ipv6_neighbor_discovery)
  absent on router: dns_servers
/ipv6/route (routeros_ipv6_route)
  absent on router: suppress_hw_offload
/routing/bgp/connection (routeros_routing_bgp_connection)
  absent on router: add_path_out cluster_id router_id
/routing/bgp/template (routeros_routing_bgp_template)
  absent on router: add_path_out as_override cluster_id remove_private_as router_id
/routing/ospf/interface-template (routeros_routing_ospf_interface_template)
  absent on router: authentication_key vlink_neighbor_id vlink_transit_area
/system/leds (routeros_system_led)
  absent on router: modem_signal_treshold
/system/logging/action (routeros_system_logging_action)
  absent on router: bsd_syslog
/system/routerboard/settings (routeros_system_routerboard_settings)
  absent on router: cpu_mode memory_data_rate memory_frequency
/system/script (routeros_system_script)
  read-only in provider: owner
  absent on router: launch_trigger
/tool/netwatch (routeros_tool_netwatch)
  absent on router: http_code_max http_code_min
