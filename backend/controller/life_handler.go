package controller

import (
	"net/http"
)

type LifeHandler struct {
}

func NewLifeHandler(w http.ResponseWriter, r *http.Request) {
	handler := LifeHandler{}
	switch r.Method {
	case "GET":
		handler.FetchLife(w, r)
	default:
		w.WriteHeader(405)
	}
}

func (l *LifeHandler) FetchLife(w http.ResponseWriter, r *http.Request) {
	var res []entity.life
}
