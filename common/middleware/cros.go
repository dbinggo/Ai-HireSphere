package middleware

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// CorsMiddleware 跨域中间件
func CorsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 设置允许的源
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// 设置允许的请求头类型
		w.Header().Set("Access-Control-Allow-Headers", "*")
		// 设置允许的请求方法
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		// 允许携带凭证（如 Cookie 等，根据实际需求决定是否设置）
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			// 对于 OPTIONS 请求，直接返回 200 状态码
			httpx.WriteJson(w, http.StatusOK, map[string]interface{}{})
			return
		}

		next(w, r)
	}
}
