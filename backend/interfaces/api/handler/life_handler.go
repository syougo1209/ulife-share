package handler

import (
	"encoding/json"
	"net/http"
	"time"

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
	var fLJ fetchLifeJSON
	fLJ = l.toFetchLifeJSON(lifes)

	res, _ := json.MarshalIndent(fLJ, "", "\t\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (l *lifeHandler) toFetchLifeJSON(lifes application.FetchOutput) fetchLifeJSON {
	lifeJSONs := make([]lifeJSON, len(lifes))
	for i, v := range lifes {
		lifeJSONs[i] = lifeJSON{
			Id:        v.Id,
			Content:   v.Content,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
	}
	return fetchLifeJSON{
		Lifes:   lifeJSONs,
		HasNext: true,
	}
}

type lifeJSON struct {
	Id        int       `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type fetchLifeJSON struct {
	Lifes   []lifeJSON `json:"lifes"`
	HasNext bool       `json:"has_next"`
}
