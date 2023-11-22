package middleware

import (
    "os"
    "log"
    "time"
    "net/http"
)

func NewLogger(l *log.Logger) middleware {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
            start := time.Now()

            defer func() {
                if r := recover(); r != nil {
                    l.Println(r)
                }
            }()

            next.ServeHTTP(w, req)

            hostname := os.Getenv("HOSTNAME")

            log.Printf(
                "%s %s %s %s %v",
                hostname,
                req.Method,
                req.RequestURI,
                req.Proto,
                time.Since(start),
            )
        })
    }
}

