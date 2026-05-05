path "database/creds/zeaz-*" { capabilities = ["read"] }
path "kv/data/zeaz/{{identity.entity.aliases.auth_kubernetes_*.metadata.service_account_namespace}}/*" { capabilities = ["read"] }
path "transit/sign/zeaz-*" { capabilities = ["update"] }
path "transit/verify/zeaz-*" { capabilities = ["update"] }
