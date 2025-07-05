package middleware

import (
	"context"
	"net/http"
	"project-app-inventory-restapi-golang-faisal/model"
	"project-app-inventory-restapi-golang-faisal/repository"
)

type contextKey string

const userContextKey = contextKey("user")

func WithUser(user *model.User, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), userContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserFromContext(r *http.Request) *model.User {
	user, ok := r.Context().Value(userContextKey).(*model.User)
	if ! ok {
		return nil
	}
	return user
}

func RoleMiddleware(allowedRoles ...string) func (http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user := GetUserFromContext(r)
			if user == nil {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return 
			}
			for _, role := range allowedRoles {
				if user.Role == role {
					next.ServeHTTP(w, r)
					return 
				}
			}
			http.Error(w, "forbidden - role not allowed", http.StatusForbidden)
		})
	}
}

func LoadUser(repo repository.UserRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("username")
			if err != nil {
				next.ServeHTTP(w, r)
				return 
			}
			user, err := repo.GetByUsername(cookie.Value)
			if err != nil {
				next.ServeHTTP(w, r)
				return 
			}
			ctx := context.WithValue(r.Context(), userContextKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}