package database

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/BurhaanAshraf/finance-api/internal/config"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func Connect(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&tls=tidb",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)
	rootCertPool := x509.NewCertPool()

	pem, err := os.ReadFile("certs/isrgrootx1.pem")
	if err != nil {
		return nil, err
	}

	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		return nil, fmt.Errorf("failed to append CA certificate")
	}

	err = mysql.RegisterTLSConfig("tidb", &tls.Config{
		RootCAs: rootCertPool,
	})
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to MySQL successfully")

	return db, nil
}
