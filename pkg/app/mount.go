package app

import (
	"net/http"
)

// Mount mounst
func Mount(mux *http.ServeMux) {
	mux.HandleFunc("/", index) // List all news
	mux.Handle("/upload/", http.StripPrefix("/upload", http.FileServer(http.Dir("upload"))))
	mux.Handle("/news/", http.StripPrefix("/news", http.HandlerFunc(newsView)))
	// mux.Handle("/news/", http.StripPrefix("/news", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	id := r.URL.Path[1:]
	// 	newsView(id).ServeHTTP(w, r)
	// }))) // /news/:id
	mux.HandleFunc("/register", adminRegister)

	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/login", adminLogin)
	adminMux.HandleFunc("/list", adminList)
	adminMux.HandleFunc("/create", adminCreate)
	adminMux.HandleFunc("/edit", adminEdit)

	mux.Handle("/admin/", http.StripPrefix("/admin", onlyAdmin(adminMux)))
}

func onlyAdmin(h http.Handler) http.Handler {
	return h
}
