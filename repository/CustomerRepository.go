package repository

import "database/sql"

type Customer struct {
	CustName string
	CustAddr string
}

func (cust *Customer)queryCust(DB *sql.DB)  {



	
}