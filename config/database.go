package config

import (
	"erajaya/internal/model"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDB menginisialisasi database
func NewDB(log *zap.Logger) *gorm.DB {
	// Ambil konfigurasi dari env
	dbUser := GetEnv("DB_USER", "default_user")
	dbPassword := GetEnv("DB_PASSWORD", "default_pass")
	dbName := GetEnv("DB_NAME", "default_db")
	dbHost := GetEnv("DB_HOST", "localhost")
	dbPort := GetEnv("DB_PORT", "5432")
	dbSSLMode := GetEnv("DB_SSLMODE", "disable")
	dbDriver := GetEnv("DB_DRIVER", "postgres")

	var (
		db  *gorm.DB
		err error
	)
	if dbDriver == "postgres" {
		// Postgres
		dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
			dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: NewDatabaseLogger(log),
		})
	} else if dbDriver == "mysql" {
		// MySQL
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local",
			dbUser, dbPassword, dbHost, dbPort, dbName)

		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: NewDatabaseLogger(log),
		})
	} else {
		log.Fatal("DB Driver " + dbDriver + " not supported")
	}

	if err != nil {
		log.Fatal("Failed to connect to database", zap.Error(err))
	}

	// AutoMigrate
	err = db.AutoMigrate(
		&model.Product{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database", zap.Error(err))
	}

	// Optional. Penambahan index untuk product
	db.Exec("CREATE INDEX IF NOT EXISTS idx_products_price_asc ON products(price ASC)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_products_price_desc ON products(price DESC)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_products_name_asc ON products(name ASC)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_products_name_desc ON products(name DESC)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_products_created_at_desc ON products(created_at DESC)")

	log.Info("Database connected and migrated")
	return db
}
