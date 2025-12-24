package middleware

type LoggingMiddleware struct {
	time   string
	method string
	code   int
}

func GetNewLogginMiddleware() *LoggingMiddleware {
	return &LoggingMiddleware{time: "1.1.2025", method: "GET", code: 200}
}
