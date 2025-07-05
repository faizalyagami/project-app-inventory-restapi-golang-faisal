package handler

import (
	"encoding/json"
	"net/http"
	"project-app-inventory-restapi-golang-faisal/middleware"
	"project-app-inventory-restapi-golang-faisal/model"
	"project-app-inventory-restapi-golang-faisal/service"
	"strconv"

	"github.com/go-chi/chi"
)

type SaleHandler struct {
	service service.SaleService
}

func NewSaleHandler(s service.SaleService) *SaleHandler {
	return &SaleHandler{s}
}

type saleRequest struct {
	UserID int64 `json:"user_id"`
	Items  []model.SaleItem `json:"items"`
}

func (h *SaleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req saleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	//ambil dari user context
	user := middleware.GetUserFromContext(r)
	if user == nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}


	sale := &model.Sale{
		UserID: req.UserID,
	}

	err := h.service.CreateSale(sale, req.Items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Transaksi penjualan berhasil",
		"sale_id": sale.ID,
	})
}

func (h *SaleHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	sales, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(sales)
}

func (h *SaleHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	sale, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, "sale not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(sale)
}