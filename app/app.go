package app

import (
	"fmt"
	"github.com/fuadaghazada/banking/domain"
	"github.com/fuadaghazada/banking/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
)

func sanityCheck() {
	if os.Getenv("SERVER_HOST") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variables not found...")
	}
}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPwd := os.Getenv("DB_PWD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPwd, dbHost, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	return client
}

func Start() {
	sanityCheck()

	router := mux.NewRouter()

	dbClient := getDbClient()

	// Wiring
	//ch := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)

	ch := CustomerHandler{service: service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service: service.NewAccountService(accountRepositoryDb)}

	// Routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerId:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerId:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)

	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")
	serverAddr := fmt.Sprintf("%s:%s", host, port)
	err := http.ListenAndServe(serverAddr, router)
	if err != nil {
		return
	}
}
