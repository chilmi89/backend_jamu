package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

// Instance Bun DB yang bisa digunakan di seluruh aplikasi
var DB *bun.DB

func Connect() {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "3306"
	}
	if name == "" {
		name = "backend_jamu"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, name)

	sqldb, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("❌ Gagal inisialisasi driver database: %v", err)
	}

	DB = bun.NewDB(sqldb, mysqldialect.New())

	// Tambahkan hook untuk logging query di terminal
	DB.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	if err := DB.Ping(); err != nil {
		log.Fatalf("❌ Database tidak terjangkau: %v", err)
	}

	log.Println("✅ Database terhubung (Bun ORM)")
}
