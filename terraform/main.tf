
//@TODO This Terraform is only for testing purpose. We will remove it later
terraform {
  required_providers {
    aviatrix-cloud-poc = {
      source = "hashicorp.com/prd/aviatrix-cloud-poc"
    }
  }
}

provider "aviatrix-cloud-poc" {}

data "aviatrix-cloud-poc_avx" "example" {}
data "aviatrix-cloud-poc_controller" "cntrl" {}
resource "aviatrix-cloud-poc_account_user" "account" {
    email = "test@test.com"
    username = "username"
    password = "test@123"
}

output "avx_example" {
  value = data.aviatrix-cloud-poc_avx.example
}

output "test_output" {
  value = data.aviatrix-cloud-poc_controller.cntrl
}

output "account_output" {
  value = aviatrix-cloud-poc_account_user.account
}
