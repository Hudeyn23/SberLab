variable "elbName" {
  description = "Name of ELB"
}

variable "listenerName" {
  description = "Name of Listener"
}

resource "sbercloud_vpc_eip" "elb_eip" {
  publicip {
    type = "5_bgp"
  }
  bandwidth {
    name = "eip-for-${var.elbName}"
    size = 5
    share_type = "PER"
    charge_mode = "bandwidth"
  }
}

resource "sbercloud_lb_loadbalancer" "elb_01" {
  name = var.elbName
  vip_subnet_id = sbercloud_vpc_subnet.subnet_01.subnet_id
}

resource "sbercloud_networking_eip_associate" "elb_eip_associate" {
  public_ip = sbercloud_vpc_eip.elb_eip.address
  port_id = sbercloud_lb_loadbalancer.elb_01.vip_port_id
}

resource "sbercloud_lb_listener" "listener_01" {
  name = var.listenerName
  protocol = "TCP"
  protocol_port = 80
  loadbalancer_id = sbercloud_lb_loadbalancer.elb_01.id
}

resource "sbercloud_lb_pool" "backend_pool" {
  name = "Backend servers group for ELB"
  protocol = "TCP"
  lb_method = "ROUND_ROBIN"
  listener_id = sbercloud_lb_listener.listener_01.id
}

resource "sbercloud_lb_monitor" "elb_health_check" {
  name = "Health check for ECS"
  type = "TCP"
  url_path = "/"
  expected_codes = "200-202"
  delay = 10
  timeout = 5
  max_retries = 3
  pool_id = sbercloud_lb_pool.backend_pool.id
}

resource "sbercloud_lb_member" "backend_server" {
  count = var.numberOfBackendServers

  address = sbercloud_compute_instance.ecs_01[count.index].access_ip_v4

  protocol_port = var.serverPort
  pool_id = sbercloud_lb_pool.backend_pool.id
  subnet_id = sbercloud_vpc_subnet.subnet_01.subnet_id

  depends_on = [
    sbercloud_lb_monitor.elb_health_check
  ]
}