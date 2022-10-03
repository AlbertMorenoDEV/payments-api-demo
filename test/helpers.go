package test

import (
	"database/sql"
	"fmt"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/infrastructure"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/ui"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"time"
)

var (
	isInitialized bool
	config        *infrastructure.Config
	services      *infrastructure.Services
	router        http.Handler
)

func setUp() {
	if isInitialized {
		return
	}

	domainEvents := make(chan domain.DomainEvent)
	config = infrastructure.NewConfig()
	services = infrastructure.BuildServices(config, domainEvents)
	router = ui.BuildRouter(services)

	dbSetUp()

	isInitialized = true

	log.Println("Testing env successfully initialized")
}

func dbSetUp() {

	sqlAddress := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		config.MysqlUser(),
		config.MysqlPass(),
		config.MysqlHost(),
		config.MysqlPort(),
		config.MysqlDb(),
	)

	db, err := sql.Open("mysql", sqlAddress)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if _, err = db.Exec(`CREATE TABLE IF NOT EXISTS users_balance (
	  user_id varchar(36) NOT NULL,
	  amount int NOT NULL,
	  currency char(3) NOT NULL,
	  PRIMARY KEY (user_id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`); err != nil {
		panic(err)
	}

	if _, err = db.Exec(`CREATE TABLE IF NOT EXISTS transactions (
	  transaction_id varchar(36) NOT NULL,
	  user_id varchar(36) NOT NULL,
	  destination_user_id varchar(36) NOT NULL,
	  amount decimal(10,2) NOT NULL,
	  currency char(3) NOT NULL,
	  creation_date datetime(6) NOT NULL,
	  PRIMARY KEY (transaction_id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`); err != nil {
		panic(err)
	}

	log.Println("DB tables successfully created")
}

func executeRequest(method, path string, body io.Reader) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, body)
	r := httptest.NewRecorder()

	router.ServeHTTP(r, req)

	return r
}

type FixedTimeProvider struct {
	date time.Time
}

func NewFixedDateTimeProvider() *FixedTimeProvider {
	return &FixedTimeProvider{}
}

func (dp *FixedTimeProvider) Now() time.Time {
	if dp.date.IsZero() {
		dp.date = time.Now()
	}

	return dp.date
}
