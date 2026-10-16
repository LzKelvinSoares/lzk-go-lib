package environment

import (
	"fmt"
	"os"
)

type DBVariables struct {
	Username string
	Password string
	Address  string
	Port     string
	Name     string
}

func (d DBVariables) GetMySQLDataSource() string {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		d.Username,
		d.Password,
		d.Address,
		d.Port,
		d.Name)
	return dataSource
}

func (d DBVariables) GetSQLServerDataSource() string {
	dataSource := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;encrypt=true",
		d.Address,
		d.Username,
		d.Password,
		d.Port,
		d.Name)
	return dataSource
}

func GetDBVariables() DBVariables {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbAddress := os.Getenv("DB_ADDRESS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	return DBVariables{
		Username: dbUser,
		Password: dbPassword,
		Address:  dbAddress,
		Port:     dbPort,
		Name:     dbName,
	}
}
