package lib

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Create single DB instance
func CreateDatabase() (*gorm.DB, error) {
    host := os.Getenv("DB_HOST")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    sslmode := os.Getenv("DB_SSLMODE")

    log.Println("Connecting to database...")
    dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " pgbouncer=true" + " sslmode=" + sslmode + " default_query_exec_mode=cache_describe"
    
    var err error
    for i := 0; i < 5; i++ {
        DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{

        })
        if err == nil {
            break
        }
        log.Printf("Retry %d/5...", i+1)
        time.Sleep(5 * time.Second)
    }

    if err != nil {
        log.Fatalf("Failed to connect to database after retries: %v", err)
    }

    AutoMigrate(DB)

    return DB, nil
}
// initialize the database
func InitDB() error{
	var err error
	DB, err = CreateDatabase()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
	}
	
	return err
}

func GetDB() *gorm.DB {
    return DB
}

func CloseDB() error {
    db, err := DB.DB()
    if err != nil {
        return err
    }
    return db.Close()
}
