package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	// create 2 separate loggers for INFO and ERROR messages.
	// other useful flags:
	// log.Llongfile - print the full file path, intead of the name only
	// log.LUTC - print the time in the UTC instead of local
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("snippet/create", snippetCreate)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog, Handler: mux,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()

	errorLog.Fatal(err)
}
