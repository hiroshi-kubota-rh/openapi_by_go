package impressionTags

import (
	"encoding/json"
	"net/http"
	"newview-backend-go/app/handler/httperror"
	"newview-backend-go/app/handler/request"
)

var username string

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := request.IDOf(r)
	if err != nil {
		httperror.BadRequest(w, err)
	}

	// impression tagからIDを取得
	impressionTagRepository := h.app.Dao.ImpressionTag()
	impressionTag, err := impressionTagRepository.FindById(ctx, uint64(id))
	//サーバーエラーは三種類用意されているので，しっかり確認するのが大事
	if err != nil {
		httperror.InternalServerError(w, err)
	}
	// jsonで返すときの形があって何ので，継承的な方法を考える．
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(impressionTag); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
