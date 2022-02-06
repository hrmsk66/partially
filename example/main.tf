terraform {
  required_providers {
      partially = {
          version = "0.1"
          source = "fastly/edu/partially"
      }
  }
}

provider "partially" {}

data "partially_datacenters" "dc" {}

output "datacenters" {
    value = data.partially_datacenters.dc
}

