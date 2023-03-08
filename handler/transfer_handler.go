package handler

import (
	"banking/model"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type TransferHandler struct {
	model.TransferUseCase
	Ctx context.Context
}

func (trans *TransferHandler) ListTransferByAccount(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("_id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ResponseWithJSON(w, http.StatusBadRequest, map[string]string{"error": "_id not true"})
		return
	}

	transfers, err := trans.TransferUseCase.ListTransferByAccount(trans.Ctx, objectID)
	if err != nil {
		ResponseWithJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}

	ResponseWithJSON(w, http.StatusOK, transfers)
}

func (trans *TransferHandler) Transfer(w http.ResponseWriter, r *http.Request) {
	var transfer model.TransferRequest
	err := json.NewDecoder(r.Body).Decode(&transfer)
	if err != nil {
		ResponseWithJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	insertedID, err := trans.TransferUseCase.Transfer(trans.Ctx, &transfer)
	if err != nil {
		ResponseWithJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	ResponseWithJSON(w, http.StatusCreated, map[string]interface{}{"_id": insertedID})
}
