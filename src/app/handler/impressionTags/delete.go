package impressionTags

import (
	"encoding/json"
	"net/http"
	"newview-backend-go/app/handler/httperror"
	"newview-backend-go/app/handler/request"
)

type ApiReturn struct {
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := request.IDOf(r)
	if err != nil {
		httperror.BadRequest(w, err)
	}

	impressionTagRepository := h.app.Dao.ImpressionTag()
	_, err = impressionTagRepository.DeleteById(ctx, uint64(id))

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ApiReturn{}); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
