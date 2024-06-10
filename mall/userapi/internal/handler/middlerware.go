package handler

import "net/http"

type UserMiddleWare struct {
}

func NewUserMiddleware() *UserMiddleWare {
	return &UserMiddleWare{}
}

func (*UserMiddleWare) LoginAndReg(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	}
}
