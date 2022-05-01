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
		DB: DBModel{DB: db},
	}
}

// Widget type for al widgets
type Widget struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	InventoryLevel int       `json:"inventory_level"`
	Price          int       `json:"price"`
	Image          string    `json:"image"`
	CreatedAt      time.Time `json:"-"`
	UpdateAt       time.Time `json:"-"`
}

// Order type for all orders
type Order struct {
	ID            int       `json:"id"`
	WidgetId      int       `json:"widget_id"`
	TransactionId int       `json:"transaction_id"`
	StatusId      int       `json:"status_id"`
	Quantity      int       `json:"quantity"`
	Amount        int       `json:"amount"`
	CreatedAt     time.Time `json:"-"`
	UpdateAt      time.Time `json:"-"`
}

// Status type for all order statuses
type Status struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"-"`
	UpdateAt  time.Time `json:"-"`
}

// TransactionStatus type for all transaction statuses
type TransactionStatus struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"-"`
	UpdateAt  time.Time `json:"-"`
}

// Transaction type for all transactions
type Transaction struct {
	ID                  int       `json:"id"`
	Amount              int       `json:"amount"`
	Currency            string    `json:"currency"`
	LastFour            string    `json:"last_four"`
	BankReturnCode      int       `json:"bank_return_code"`
	TransactionStatusId int       `json:"transaction_status_id"`
	CreatedAt           time.Time `json:"-"`
	UpdateAt            time.Time `json:"-"`
}

// User type for all users
type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"-"`
	UpdateAt  time.Time `json:"-"`
}

// GetWidget gets one widget by id
func (m *DBModel) GetWidget(id int) (Widget, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var widget Widget

	row := m.DB.QueryRowContext(ctx, "select w.id, w.name from products.widgets w where id = $1", id)
	err := row.Scan(&widget.ID, &widget.Name)
	if err != nil {
		return widget, err
	}
	return widget, nil
}
