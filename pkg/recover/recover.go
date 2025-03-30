package recovery

import (
	"fmt"
	"log"
	"net/http"

	"go.uber.org/zap"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				zap.L().Error(fmt.Sprintf("Recovered by error : %v", r))
				log.Println("Recovered by error : ", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
