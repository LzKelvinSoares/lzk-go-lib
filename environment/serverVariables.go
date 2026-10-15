package environment

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ServerVariables struct {
	Address string
	Port    string
}

func (s ServerVariables) GetAddressToServe() string {
	return fmt.Sprintf("%s:%s", s.Address, s.Port)
}

func GetServerVariables() ServerVariables {
	sanityCheck()
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	return ServerVariables{
		Address: address,
		Port:    port,
	}
}

func sanityCheck() {
	godotenv.Load()
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variables not defined!")
	}
}
