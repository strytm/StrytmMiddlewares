package middlewares

import (
	"fmt"
	"net/http"

	"github.com/strytm/StrytmMessages/messages"
)

func RecoveryMiddleware() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			defer func() {
				if err := recover(); err != nil {
					messageModel := messages.MessageModelStruct{}
					messageModel.ResponseWriter = w

					w.Header().Set("Connection", "close")

					messageModel.ShowStringMessageAndStatusCode(fmt.Errorf("%s", err).Error(), http.StatusInternalServerError, nil)

				}
			}()

			h.ServeHTTP(w, r)
		})
	}
}
