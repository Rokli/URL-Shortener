package databases

import (
	"database/sql"
)

type postgresUrlRepository struct {
	db *sql.DB
}

// func getData() []domain.Url {
// 	return []domain.Url{domain.CreateNewUrl()}
// }

// func GetUrl(id int) string { return "google.com" }

// func CreateUrl(long_url string) error { return nil }

// func UpdateUrl(id int) error { return nil }

// func GetDBConnect() sql.DB {
// 	connStr := "user=app password=testpass dbname=shortener sslmode=disable"
// 	db, err := sql.Open("postgres", connStr)

// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	return *db
// }
