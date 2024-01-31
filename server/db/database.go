package db

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func CreateDatabase() (*gorm.DB, error) {
    godotenv.Load()

    host := os.Getenv("DB_HOST")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    sslmode := os.Getenv("DB_SSLMODE")

    dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=" + sslmode
    
    var err error
    for i := 0; i < 5; i++ {
        DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err == nil {
            break
        }
        log.Printf("Retry %d/5...", i+1)
        time.Sleep(5 * time.Second)
    }

    if err != nil {
        log.Fatalf("Failed to connect to database after retries: %v", err)
    }

    return DB, nil
}
func Init() error{
	var err error
	DB, err = CreateDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	
	return err
}

func GetDB() *gorm.DB {
    return DB
}

