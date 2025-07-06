package model

type SalesReport struct {
	TotalTransactions int     `json:"total_transactions"`
	TotalItemsSold    int     `json:"total_items_sold"`
	TotalRevenue      float64 `json:"total_revenue"`
}
