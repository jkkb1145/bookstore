package repository

import (
	"database/sql"
	"demo02/global"
)

type OrderDAO struct {
	OrderDB *sql.DB
}

func NewOrderDAO() *OrderDAO {
	return &OrderDAO{
		OrderDB: global.GetDb(),
	}
}
