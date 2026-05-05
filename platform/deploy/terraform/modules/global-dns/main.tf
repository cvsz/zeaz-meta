variable "zone_id" { type = string }
variable "wildcard_domain" { type = string }
variable "primary_endpoint" { type = string }
variable "failover_endpoint" { type = string }

resource "cloudflare_record" "wildcard_primary" {
  zone_id = var.zone_id
  name    = var.wildcard_domain
  type    = "CNAME"
  content = var.primary_endpoint
  proxied = true
  ttl     = 1
}

resource "cloudflare_record" "api_primary" {
  zone_id = var.zone_id
  name    = "api.zeaz.dev"
  type    = "CNAME"
  content = var.primary_endpoint
  proxied = true
  ttl     = 1
}

resource "cloudflare_record" "failover" {
  zone_id = var.zone_id
  name    = "failover.zeaz.dev"
  type    = "CNAME"
  content = var.failover_endpoint
  proxied = true
  ttl     = 1
}

resource "cloudflare_ruleset" "waf_rate_limit" {
  zone_id = var.zone_id
  name    = "zeaz-dev-api-waf-rate-limit"
  kind    = "zone"
  phase   = "http_ratelimit"

  rules {
    action      = "block"
    description = "Block abusive API clients while preserving tenant burst traffic"
    expression  = "(http.host eq \"api.zeaz.dev\" or ends_with(http.host, \".zeaz.dev\")) and http.request.uri.path contains \"/v1/\""
    ratelimit {
      characteristics     = ["cf.colo.id", "ip.src"]
      period              = 60
      requests_per_period = 1200
      mitigation_timeout  = 600
    }
  }
}

resource "cloudflare_ruleset" "managed_waf" {
  zone_id = var.zone_id
  name    = "zeaz-dev-managed-waf"
  kind    = "zone"
  phase   = "http_request_firewall_managed"

  rules {
    action      = "execute"
    description = "Cloudflare managed ruleset"
    expression  = "true"
    action_parameters { id = "efb7b8c949ac4650a09736fc376e9aee" }
  }
}

resource "cloudflare_ruleset" "ssl_redirect" {
  zone_id = var.zone_id
  name    = "zeaz-dev-enterprise-ssl"
  kind    = "zone"
  phase   = "http_request_dynamic_redirect"
  rules {
    action      = "redirect"
    description = "Force HTTPS"
    expression  = "http.request.scheme eq \"http\""
    action_parameters {
      from_value {
        status_code = 301
        target_url { expression = "concat(\"https://\", http.host, http.request.uri.path)" }
      }
    }
  }
}
