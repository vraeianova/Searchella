package main

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/api/search/{search_phrase}/", func(w http.ResponseWriter, r *http.Request) {

		search_phrase := chi.URLParam(r, "search_phrase")

		query := `{
			"search_type": "match",
			"query":
			{
				"term": "` + search_phrase + `"	
			},
			"from": 0,
			"max_results": 20,
			"_source": []
		}`

		req, err := http.NewRequest("POST", "http://localhost:4080/api/messages/_search", strings.NewReader(query))
		if err != nil {
			log.Fatal(err)
		}
		req.SetBasicAuth("admin", "Complexpass#123")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		log.Println(resp.StatusCode)
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		w.Write([]byte(body))
	})

	http.ListenAndServe(":8080", r)
}
