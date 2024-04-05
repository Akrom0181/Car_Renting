package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"rent-car/api/models"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/google/uuid"
)

type customerRepo struct {
	db *pgxpool.Pool
}

func NewCustomer(db *pgxpool.Pool) customerRepo {
	return customerRepo{
		db: db,
	}
}

func (c *customerRepo) Create(ctx context.Context, customer models.Customer) (string, error) {
	id := uuid.New()

	query := `insert into customers(id, first_name, last_name, gmail, phone, is_blocked, login, password) values($1,$2,$3,$4,$5,$6,$7,$8)`

	_, err := c.db.Exec(context.Background(), query,
		id.String(),
		customer.FirstName,
		customer.LastName,
		customer.Gmail,
		customer.Phone,
		customer.Is_Blocked,
		customer.Login,
		customer.Password)
	if err != nil {
		return "error:", err
	}
	return id.String(), nil
}

func (c *customerRepo) ChangePassword(ctx context.Context, customer models.Customer) (string, error) {
	query := `UPDATE customers SET password = $1 where phone = $2`
	_, err := c.db.Exec(context.Background(), query, 
		customer.Password, 
		customer.Phone)
	if err != nil {
		return "", err
	}
	return "Successfully changed", nil
}

func (c *customerRepo) GetPassword(ctx context.Context, customer models.GetPassword) (string, error) {
	query := `SELECT password from customers WHERE id=$1`
	_, err := c.db.Exec(context.Background(), query, customer.Password)
	if err != nil {
		return "", err
	}
	return customer.Password, nil
}

func (c *customerRepo) Update(ctx context.Context, customer models.Customer) (string, error) {
	query := `update customers set 
	first_name=$1,
	last_name=$2,
	gmail=$3,
	phone=$4,
	is_blocked=$5,
	updated_at=CURRENT_TIMESTAMP
	WHERE id = $6 AND deleted_at=0
	`
	_, err := c.db.Exec(context.Background(), query,
		customer.FirstName,
		customer.LastName,
		customer.Gmail,
		customer.Phone,
		customer.Is_Blocked,
		customer.Id)
	if err != nil {
		return "", err
	}
	return customer.Id, nil
}

func (c *customerRepo) GetAll(ctx context.Context, req models.GetAllCustomersRequest) (models.GetAllCustomersResponse, error) {
	var (
		resp   = models.GetAllCustomersResponse{}
		filter = ""
	)

	offset := (req.Page - 1) * req.Limit
	if req.Search != "" {
		filter += fmt.Sprintf(` and first_name ILIKE '%%%v%%' `, req.Search)
	}

	filter += fmt.Sprintf("OFFSET %v LIMIT %v", offset, req.Limit)
	fmt.Println("filter:", filter)
	rows, err := c.db.Query(context.Background(), `select count(id) over(),
			id,
			first_name,
			last_name,	
			gmail,
			phone,
			is_blocked,
			created_at::date,
			updated_at FROM customers WHERE deleted_at = 0 `+filter+``)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			customer   = models.Customer{}
			Updated_at sql.NullString
			CreatedAt sql.NullString
		)
		if err := rows.Scan(
			&resp.Count,
			&customer.Id,
			&customer.FirstName,
			&customer.LastName,
			&customer.Gmail,
			&customer.Phone,
			&customer.Is_Blocked,
			&CreatedAt,
			&Updated_at,
		); err != nil {
			return resp, err
		}
		resp.Customers = append(resp.Customers, customer)
	}
	return resp, nil
}
func (c *customerRepo) GetByID(ctx context.Context, id string) (models.Customer, error) {
	customer := models.Customer{}

	var (
		first_name sql.NullString
		last_name  sql.NullString
		gmail      sql.NullString
		phone      sql.NullString
		is_blocked sql.NullBool
		login      sql.NullString
		created_at sql.NullString
		updated_at sql.NullString
	)

	if err := c.db.QueryRow(context.Background(), `select id, first_name, last_name, gmail, phone, is_blocked, login, created_at, updated_at from customers where id = $1`, id).Scan(
		&customer.Id,
		&first_name,
		&last_name,
		&gmail,
		&phone,
		&is_blocked,
		&login,
		&created_at,
		&updated_at); err != nil {
		return models.Customer{}, err
	}
	return models.Customer{
		Id:         customer.Id,
		FirstName:  first_name.String,
		LastName:   last_name.String,
		Gmail:      gmail.String,
		Phone:      phone.String,
		Is_Blocked: is_blocked.Bool,
		Login:      login.String,
		CreatedAt:  created_at.String,
		UpdatedAt:  updated_at.String,
	}, nil
}

