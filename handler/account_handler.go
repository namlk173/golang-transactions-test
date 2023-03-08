package handler

import (
	"banking/model"
	"context"
	"encoding/json"
	"net/http"
)

type AccountHandler struct {
	model.AccountUseCase
	Ctx context.Context
}

func (acc *AccountHandler) ListAccount(w http.ResponseWriter, r *http.Request) {
	accounts, err := acc.AccountUseCase.ListAccount(acc.Ctx)
	if err != nil {
		ResponseWithJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}

	ResponseWithJSON(w, http.StatusOK, accounts)
}

func (acc *AccountHandler) InsertAccount(w http.ResponseWriter, r *http.Request) {
	var account model.Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		ResponseWithJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	insertedID, err := acc.AccountUseCase.InsertAccount(acc.Ctx, &account)
	if err != nil {
		ResponseWithJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	ResponseWithJSON(w, http.StatusCreated, map[string]interface{}{"_id": insertedID})
}
