package middlewares

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type PostReq struct {
	Token string `json:"token"`
}

func Authenticate(token string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// attempts to decode the JSON body
			decoder := json.NewDecoder(r.Body)
			var p PostReq
			err := decoder.Decode(&p)
			if err != nil {
				panic(err)
			}

			if p.Token != token {
				log.Error("Posted request Token does not match configured token. - Request killed")
				return
			}

			next.ServeHTTP(w, r)
		})
	}

}
