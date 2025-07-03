package handler

import (
	"encoding/json"
	"net/http"
	"project-app-inventory-restapi-golang-faisal/model"
	"project-app-inventory-restapi-golang-faisal/service"
	"project-app-inventory-restapi-golang-faisal/utils"
	"strconv"

	"github.com/go-chi/chi"
)


type ItemHandler struct {
  	service service.ItemService
}

func NewItemHandler(s service.ItemService) *ItemHandler {
	return  &ItemHandler{s}
}

func (h *ItemHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	items, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(items)
}

func (h *ItemHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	item, err := h.service.GetByID(id)
	if err !=nil {
		http.Error(w, "Item Not Found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(item)
}

func (h *ItemHandler) Create(w http.ResponseWriter, r *http.Request) {
	var item model.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.Create(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := utils.Validate.Struct(item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Item berhasil ditambahkan",
	})
}

func (h *ItemHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	var item model.Item

	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	item.ID = id
	if err := h.service.Update(&item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := utils.Validate.Struct(item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Item barhasil diperbaharui",
	})
}

func (h *ItemHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err := h.service.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Item berhasil dihapus",
	})
}