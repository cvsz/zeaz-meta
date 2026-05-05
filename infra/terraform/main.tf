terraform {
  required_version = ">= 1.6.0"
  required_providers {
    aws = { source = "hashicorp/aws", version = "~> 5.0" }
    cloudflare = { source = "cloudflare/cloudflare", version = "~> 4.0" }
  }
}

variable "regions" { type = list(string) default = ["us-east-1", "us-west-2"] }
variable "cluster_name" { type = string default = "zeaz-v19" }
variable "cloudflare_zone_id" { type = string default = "" }

provider "aws" { region = var.regions[0] }
provider "aws" { alias = "dr" region = var.regions[1] }
provider "cloudflare" {}

module "primary_network" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "5.8.1"
  name = "${var.cluster_name}-primary"
  cidr = "10.40.0.0/16"
  azs = ["${var.regions[0]}a", "${var.regions[0]}b", "${var.regions[0]}c"]
  private_subnets = ["10.40.1.0/24", "10.40.2.0/24", "10.40.3.0/24"]
  public_subnets = ["10.40.101.0/24", "10.40.102.0/24", "10.40.103.0/24"]
  enable_nat_gateway = true
  single_nat_gateway = false
}

resource "cloudflare_record" "wildcard" {
  count = var.cloudflare_zone_id == "" ? 0 : 1
  zone_id = var.cloudflare_zone_id
  name = "*"
  type = "CNAME"
  value = "gateway.zeaz.dev"
  proxied = true
}

output "cluster_name" { value = var.cluster_name }
output "primary_vpc_id" { value = module.primary_network.vpc_id }