func (c *customerRepo) Delete(ctx context.Context, id string) error {
	queary := `delete from customers where id = $1`
	_, err := c.db.Exec(context.Background(), queary, id)
	if err != nil {
		return err
	}
	return nil
}

func (c *customerRepo) GetCustomerCars(ctx context.Context, req models.GetAllCustomersRequest) ([]models.GetOrder, error) {
	var orders []models.GetOrder

	offset := (req.Page - 1) * req.Limit

	query := `SELECT DISTINCT
                o.id,
				c.id AS car_id,
                c.name AS car_name,
                c.brand AS car_brand,
				c.model as car_model,
                cu.id AS customer_id,
                cu.first_name AS customer_first_name,
                o.from_date,
                o.to_date,
                o.amount,
                o.created_at,
                o.updated_at 
            FROM 
                orders o
            JOIN 
                cars c ON o.car_id = c.id
            JOIN 
                customers cu ON o.customer_id = cu.id
			WHERE 1=1`

	if req.Search != "" {
		query += fmt.Sprintf(`AND c.name ILIKE '%%%s%%'`, req.Search)
	}
	query += fmt.Sprintf("ORDER BY o.created_at OFFSET %d LIMIT %d", offset, req.Limit)

	rows, err := c.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order models.GetOrder
		err := rows.Scan(
			&order.Id,
			&order.Car.Id,
			&order.Car.Name,
			&order.Car.Brand,
			&order.Car.Model,
			&order.Customer.Id,
			&order.Customer.FirstName,
			&order.FromDate,
			&order.ToDate,
			&order.Amount,
			&order.CreatedAt,
			&order.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}

func (o *customerRepo) GetByIDCustomerCar(ctx context.Context, id string) (models.GetOrder, error) {
	order := models.GetOrder{
		Car:      models.Car{},
		Customer: models.Customer{},
	}

	if err := o.db.QueryRow(context.Background(), `SELECT 
            o.id,
            c.id AS car_id,
            c.name AS car_name,
            c.brand AS car_brand,
            c.model as car_model,
            cu.id AS customer_id,
            cu.first_name AS customer_first_name,
            cu.gmail AS customer_gmail,
            o.from_date,
            o.to_date,
            o.status,
            o.amount,
            o.created_at,
            o.updated_at 
            FROM 
                orders o
            JOIN 
                cars c ON o.car_id = c.id
            JOIN 
                customers cu ON o.customer_id = cu.id
            WHERE cu.id=$1`, id).Scan(
		&order.Id,
		&order.Car.Id,
		&order.Car.Name,
		&order.Car.Brand,
		&order.Car.Model,
		&order.Customer.Id,
		&order.Customer.FirstName,
		&order.Customer.Gmail,
		&order.FromDate,
		&order.ToDate,
		&order.Status,
		&order.Amount,
		&order.CreatedAt,
		&order.UpdatedAt,
	); err != nil {
		return models.GetOrder{}, err
	}
	return order, nil
}

// func (o *customerRepo) GetCustomerCarss(id string) (models.GetOrder, error) {
//     order := models.GetOrder{
//         Car:      models.Car{},
//         Customer: models.Customer{},
//     }

//     if err := o.db.QueryRow(`SELECT DISTINCT
//             o.id,
//             c.name AS car_name,
//             cu.first_name AS customer_first_name,
//             o.from_date,
//             o.to_date,
//             o.created_at
//             FROM
//                 orders o
//             JOIN
//                 cars c ON o.car_id = c.id
//             JOIN
//                 customers cu ON o.customer_id = cu.id
//             WHERE cu.id=$1`, id).Scan(

//         &order.Id,
//         &order.Car.Name,
//         &order.Customer.Id,
//         &order.Customer.FirstName,
//         &order.FromDate,
//         &order.ToDate,
//         &order.CreatedAt,
//     ); err != nil {
//         return models.GetOrder{}, err
//     }
//     return order, nil
// }
