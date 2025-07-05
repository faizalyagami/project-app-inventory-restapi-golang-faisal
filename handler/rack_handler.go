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

type RackHandler struct {
	service service.RackService
}

func NewRackHandler(s service.RackService) *RackHandler {
	return &RackHandler{s}
}

func (h *RackHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	racks, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(racks)
}

func (h *RackHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	
	rack, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(rack)
}

func (h *RackHandler) Create(w http.ResponseWriter, r *http.Request) {
	var rack model.Rack
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	//decode dan tolak field asing
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&rack); err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "error",
			"message": "Format JSON tidak valid: " + err.Error(),
		})
		return
	}
	
	//validation input
	if err := utils.Validate.Struct(rack); err != nil {
		var validationErrors []string
		for _, e := range err.(validator.ValidationErrors) {
			field := e.Field()
			switch field {
			case "Name":
				validationErrors = append(validationErrors, "Nama rack wajib diisi.")
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

	if err := h.service.Create(&rack); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"message": "rack berhasil ditambahkan!",
		"data": rack,
	})
}
func (h *RackHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	var rack model.Rack

	if err := json.NewDecoder(r.Body).Decode(&rack); err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	rack.ID = id
	if err := h.service.Update(&rack); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := utils.Validate.Struct(rack); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message" :"rack berhasil dirubah!",
	})

}
func (h *RackHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err := h.service.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "rack berhasil dihapus!",
	})

}