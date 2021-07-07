variable "numberOfBackendServers" {
  description = "Secret Key to access SberCloud"
  sensitive = true
}

data "sbercloud_images_image" "ubuntu_image" {
  name = "Ubuntu 20.04 server 64bit"
  most_recent = true
}

resource "sbercloud_compute_instance" "ecs_01" {
  count = var.numberOfBackendServers

  name = "ecs-${count.index}"
  image_id = data.sbercloud_images_image.ubuntu_image.id
  flavor_id = "s6.medium.2"
  security_groups = [sbercloud_networking_secgroup.sg_01.name]
  availability_zone = "ru-moscow-1a"
  key_pair = "KeyPair-borodin"

  user_data = file("./script.sh")

  system_disk_type = "SAS"
  system_disk_size = 20


  network {
    uuid = sbercloud_vpc_subnet.subnet_01.id
  }
}

resource "sbercloud_compute_instance" "ecs_master" {
  name = "ecs-master"
  image_id = data.sbercloud_images_image.ubuntu_image.id
  flavor_id = "s6.medium.2"
  security_groups = [sbercloud_networking_secgroup.sg_01.name]
  availability_zone = "ru-moscow-1a"
  key_pair = "KeyPair-borodin"

  user_data = file("./script.sh")

  system_disk_type = "SAS"
  system_disk_size = 20


  network {
    uuid = sbercloud_vpc_subnet.subnet_01.id
  }
}