package main

import (
	"fmt"
	"net/http"
)

type Server struct{}

func main() {
	server := NewServer()
	err := server.Run()
	if err != nil {
		panic(err)
	}
}

func InitRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/docs", docsHandler)
	mux.HandleFunc("/openapi.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/openapi.json")
	})

	return mux
}

func (s *Server) Run() error {
	fmt.Println("Web server started at :8080")
	return http.ListenAndServe(":8080", InitRouter())
}

func NewServer() Server {
	return Server{}
}

func docsHandler(w http.ResponseWriter, r *http.Request) {
	const bodyHTML = `<!DOCTYPE html>
<html>
<head><title>API Docs</title></head>
<body>
  <script id="api-reference" data-url="/openapi.json"></script>
  <script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
</body>
</html>
`
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	_, err := fmt.Fprint(w, bodyHTML)
	if err != nil {
		panic(err)
	}
}
