variable "primary_region" { type = string default = "us-east-1" }
variable "failover_region" { type = string default = "us-west-2" }
variable "cloudflare_zone_id" { type = string }
variable "cloudflare_api_token" { type = string sensitive = true }
