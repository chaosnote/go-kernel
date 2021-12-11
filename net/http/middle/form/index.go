package form

import (
	"net/http"
)

// form 格式驗證
//

// Handler ...
func Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if e := r.ParseForm(); e != nil {

			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)

			return
		}

		next.ServeHTTP(w, r)
	})
}
