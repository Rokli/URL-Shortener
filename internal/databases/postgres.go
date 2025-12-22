package databases

type Url struct {
	id        int
	long_url  string
	short_url string
}

func getData() []Url { return []Url{{id: 1, long_url: "https://google.com", short_url: "googl"}} }

func GetUrl(id int) string { return "google.com" }

func CreateUrl(long_url string) error { return nil }

func UpdateUrl(id int) error { return nil }

// func GetDBConnect() sql.DB {
// 	connStr := "user=app password=testpass dbname=shortener sslmode=disable"
// 	db, err := sql.Open("postgres", connStr)

// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	return *db
// }
