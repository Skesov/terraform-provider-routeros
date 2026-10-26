resource "routeros_ip_service_webserver" "settings" {
  webfig_plain = false
  graphs_plain = false
  index_plain  = false
  rest_plain   = false
  crl_plain    = false
  scep_plain   = false
  acme_plain   = false
}
