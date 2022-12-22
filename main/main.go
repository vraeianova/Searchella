package main

func main() {
	// r := chi.NewRouter()

	// // Usa middleware para manejar el registro y las solicitudes con formato JSON
	// r.Use(middleware.Logger)
	// r.Use(middleware.Recoverer)

	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello World!"))
	// })
	// http.ListenAndServe(":9001", r)
	mux := Routes()
	server := NewServer(mux)
	server.Run()
}
