package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	const addr = ":8080"
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlePage)

	// block until server has unrecoverable error
	fmt.Println("server started on ", addr)
	srv := http.Server{
		Handler:      mux,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	err := srv.ListenAndServe()
	log.Fatal(err)

}

func handlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	w.WriteHeader(http.StatusOK)
	const page = `
      <html> 
         <body>
            <p> Hello from the docker container!!!!!!!!!! </p>
         </body>
      </html>
      `
	w.Write([]byte(page))
}
