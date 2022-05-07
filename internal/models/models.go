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
	WidgetID      int       `json:"widget_id"`
	TransactionID int       `json:"transaction_id"`
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
	TransactionStatusID int       `json:"transaction_status_id"`
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

	row := m.DB.QueryRowContext(ctx, `
		select 
		    w.id, w.name, w.description, w.inventory_level, 
		    w.price, coalesce(w.image, ''), w.created_at, w.updated_at
		from products.widgets w where id = $1`, id)
	err := row.Scan(
		&widget.ID,
		&widget.Name,
		&widget.Description,
		&widget.InventoryLevel,
		&widget.Price,
		&widget.Image,
		&widget.CreatedAt,
		&widget.UpdateAt,
	)
	if err != nil {
		return widget, err
	}
	return widget, nil
}

// InsertTransaction insert transaction and returns id
func (m *DBModel) InsertTransaction(txn Transaction) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	SQL := `
		insert into products.transactions (
			amount, currency, last_four, bank_return_code, transaction_status_id, create_at, updated_at
		)
		values ($1, $2, $3, $4, $5, $6, $7)
	`
	result, err := m.DB.ExecContext(ctx, SQL,
		txn.Amount,
		txn.Currency,
		txn.LastFour,
		txn.BankReturnCode,
		txn.TransactionStatusID,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}


// InsertOrder insert order and returns id
func (m *DBModel) InsertOrder(order Order) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	SQL := `
		insert into products.orders (
			widget_id, transaction_id, status_id, quantity, amount, created_at, updated_at
		)
		values ($1, $2, $3, $4, $5, $6, $7)
	`
	result, err := m.DB.ExecContext(ctx, SQL,
		order.WidgetID,
		order.TransactionID,
		order.StatusId,
		order.Quantity,
		order.Amount,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}