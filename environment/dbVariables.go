package environment

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/LzKelvinSoares/lzk-go-lib/constants"
	"github.com/LzKelvinSoares/lzk-go-lib/errs"
	"github.com/LzKelvinSoares/lzk-go-lib/logger"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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

func (d DBVariables) GetMongoDBDataSource() string {
	dataSource := fmt.Sprintf("mongodb+srv://%s:%s@%s/",
		d.Username,
		d.Password,
		d.Address)
	return dataSource
}

func (d DBVariables) GetCosmosDBDataSource() string {
	dataSource := fmt.Sprintf(
		"mongodb+srv://%s:%s@%s.documents.azure.com:10255/?ssl=true",
		d.Username,
		d.Password,
		d.Address)
	return dataSource
}

func (d DBVariables) GetDataSource(dbType constants.DBType) (string, *errs.AppError) {
	switch dbType {
	case constants.MySQL:
		return d.GetMySQLDataSource(), nil
	case constants.SQLServer:
		return d.GetSQLServerDataSource(), nil
	case constants.MongoDB:
		return d.GetMongoDBDataSource(), nil
	default:
		logger.Error("DB Type not implemented!")
		return "", errs.NewUnexpectedError("DB Type not implemented!")
	}
}

func GetSQLClient(dbType constants.DBType) *sqlx.DB {
	dbVariables := GetDBVariables()
	dataSource, dbErr := dbVariables.GetDataSource(dbType)
	if dbErr != nil {
		log.Fatal("open: ", dbErr)
		panic(dbErr)
	}

	client, err := sqlx.Open(strings.ToLower(string(dbType)), dataSource)
	if err != nil {
		log.Fatal("open: ", err)
		panic(err)
	}

	client.SetConnMaxLifetime(5 * time.Minute)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(5)
	return client
}

func GetNoSQLClient(dbType constants.DBType) *mongo.Client {
	dbVariables := GetDBVariables()
	dataSource, dbErr := dbVariables.GetDataSource(dbType)
	if dbErr != nil {
		log.Fatal("open: ", dbErr)
		panic(dbErr)
	}
	opts := options.Client().ApplyURI(dataSource)

	client, err := mongo.Connect(opts)
	if err != nil {
		log.Fatal("open: ", err)
		panic(err)
	}

	return client
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
