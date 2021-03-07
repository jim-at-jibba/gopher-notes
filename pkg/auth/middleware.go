package auth

import (
	"context"
	"fmt"
	"github.com/jim-at-jibba/gopher-notes/pkg/jwt"
	"github.com/jim-at-jibba/gopher-notes/pkg/model"
	"net/http"

	"github.com/jim-at-jibba/gopher-notes/pkg/service"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware(userService service.UserService) func(http.Handler) http.Handler {
	fmt.Println("In auth middleware")
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			fmt.Println("Header %v:", header)
			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			tokenStr := header
			username, err := jwt.ParseToken(tokenStr)
			fmt.Println("Username %v:", username)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			// check if user exists in db
			id, err := userService.GetUserIdByUsername(username)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			var user model.User
			user.ID = id
			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *model.User {
	raw, _ := ctx.Value(userCtxKey).(*model.User)
	return raw
}
