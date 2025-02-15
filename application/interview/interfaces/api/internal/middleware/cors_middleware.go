package middleware

import (
	"Ai-HireSphere/common/middleware"
	"net/http"
)

type CorsMiddleware struct {
}

func NewCorsMiddleware() *CorsMiddleware {
	return &CorsMiddleware{}
}

func (m *CorsMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return middleware.CorsMiddleware(next)
}
