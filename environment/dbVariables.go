package environment

import (
	"fmt"
	"os"

	"github.com/LzKelvinSoares/lzk-go-lib/constants"
	"github.com/LzKelvinSoares/lzk-go-lib/errs"
	"github.com/LzKelvinSoares/lzk-go-lib/logger"
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

func (d DBVariables) GetDataSource(dbType constants.DBType) (string, *errs.AppError) {
	switch dbType {
	case constants.MySQL:
		return d.GetMySQLDataSource(), nil
	case constants.SQLServer:
		return d.GetSQLServerDataSource(), nil
	default:
		logger.Error("DB Type not implemented!")
		return "", errs.NewUnexpectedError("DB Type not implemented!")
	}
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
