package middleware

import (
	"belajar-api/helper"
	"belajar-api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	// untuk meneruskan request ke handler selanjutnya
	Handle http.Handler
}

func NewAuthMiddleware(handle http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handle: handle}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "RAHASIA" == request.Header.Get("X-API-Key") {
		middleware.Handle.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}

		helper.WriteToResponseBody(writer, webResponse)
	}
}
