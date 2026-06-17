package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Nuttachai-K/cafe-finder-api/internal/service"
)

type CafeHandler struct {
	service service.CafeService
}

func NewCafeHandler(service service.CafeService) *CafeHandler {
	return &CafeHandler{
		service: service,
	}
}

func (h *CafeHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(
			w,
			"Invalid cafe id",
			http.StatusBadRequest,
		)
		return
	}

	cafe, err := h.service.GetByID(id)
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusNotFound,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cafe); err != nil {
		http.Error(
			w,
			"Failed to encode response",
			http.StatusInternalServerError,
		)
	}
}
