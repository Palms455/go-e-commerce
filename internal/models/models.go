package models

import (
	"context"
	"database/sql"
	"time"
)

// DBModel type for database connection
type DBModel struct {
	DB *sql.DB
}

// Models is the wrapper fo all models
type Models struct {
	DB DBModel
}

// NewModels returns a model with db connection
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB:db},
	}
}

// Widget type for al widgets
type Widget struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	InventoryLevel int `json:"inventory_level"`
	Price int `json:"price"`
	CreatedAt time.Time `json:"-"`
	UpdateAt time.Time `json:"-"`
}


func (m *DBModel) GetWidget(id int) (Widget, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	var widget Widget

	row := m.DB.QueryRowContext(ctx, "select w.id, w.name from products.widgets w where id = $1", id)
	err := row.Scan(&widget.ID, &widget.Name)
	if err != nil {
		return widget, err
	}
	return widget, nil
}