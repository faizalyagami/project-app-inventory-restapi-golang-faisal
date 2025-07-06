package repository

import (
	"database/sql"
	"project-app-inventory-restapi-golang-faisal/model"
)


type ItemRepository interface {
  GetAll() ([]model.Item, error)
  GetByID(id int64) (*model.Item, error)
  Create(item *model.Item) error
  Update(item *model.Item) error
  Delete(id int64) error
  GetLowStockItems(threshold int64) ([]model.Item, error)
}

type itemRepository struct {
  db *sql.DB
}

func NewItemRepository(db *sql.DB) ItemRepository {
	return &itemRepository{db: db}
}

func (r *itemRepository) GetAll() ([]model.Item, error) {
	rows, err := r.db.Query(`SELECT * FROM items ORDER BY id`)
	if err != nil {
		return  nil, err
	}
	defer rows.Close()

	var items []model.Item
	for rows.Next(){
		var i model.Item
		if err := rows.Scan(&i.ID, &i.Name,&i.CategoryID,&i.RackID, &i.WarehouseID,&i.Stock, &i.Price, &i.MinStock, new(string), new(string)); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	return  items, nil
}

func (r *itemRepository) GetByID(id int64) (*model.Item, error) {
	row := r.db.QueryRow(`SELECT * FROM items WHERE id=$1`, id)

	var i model.Item
	if err := row.Scan(&i.ID, &i.Name,&i.CategoryID,&i.RackID, &i.WarehouseID,&i.Stock, &i.Price, &i.MinStock, new(string), new(string)); err != nil {
		return  nil, err
	}
	return &i, nil
}

func (r *itemRepository) Create(item *model.Item) error {
	_, err := r.db.Exec(`INSERT INTO items (name, category_id, rack_id, warehouse_id, stock, price, min_stock)VALUES ($1,$2,$3,$4,$5,$6,$7)`,
		item.Name, item.CategoryID, item.RackID, item.WarehouseID, item.Stock, item.Price, item.MinStock)
	return err
}

func (r *itemRepository) Update(item *model.Item) error {
	_, err := r.db.Exec(`UPDATE items SET name=$1, category_id=$2, rack_id=$3, warehouse_id=$4, stock=$5, price=$6, min_stock=$7`,
		item.Name, item.CategoryID, item.RackID, item.WarehouseID, item.Stock, item.Price, item.MinStock)
	return err 
}

func (r *itemRepository) Delete(id int64) error {
	_, err := r.db.Exec(`DELETE FROM items WHERE id=$1`, id)
	return err
}

func (r *itemRepository) GetLowStockItems(thresold int64) ([]model.Item, error) {
	rows, err := r.db.Query(`SELECT id, name, category_id, rack_id, warehouse_id, stock, price FROM items WHERE stock < $1`, thresold)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.Item
	for rows.Next(){
		var item model.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.CategoryID, &item.RackID, &item.WarehouseID, &item.Stock, &item.Price); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}