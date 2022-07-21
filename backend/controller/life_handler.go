package controller

import (
	"encoding/json"
	"net/http"

	"github.com/syougo1209/ulife-share/model/entity"
)

type lifeHandler struct {
	LUsecase application.LifeUsecase
}

func NewLifeHandler(w http.ResponseWriter, r *http.Request) {
	handler := lifeHandler{}
	switch r.Method {
	case "GET":
		handler.FetchLife(w, r)
	default:
		w.WriteHeader(405)
	}
}

func (l *lifeHandler) FetchLife(w http.ResponseWriter, r *http.Request) {
	var res []entity.Life
	for i := 0; i < 5; i++ {
		res = append(res, entity.Life{Id: i + 1, Content: "newcontent"})
	}
	output, _ := json.MarshalIndent(res, "", "\t\t")
	w.Write(output)
}
