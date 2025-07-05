package handler

import (
	"encoding/json"
	"net/http"
	"project-app-inventory-restapi-golang-faisal/model"
	"project-app-inventory-restapi-golang-faisal/service"
	"project-app-inventory-restapi-golang-faisal/utils"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
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

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	//decode dan tolak field asing
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&item); err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "error",
			"message": "Format JSON tidak valid: " + err.Error(),
		})
		return
	}
	
	//validation input
	if err := utils.Validate.Struct(item); err != nil {
		var validationErrors []string
		for _, e := range err.(validator.ValidationErrors) {
			field := e.Field()
			switch field {
			case "Name":
				validationErrors = append(validationErrors, "Nama barang wajib diisi.")
			default:
				validationErrors = append(validationErrors, field+" tidak valid")
			}
		}
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "error",
			"message": "Validasi Gagal!",
			"errors": validationErrors,
		})
		return
	}
	//simpan ke DB
	if err := h.service.Create(&item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"message": "Item berhasil ditambahkan",
		"data": item,
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

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status" : "success",
		"message": "Item berhasil diperbaharui",
		"data": item,
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