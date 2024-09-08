package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
  "os"
)


func main() {
	m := http.NewServeMux()
  addr := ":"+os.Getenv("PORT")

	m.HandleFunc("/", handlePage)

	srv := http.Server{
		Handler:      m,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	// this blocks forever, until the server
	// has an unrecoverable error
	fmt.Println("server started on ", addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

func handlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	const page = `<html>
<head></head>
<body>
	<p> Hi Docker, I pushed a new version!</p>
</body>
</html>
`
	w.Write([]byte(page))
}
