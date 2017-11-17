package routes

import (
	"net/http"

	"github.com/joshadambell/brickandclick/routes/handlers"
	goji "goji.io"
	"goji.io/pat"
)

func Register(mux *goji.Mux) {
	mux.Use(corsHandler)
	mux.HandleFunc(pat.Get("/search/:searchText"), handlers.Search)
}

func corsHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Methods", "*")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "x-pingpong")
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
