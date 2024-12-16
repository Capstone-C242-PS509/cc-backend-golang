package postgresql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db       *gorm.DB
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname   = os.Getenv("DB_NAME")
)

// connectWithConnector creates a database connection using Cloud SQL Connector.
// func connectWithConnector() (*sql.DB, error) {
// 	dsn := fmt.Sprintf("user=%s password=%s database=%s", user, password, dbname)
// 	config, err := pgx.ParseConfig(dsn)
// 	if err != nil {
// 		return nil, fmt.Errorf("pgx.ParseConfig: %w", err)
// 	}

// 	var opts []cloudsqlconn.Option
// 	if usePrivate != "" {
// 		opts = append(opts, cloudsqlconn.WithDefaultDialOptions(cloudsqlconn.WithPrivateIP()))
// 	}

// 	dialer, err := cloudsqlconn.NewDialer(context.Background(), opts...)
// 	if err != nil {
// 		return nil, fmt.Errorf("cloudsqlconn.NewDialer: %w", err)
// 	}

// 	config.DialFunc = func(ctx context.Context, network, instance string) (net.Conn, error) {
// 		return dialer.Dial(ctx, instanceConnectionName)
// 	}

// 	dbURI := stdlib.RegisterConnConfig(config)
// 	sqlDB, err := sql.Open("pgx", dbURI)
// 	if err != nil {
// 		return nil, fmt.Errorf("sql.Open: %w", err)
// 	}

// 	return sqlDB, nil
// }

// connectWithPublicIP connects to the database using Public IP.
func connectWithPublicIP() (*sql.DB, error) {
	// Gunakan Public IP dari instance Cloud SQL Anda
	publicIP := os.Getenv("DB_HOST") // Pastikan Public IP Cloud SQL diatur di DB_HOST

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s",
		publicIP,
		user,
		password,
		dbname,
	)

	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	return sqlDB, nil
}

// GetDBInstance returns the initialized Gorm DB instance.
func GetDBInstance() *gorm.DB {
	return db
}

func init() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "json/original-advice-438105-i6-9ed330e0dc52.json")

	// Menggunakan Public IP
	sqlDB, err := connectWithPublicIP()
	if err != nil {
		log.Fatalf("Failed to connect to Cloud SQL: %v", err)
	}

	db, err = gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to initialize Gorm: %v", err)
	}

	log.Println("Connected to Cloud SQL database!")
}
