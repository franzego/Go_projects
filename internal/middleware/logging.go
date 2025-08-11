package middleware


import (
    "log"
    "time"
    "net/http"
)

//func LogMiddleware(next http.Handler) http.Handler {
  //  return http.Handler(func(w http.ResponseWriter, r *http.Request))

//}

// struct type to wrap around the responsewriter because w cannot send back status number

type statusrecorder struct {
    http.ResponseWriter
    status int
}

// This method stores the response code for logging later

func (r *statusrecorder) WriteHeader(code int) {
    r.status = code // stroing the status in <code> to be used later
    r.ResponseWriter.WriteHeader(code) // still return the response message to the client
}

// The midddleware function that does the logging

func LogMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        start := time.Now() // stores the start time
        

        rec := &statusrecorder{ResponseWriter: w, status: http.StatusOK}

        next.ServeHTTP(rec, r)

        duration := time.Since(start) // stores the duration of it all
        log.Printf("%s %s %d %s", r.Method, r.URL.Path, rec.status, duration)
    })
}
