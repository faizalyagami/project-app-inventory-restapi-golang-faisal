package model

type Item struct {
    ID          int64  `json:"id"`
    Name        string `json:"name" validate:"required"`
    CategoryID  int64  `json:"category_id" validate:"required"`
    RackID      int64  `json:"rack_id"`
    WarehouseID int64  `json:"warehouse_id"`
    Stock       int    `json:"stock" validate:"gte=0"`
    Price       int    `json:"price" validate:"gte=0"`
    MinStock    int    `json:"min_stock" validate:"gte=0"`
}
