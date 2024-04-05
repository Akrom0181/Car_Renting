// package postgres

// import (
// 	"database/sql"
// 	"fmt"
// 	"rent-car/api/models"
// )

// type cusRepo struct {
// 	db *sql.DB
// }

// func NewCus(db *sql.DB) cusRepo {
// 	return cusRepo{
// 		db: db,
// 	}
// }

// func (c *cusRepo) GetCusCars(id string, req models.GetAllCarsRequest) ([]models.GetOrder, error) {
// 	var orders []models.GetOrder

// 	offset := (req.Page - 1) * req.Limit

// 	query := `SELECT 
// 			o.id,
// 			c.name AS car_name,
// 			cu.first_name AS customer_first_name,
// 			o.from_date,
// 			o.to_date,
// 			o.created_at
// 			FROM 
// 				orders o
// 			JOIN 
// 				cars c ON o.car_id = c.id
// 			JOIN 
// 				customers cu ON o.customer_id = cu.id
// 			WHERE cu.id = $1`

// 	if req.Search != "" {
// 		query += fmt.Sprintf(` AND c.name ILIKE '%%%s%%'`, req.Search)
// 	}
// 	query += fmt.Sprintf(" ORDER BY o.created_at OFFSET %d LIMIT %d", offset, req.Limit)

// 	rows, err := c.db.Query(query, id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var order models.GetOrder
// 		err := rows.Scan(
// 			&order.Id,
// 			&order.Car.Name,
// 			&order.Customer.FirstName,
// 			&order.FromDate,
// 			&order.ToDate,
// 			&order.CreatedAt,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}
// 		orders = append(orders, order)
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return orders, nil
// }

// func (c *customerRepo) GetByIDC(id string) (models.Customer, error) {
// 	customer := models.Customer{}

// 	if err := c.db.QueryRow(`SELECT id, 
// 	    first_name,
// 	    last_name,
// 	    gmail,
// 		phone
// 		FROM customers WHERE id=$1`, id).Scan(
// 			&customer.Id,
// 			&customer.FirstName,
// 			&customer.LastName,
// 			&customer.Gmail,
// 			&customer.Phone); err != nil {
// 				return models.Customer{}, err
// 			}
// 	if err := c.db.QueryRow(`SELECT COUNT(*), COUNT(DISTINCT car_id) FROM order WHERE customer_id=$1`,id).Scan(
// 		&customer.OrdersCount,
// 		&customer.CarsCount); err != nil {
// 			return models.Customer{}, err
// 		}	
// 	return customer, nil	
// }




// // SELECT Category, MIN(ProductID) AS FirstProductID, ProductName
// // FROM Products
// // GROUP BY Category;


// // -- Category | FirstProductID | ProductName
// // -- ---------|----------------|-------------
// // -- Clothing | 1              | T-Shirt
// // -- Kitchen  | 4              | Coffee Mug
package postgres