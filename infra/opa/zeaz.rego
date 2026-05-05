package zeaz.authz

default allow := false

allow if {
  input.tenant_id != ""
  input.user_id != ""
  input.action in input.scopes
  not input.break_glass
}

allow if {
  input.break_glass
  input.action == "audit:read"
  input.ticket_id != ""
}
