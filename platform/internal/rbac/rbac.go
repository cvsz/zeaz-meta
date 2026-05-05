package rbac

import "net/http"

type Principal struct {
	TenantID string
	UserID   string
	Scopes   map[string]bool
}

func FromRequest(r *http.Request) Principal {
	scopes := map[string]bool{}
	for _, scope := range r.Header.Values("X-Zeaz-Scope") {
		scopes[scope] = true
	}
	return Principal{TenantID: r.Header.Get("X-Tenant-ID"), UserID: r.Header.Get("X-User-ID"), Scopes: scopes}
}

func Require(scope string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := FromRequest(r)
		if p.TenantID == "" || p.UserID == "" || !p.Scopes[scope] {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
