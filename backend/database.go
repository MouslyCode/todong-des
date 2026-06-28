package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	rootCertPool := x509.NewCertPool()
	pem, errs := os.ReadFile("./global-bundle.pem")
	if errs != nil {
		log.Fatalf("Failed to read cert: %v", errs)
	}
	rootCertPool.AppendCertsFromPEM(pem)
	mysql.RegisterTLSConfig("rds", &tls.Config{
		RootCAs: rootCertPool,
	})

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&tls=rds",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var err error
	for i := 0; i < 10; i++ {
		DB, err = gorm.Open(gormmysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Database not ready, retrying in 3s... (%d/10)", i+1)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	DB.AutoMigrate(&Todo{})
	log.Println("Database connected!")
}
