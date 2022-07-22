package handler

import (
	"encoding/json"
	"net/http"

	application "github.com/syougo1209/ulife-share/application/usecase"
	httputils "github.com/syougo1209/ulife-share/interfaces/api/handler/utils"
)

type lifeHandler struct {
	LUseCase application.LifeUseCase
}

func NewLifeHandler(w http.ResponseWriter, r *http.Request, LUseCase application.LifeUseCase) {
	handler := lifeHandler{LUseCase}
	switch r.Method {
	case "GET":
		handler.FetchLife(w, r)
	default:
		w.WriteHeader(405)
	}
}

func (l *lifeHandler) FetchLife(w http.ResponseWriter, r *http.Request) {
	lifes, err := l.LUseCase.Fetch()
	if err != nil {
		w.WriteHeader(httputils.GetStatusCode(err))
		return
	}

	res, _ := json.MarshalIndent(lifes, "", "\t\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
