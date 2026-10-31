resource "routeros_ip_smb_users" "guest" {
  name     = "guest"
  disabled = true
}
