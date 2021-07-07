variable "accessKey" {
  description = "Access Key to access SberCloud"
  sensitive = true
}

variable "secretKey" {
  description = "Secret Key to access SberCloud"
  sensitive = true
}

variable "projectName" {
  description = "IAM project where to deploy infrastructure"
}

terraform {
  required_providers {
    sbercloud = {
      source = "sbercloud-terraform/sbercloud"
    }
  }
}

provider "sbercloud" {
  auth_url = "https://iam.ru-moscow-1.hc.sbercloud.ru/v3"
  region = "ru-moscow-1"

  access_key = var.accessKey
  secret_key = var.secretKey
  project_name = var.projectName
}