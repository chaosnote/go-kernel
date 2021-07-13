package form

import (
	"net/http"

	"github.com/chaosnote/go-kernel/net/http/response"
)

// form 格式驗證
//

// Handler ...
func Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if e := r.ParseForm(); e != nil {

			response.NewData(response.OK, e.Error()).Write(w, r) // http.Error(w, fmt.Sprintf("ParseForm() err: %v", err), http.StatusForbidden)

			return
		}

		next.ServeHTTP(w, r)
	})
}
