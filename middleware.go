package mr

import (
	"context"
	"net/http"
)

type mrCtxKey string

// RepoCtxKey context key used to store/retrieve the repository
const RepoCtxKey = mrCtxKey("repo")

// Middleware creates a http middleware that injects a Repository instance in
// each request
//
// Wrap your existing handlers if you use plain net/http:
//
//     withRepo := mr.Middleware(repo)
//     http.Handle("/url", withRepo(MyHandler))
//
// Alternatively, if you use something like gorilla/mux you can use it as a
// mux.MiddlewareFunc:
//
//     r := mux.NewRouter()
//     r.Use(mr.Middleware(repo))
//     r.Handle("/url", MyHandler)
//
func Middleware(repo Repository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), RepoCtxKey, repo)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// FromCtx extracts a Repositor instance from the given context. Returns the
// instance or nil.
func FromCtx(ctx context.Context) Repository {
	if repo, ok := ctx.Value(RepoCtxKey).(Repository); ok {
		return repo
	}
	return nil
}

// C extracts a Repository from "r"s context and calls "C(name)" on it. For use
// in conjuction with "mr.Middleware()":
//
//     func MyHandler(w http.ResponseWriter, r *http.Request) {
//             var users []User
//             mr.C(r, "users").FindAll(&users)
//             json.NewEncoder(w).Encode(&users)
//     }
func C(r *http.Request, name string) Collection {
	if repo := FromCtx(r.Context()); repo != nil {
		return repo.C(name)
	}
	return nil
}
