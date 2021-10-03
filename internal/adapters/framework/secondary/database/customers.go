package database

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

func (dba *Adapter) CreateCustomer(octyCustomerID string, octyProfileID string, shopifyCustomerID string) error {
	if shopifyCustomerID != "" {
		statement, err := dba.conn.Prepare("INSERT INTO Customers (octy_customer_id, octy_profile_id, shopify_customer_id, CreatedAt) VALUES (?, ?, ?, ?)")
		if err != nil {
			return err
		}
		statement.Exec(
			octyCustomerID,
			octyProfileID,
			shopifyCustomerID,
			time.Now().UTC())
	} else {
		statement, err := dba.conn.Prepare("INSERT INTO Customers (octy_customer_id, octy_profile_id, CreatedAt) VALUES (?, ?, ?)")
		if err != nil {
			return err
		}
		_, err = statement.Exec(
			octyCustomerID,
			octyProfileID,
			time.Now().UTC())
		if err != nil {
			return err
		}
	}

	return nil
}

func (dba *Adapter) GetCustomer(octyCustomerID string) (map[string]string, error) {
	query := `SELECT octy_customer_id, octy_profile_id, shopify_customer_id FROM Customers WHERE octy_customer_id =? ORDER BY CreatedAt DESC LIMIT 1;`
	row := dba.conn.QueryRow(query, octyCustomerID)

	var customerID *string
	var octyProfileID *string
	var shopifyCustomerID *string

	switch err := row.Scan(&customerID, &octyProfileID, &shopifyCustomerID); err {
	case sql.ErrNoRows:
		log.Println("No rows were returned!")
		return nil, errors.New("empty result")
	case nil:
		log.Println("Customer found!")
	default:
		return nil, err
	}

	var customer = map[string]string{}
	if shopifyCustomerID != nil {
		customer["octyCustomerID"] = *customerID
		customer["octyProfileID"] = *octyProfileID
		customer["shopifyCustomerID"] = *shopifyCustomerID
	} else {
		customer["octyCustomerID"] = *customerID
		customer["octyProfileID"] = *octyProfileID
		customer["shopifyCustomerID"] = ""
	}

	return customer, nil
}

func (dba *Adapter) UpdateCustomer(octyCustomerID string, octyProfileID string, shopifyCustomerID string) error {
	if shopifyCustomerID != "" {
		statement, err := dba.conn.Prepare("update Customers set octy_customer_id=?, octy_profile_id=?, shopify_customer_id=?, UpdatedAt=? where octy_customer_id=?")
		if err != nil {
			return err
		}
		_, err = statement.Exec(
			octyCustomerID,
			octyProfileID,
			shopifyCustomerID,
			time.Now().UTC(),
			octyCustomerID)
		if err != nil {
			return err
		}
	} else {
		statement, err := dba.conn.Prepare("update Customers set octy_customer_id=?, octy_profile_id=?, UpdatedAt=? where octy_customer_id=?")
		if err != nil {
			return err
		}
		_, err = statement.Exec(
			octyCustomerID,
			octyProfileID,
			time.Now().UTC(),
			octyCustomerID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (dba *Adapter) DeleteCustomer(octyCustomerID string) error {
	statement, err := dba.conn.Prepare("delete from Customers where octy_customer_id=?")
	if err != nil {
		return err
	}
	_, err = statement.Exec(octyCustomerID)
	if err != nil {
		return err
	}
	return nil
}
