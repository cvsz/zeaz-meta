variable "primary_region" {
  type    = string
  default = "us-east-1"
}

variable "failover_region" {
  type    = string
  default = "us-west-2"
}

variable "cloudflare_zone_id" {
  type = string
}

variable "cloudflare_api_token" {
  type      = string
  sensitive = true
}

variable "enable_gcp_hybrid" {
  type    = bool
  default = false
}

variable "gcp_project_id" {
  type    = string
  default = "zeaz-prod"
}

variable "gcp_region" {
  type    = string
  default = "us-central1"
}
