package middlewares

import (
	"net/http"

	"github.com/strytm/StrytmJwt/jwt"
	"github.com/strytm/StrytmMessages/messages"
	"github.com/strytm/StrytmUtils/utils"
)

var messageModel = messages.MessageModelStruct{}

func AuthMiddleware(key string) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			messageModel.ResponseWriter = w

			tokenFromHeader, err := utils.ExteraxtTokenFromHeader("Authorization", r)

			if err != nil {
				messageModel.ShowStringMessageAndStatusCode(err.Error(), http.StatusBadRequest, nil)

			} else {
				_, jwtErr := jwt.IsValidToken(tokenFromHeader, key)
				if jwtErr != nil {
					messageModel.ShowStringMessageAndStatusCode(jwtErr.Error(), http.StatusUnauthorized, nil)
					return
				}
			}

			h.ServeHTTP(w, r)
		})
	}
}
