variable "zone_id" { type = string }
variable "wildcard_domain" { type = string }
variable "primary_endpoint" { type = string }
variable "failover_endpoint" { type = string }

resource "cloudflare_record" "wildcard_primary" {
  zone_id = var.zone_id
  name    = var.wildcard_domain
  type    = "CNAME"
  value   = var.primary_endpoint
  proxied = true
  ttl     = 1
}

resource "cloudflare_ruleset" "ssl_security" {
  zone_id = var.zone_id
  name    = "zeaz-dev-enterprise-ssl"
  kind    = "zone"
  phase   = "http_request_dynamic_redirect"
  rules {
    action = "redirect"
    expression = "http.request.scheme eq \"http\""
    action_parameters { from_value { status_code = 301 target_url { expression = "concat(\"https://\", http.host, http.request.uri.path)" } } }
  }
}
