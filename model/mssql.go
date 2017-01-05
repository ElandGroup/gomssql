package model

import (
	"database/sql"

	_ "github.com/mattn/go-adodb"
)

type Mssql struct {
	*sql.DB
	Conn string
}

func (m *Mssql) Open() (err error) {
	if m.DB, err = sql.Open("adodb", "Provider=SQLOLEDB;Data Source=qds114812583.my3w.com,1433;Initial Catalog=qds114812583_db;User ID=qds114812583;password=Eland123;"); err != nil {
		return err
	}
	return nil
}
