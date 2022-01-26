package accounts

import (
	"encoding/json"
	"net/http"

	"newview-backend-go/app/domain/object"
	"newview-backend-go/app/handler/httperror"
)

// Request body for `POST /v1/accounts`
type AddRequest struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
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
	account := new(object.Account)
	account.Username = req.Username
	if err := account.SetPassword(req.Password); err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	// repositoryを呼び出し，CreateNewUser関数の実行
	accountRepository := h.app.Dao.Account()
	result, err := accountRepository.CreatNewUser(ctx, account)
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
	account, err = accountRepository.FindById(ctx, id)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	// 取得したaccount情報をjson化して返す
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(account); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
