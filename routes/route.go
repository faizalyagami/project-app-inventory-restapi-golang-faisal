package routes

import (
	"database/sql"
	"net/http"
	"project-app-inventory-restapi-golang-faisal/handler"
	"project-app-inventory-restapi-golang-faisal/repository"
	"project-app-inventory-restapi-golang-faisal/service"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/lib/pq"
)


func SetUpRouter() http.Handler  {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	db, _ :=sql.Open("postgres", "user=postgres password=postgres dbname=system_inventory sslmode=disable")

	itemRepo := repository.NewItemRepository(db)
	itemService := service.NewItemService(itemRepo)
	ItemHandler := handler.NewItemHandler(itemService)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	warehouseRepo := repository.NewWarehouseRepository(db)
	warehouseService := service.NewWarehouseService(warehouseRepo)
	warehouseHandler := handler.NewWarehouseHandler(warehouseService)
	
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Route("/items", func (r chi.Router)  {
		r.Get("/", ItemHandler.GetAll)
		r.Get("/{id}", ItemHandler.GetByID)
		r.Post("/", ItemHandler.Create)
		r.Put("/{id}", ItemHandler.Update)
		r.Delete("/{id}", ItemHandler.Delete)
	})

	r.Route("/categories", func(r chi.Router) {
		r.Get("/", categoryHandler.GetAll)
		r.Post("/", categoryHandler.Create)
		r.Put("/{id}", categoryHandler.Update)
		r.Delete("/{id}", categoryHandler.Delete)
	})

	r.Route("/warehouses", func(r chi.Router) {
		r.Get("/", warehouseHandler.GetAll)
		r.Post("/", warehouseHandler.Create)
		r.Put("/{id}", warehouseHandler.Update)
		r.Delete("/{id}", warehouseHandler.Delete)
	})
	return  r
}