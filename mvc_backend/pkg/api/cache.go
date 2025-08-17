package api

import (
	"net/http"
	"time"
	"bytes"
	"github.com/patrickmn/go-cache"
	"github.com/gorilla/mux"
	"strings"
)

var c = cache.New(3*time.Hour, 10*time.Minute)

type responseWriter struct {
    http.ResponseWriter
    buf        bytes.Buffer
    statusCode int
}

func (rw *responseWriter) WriteHeader(statusCode int) {
    rw.statusCode = statusCode
    rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *responseWriter) Write(data []byte) (int, error) {
    rw.buf.Write(data)
    return rw.ResponseWriter.Write(data)
}

func CacheMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet && r.Method != http.MethodHead {
            next(w, r)
            return
        }

		var cacheKey string

        if strings.HasPrefix(r.URL.Path, "/menu/cat/") {
            vars := mux.Vars(r)
            id := vars["id"]

            cacheKey = "/menu/cat/" + id
        } else {
            cacheKey = r.URL.String()
        }

		cachedResponse, found := c.Get(cacheKey)
		if found {
			w.Header().Set("Content-Type", "application/json")
            w.Write(cachedResponse.([]byte))
			return
		}

		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next(rw, r)

		if rw.statusCode >= 200 && rw.statusCode < 300 {
            c.Set(cacheKey, rw.buf.Bytes(), cache.DefaultExpiration)

		}
	}
}
