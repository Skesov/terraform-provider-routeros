resource "routeros_queue_interface" "ether1" {
  interface = "ether1"
  queue     = "only-hardware-queue"
}
