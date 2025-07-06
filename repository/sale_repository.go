package repository

import (
	"database/sql"
	"project-app-inventory-restapi-golang-faisal/model"
)

type SaleRepository interface {
	CreateSale(sale *model.Sale, items []model.SaleItem) error
	GetAll() ([]model.Sale, error)
	GetByID(id int64) (*model.Sale, error)
	GetSalesReport() (*model.SalesReport, error)
}

type saleRepository struct {
	db *sql.DB
}

func NewSaleRepository(db *sql.DB) SaleRepository {
	return &saleRepository{db: db}
}

func (r *saleRepository) CreateSale(sale *model.Sale, items []model.SaleItem) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func ()  {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = tx.QueryRow(`INSERT INTO sales(user_id, total) VALUES($1, $2) RETURNING id`, sale.UserID, sale.Total).Scan(&sale.ID)
	if err != nil {
		return err
	}

	for i, _ := range items {
		items[i].SaleID = sale.ID

		_, err := tx.Exec(`INSERT INTO sale_items(sale_id, item_id, quantity, price) VALUES($1, $2, $3, $4)`, items[i].SaleID, items[i].ItemID, items[i].Quantity, items[i].Price)
		if err != nil {
			return err
		}

		//kurangi stock barang
		res, err := tx.Exec(`UPDATE items SET stock = stock - $1 WHERE id = $2 AND stock >= $1`, items[i].Quantity, items[i].ItemID)
		if err != nil {
			return err
		}
		rowsAffected, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return sql.ErrNoRows
		}
	}
	return  nil
}

func (r *saleRepository) GetAll()([]model.Sale, error) {
	rows, err := r.db.Query(`SELECT id, user_id, total, crated_at FROM sales ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sales []model.Sale
	for rows.Next(){
		var s model.Sale
		if err := rows.Scan(&s.ID, &s.UserID, &s.Total, &s.CreatedAt); err != nil {
			return nil, err
		}
		sales = append(sales, s)
	}
	return sales, nil
}

func (r *saleRepository) GetByID(id int64) (*model.Sale, error) {
	var s model.Sale
	err := r.db.QueryRow(`SELECT id, user_id, total, created_at FROM sales WHERE id = $1`, id).Scan(&s.ID, &s.UserID, &s.Total, &s.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *saleRepository) GetSalesReport() (*model.SalesReport, error) {
	query := `
		SELECT 
			COUNT(DISTINCT s.id) total_transactions,
			COALESCE(SUM(si.quantity), 0) AS total_items_sold,
			COALESCE(SUM(si.quantity * i.price), 0) AS total_revenue
		FROM sales s
		JOIN sale_items si ON s.id = si.sale_id
		JOIN items i ON si.item_id = i.id;
		`
	var report model.SalesReport
	err := r.db.QueryRow(query).Scan(
		&report.TotalTransactions,
		&report.TotalItemsSold,
		&report.TotalRevenue,
	)
	if err != nil {
		return nil, err
	}
	return &report, nil
}