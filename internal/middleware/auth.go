package middleware

import (
    "context"
    "net/http"
)

var validTokens = map[string]string{
    "74edf612f393b4eb01fbc2c29dd96671": "12345",
    "d88b4b1e77c70ba780b56032db1c259b": "98765",
}

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if userID, exists := validTokens[token]; exists {
            ctx := context.WithValue(r.Context(), "userID", userID)
            next.ServeHTTP(w, r.WithContext(ctx))
        } else {
            http.Error(w, "Forbidden", http.StatusForbidden)
        }
    })
}
