package model

import "time"

type Sale struct {
    ID        int64     `json:"id"`
    UserID    int64     `json:"user_id"` // user yang buat transaksi
    Total     int64     `json:"total"`
    CreatedAt time.Time `json:"created_at"`
}

type SaleItem struct {
    ID       int64 `json:"id"`
    SaleID   int64 `json:"sale_id"`
    ItemID   int64 `json:"item_id"`
    Quantity int64 `json:"quantity"`
    Price    int64 `json:"price"` // harga per unit saat transaksi
}
