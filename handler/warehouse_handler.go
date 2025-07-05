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

type WarehouseHandler struct {
	service service.WarehouseService
}

func NewWarehouseHandler(s service.WarehouseService) *WarehouseHandler {
	return &WarehouseHandler{s}
}

func (h *WarehouseHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	warehouses, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(warehouses)
}
func (h *WarehouseHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	warehouse, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, "Warehouse not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(warehouse)
}
func (h *WarehouseHandler) Create(w http.ResponseWriter, r *http.Request) {
	var warehouse model.Warehouse

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}
	
	//decode dan tolak field asing
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&warehouse); err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "error",
			"message": "Format JSON tidak valid: " + err.Error(),
		})
		return
	}
	
	//validation input
	if err := utils.Validate.Struct(warehouse); err != nil {
		var validationErrors []string
		for _, e := range err.(validator.ValidationErrors) {
			field := e.Field()
			switch field {
			case "Name":
				validationErrors = append(validationErrors, "Nama gudang wajib diisi.")
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
	if err := h.service.Create(&warehouse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Warehouse berhasil ditambah!",
	})
}
func (h *WarehouseHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	var warehouse model.Warehouse

	if err := json.NewDecoder(r.Body).Decode(&warehouse); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	warehouse.ID = id
	if err := h.service.Update(&warehouse); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := utils.Validate.Struct(warehouse); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "warehouse berhasil diubah!",
	})
}
func (h *WarehouseHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err := h.service.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "warehouse berhasil dihapus",
	})
}