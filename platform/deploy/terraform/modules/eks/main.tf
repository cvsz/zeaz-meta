variable "name" { type = string }
variable "region" { type = string }
variable "vpc_cidr" { type = string default = "10.42.0.0/16" }

locals {
  azs = ["${var.region}a", "${var.region}b", "${var.region}c"]
}

resource "aws_kms_key" "cluster" {
  description             = "${var.name} envelope encryption"
  deletion_window_in_days = 30
  enable_key_rotation     = true
}

resource "aws_cloudwatch_log_group" "cluster" {
  name              = "/aws/eks/${var.name}/cluster"
  retention_in_days = 365
  kms_key_id        = aws_kms_key.cluster.arn
}

module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "~> 5.8"

  name = var.name
  cidr = var.vpc_cidr
  azs  = local.azs

  private_subnets = [cidrsubnet(var.vpc_cidr, 4, 0), cidrsubnet(var.vpc_cidr, 4, 1), cidrsubnet(var.vpc_cidr, 4, 2)]
  public_subnets  = [cidrsubnet(var.vpc_cidr, 4, 8), cidrsubnet(var.vpc_cidr, 4, 9), cidrsubnet(var.vpc_cidr, 4, 10)]

  enable_nat_gateway     = true
  single_nat_gateway     = false
  one_nat_gateway_per_az = true
  enable_dns_hostnames   = true
  enable_dns_support     = true

  public_subnet_tags = {
    "kubernetes.io/role/elb" = "1"
  }
  private_subnet_tags = {
    "kubernetes.io/role/internal-elb" = "1"
  }
}

module "eks" {
  source  = "terraform-aws-modules/eks/aws"
  version = "~> 20.8"

  cluster_name                   = var.name
  cluster_version                = "1.30"
  cluster_endpoint_public_access = false
  cluster_endpoint_private_access = true
  enable_cluster_creator_admin_permissions = false
  cluster_enabled_log_types = ["api", "audit", "authenticator", "controllerManager", "scheduler"]
  cloudwatch_log_group_kms_key_id = aws_kms_key.cluster.arn

  vpc_id     = module.vpc.vpc_id
  subnet_ids = module.vpc.private_subnets

  cluster_encryption_config = {
    provider_key_arn = aws_kms_key.cluster.arn
    resources        = ["secrets"]
  }

  eks_managed_node_groups = {
    system = {
      instance_types = ["m7i.large"]
      min_size       = 3
      max_size       = 12
      desired_size   = 3
      labels         = { pool = "system" }
    }
    workloads = {
      instance_types = ["m7i.xlarge"]
      min_size       = 6
      max_size       = 60
      desired_size   = 9
      labels         = { pool = "workloads" }
    }
    signer = {
      instance_types = ["c7i.large"]
      min_size       = 3
      max_size       = 15
      desired_size   = 5
      labels         = { pool = "signer", dedicated = "tss" }
      taints         = { signer = { key = "dedicated", value = "tss", effect = "NO_SCHEDULE" } }
    }
  }
}

output "cluster_name" { value = module.eks.cluster_name }
output "vpc_id" { value = module.vpc.vpc_id }
output "private_subnets" { value = module.vpc.private_subnets }
output "ingress_hostname" { value = "${var.name}.elb.${var.region}.amazonaws.com" }
