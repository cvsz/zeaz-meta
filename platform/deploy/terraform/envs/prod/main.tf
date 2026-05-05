terraform {
  required_version = ">= 1.6.0"
  required_providers {
    aws = { source = "hashicorp/aws", version = "~> 5.0" }
    cloudflare = { source = "cloudflare/cloudflare", version = "~> 4.0" }
    vault = { source = "hashicorp/vault", version = "~> 4.0" }
  }
  backend "s3" {
    bucket = "zeaz-terraform-state-prod"
    key    = "global/zeaz-meta.tfstate"
    region = "us-east-1"
    encrypt = true
    dynamodb_table = "zeaz-terraform-locks"
  }
}

provider "aws" { region = var.primary_region }
provider "aws" { alias = "failover" region = var.failover_region }
provider "cloudflare" { api_token = var.cloudflare_api_token }

module "primary_eks" {
  source = "../../modules/eks"
  name = "zeaz-prod-primary"
  region = var.primary_region
}
module "failover_eks" {
  source = "../../modules/eks"
  providers = { aws = aws.failover }
  name = "zeaz-prod-failover"
  region = var.failover_region
}
module "global_dns" {
  source = "../../modules/global-dns"
  zone_id = var.cloudflare_zone_id
  wildcard_domain = "*.zeaz.dev"
  primary_endpoint = module.primary_eks.ingress_hostname
  failover_endpoint = module.failover_eks.ingress_hostname
}
