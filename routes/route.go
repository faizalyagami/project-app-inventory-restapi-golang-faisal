package routes

import (
	"database/sql"
	"net/http"
	"project-app-inventory-restapi-golang-faisal/handler"
	"project-app-inventory-restapi-golang-faisal/middleware"
	"project-app-inventory-restapi-golang-faisal/repository"
	"project-app-inventory-restapi-golang-faisal/service"

	chimiddleware "github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
)


func SetUpRouter() http.Handler  {
	r := chi.NewRouter()
	r.Use(chimiddleware.Logger)

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

	rackRepo := repository.NewRackRepository(db)
	rackService := service.NewRackService(rackRepo)
	rackHandler := handler.NewRackHandler(rackService)

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	saleRepo := repository.NewSaleRepository(db)
	saleService := service.NewSaleServer(saleRepo)
	saleHandler := handler.NewSaleHandler(saleService)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	r.Use(middleware.LoadUser(userRepo))
	
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Post("/login", authHandler.Login)

	r.Route("/items", func (r chi.Router)  {
		r.Use(middleware.RoleMiddleware("admin", "staff"))
		r.Get("/", ItemHandler.GetAll)
		r.Get("/{id}", ItemHandler.GetByID)
		r.Post("/", ItemHandler.Create)
		r.Put("/{id}", ItemHandler.Update)
		r.Delete("/{id}", ItemHandler.Delete)
		r.Get("/low-stock", ItemHandler.GetLowStockItems)
	})

	r.Route("/categories", func(r chi.Router) {
		r.Use(middleware.RoleMiddleware("admin", "staff"))
		r.Get("/", categoryHandler.GetAll)
		r.Get("/", categoryHandler.GetByID)
		r.Post("/", categoryHandler.Create)
		r.Put("/{id}", categoryHandler.Update)
		r.Delete("/{id}", categoryHandler.Delete)
	})

	r.Route("/warehouses", func(r chi.Router) {
		r.Use(middleware.RoleMiddleware("admin", "staff"))
		r.Get("/", warehouseHandler.GetAll)
		r.Get("/{id}", warehouseHandler.GetByID)
		r.Post("/", warehouseHandler.Create)
		r.Put("/{id}", warehouseHandler.Update)
		r.Delete("/{id}", warehouseHandler.Delete)
	})

	r.Route("/racks", func(r chi.Router) {
		r.Use(middleware.RoleMiddleware("admin", "staff"))
		r.Get("/", rackHandler.GetAll)
		r.Get("/{id}", rackHandler.GetByID)
		r.Post("/", rackHandler.Create)
		r.Put("/{id}", rackHandler.Update)
		r.Delete("/{id}", rackHandler.Delete)
	})

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", authHandler.Login)
		r.Post("/register", authHandler.Register)
	})

	r.Route("/sales", func(r chi.Router) {
		r.Use(middleware.RoleMiddleware("admin", "staff"))
		r.Post("/", saleHandler.Create)
		r.Get("/", saleHandler.GetAll)
		r.Get("/{id}", saleHandler.GetByID)
		
		r.Route("/report", func(r chi.Router) {
			r.Use(middleware.RoleMiddleware("admin", "staff", "owner"))
			r.Get("/", saleHandler.GetSalesReport)
		})
	})

	r.Route("/users", func(r chi.Router) {
		r.Use(middleware.RoleMiddleware("admin"))
		r.Get("/", userHandler.GetAll)
		r.Post("/", userHandler.Create)
		r.Put("/{id}", userHandler.Update)
		r.Delete("/{id}", userHandler.Delete)
	})

	return  r
}