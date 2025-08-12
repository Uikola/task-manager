package http

import "net/http"

func NewServer() http.Handler {
	router := http.NewServeMux()

	addRoutes(router)

	var handler http.Handler = router

	return handler
}

func addRoutes(_ *http.ServeMux) {
}
