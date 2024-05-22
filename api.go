package main

import (
	"log"
	"net/http"
	"time"
)

type APIServer struct {
	Addr string
}

func newServer(addr string) *APIServer {
	return &APIServer{
		Addr: addr,
	}
}

func (s *APIServer) Run() {
	router := http.NewServeMux()
	router.HandleFunc("GET /", handlePage)

	server := http.Server{
		Addr: s.Addr,
		Handler: router,
		WriteTimeout: 30 * time.Second,
		ReadTimeout: 30 * time.Second,
	}

	log.Printf("Server has started %s", s.Addr)

	log.Fatal(server.ListenAndServe())
}


func handlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	const page = `<html>
<head></head>
<body>
	<p> Hello from Docker! I'm a Go server. </p>
</body>
</html>
`
	w.Write([]byte(page))
}