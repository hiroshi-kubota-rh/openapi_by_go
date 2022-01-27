package impressionTags

import (
	"encoding/json"
	"net/http"

	"newview-backend-go/app/domain/object"
	"newview-backend-go/app/handler/httperror"
)

// Request body for `POST /v1/accounts`
type AddRequest struct {
	Name string `json:"impressionTag"`
}

// Handle request for `POST /v1/accounts`

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// jsonで渡されたuser名とpasswordをAddRequest構造体に保存
	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httperror.BadRequest(w, err)
		return
	}

	// object AccountにAddRequestの中身を保存
	ImpressionTag := new(object.ImpressionTag)
	ImpressionTag.Name = req.Name

	// repositoryを呼び出し，CreateNewUser関数の実行
	impressionTagRepository := h.app.Dao.ImpressionTag()
	result, err := impressionTagRepository.CreatNewImpressionTag(ctx, ImpressionTag)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	// 追加されたaccountのidを取得
	id, err := result.LastInsertId()
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	// idをもとにaccount情報の取得
	ImpressionTag, err = impressionTagRepository.FindById(ctx, uint64(id))
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	// 取得したaccount情報をjson化して返す
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ImpressionTag); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
