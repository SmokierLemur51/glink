package main

import (
	"net/http"
	"flag"
	"log"
	"net"
	"fmt"

	"github.com/SmokierLemur51/glink/router"

)


func main() {
	var (
		host = flag.String("host", "", "host addr to listen on")
		port = flag.String("port", "5000", "port number for http listener")
	)
	flag.Parse()
	addr := net.JoinHostPort(*host, *port)
	if err := runHttp(addr); err != nil {
		log.Fatal(err)
	}
}


func runHttp(listenAddr string) error {
	s := http.Server{
		Addr: listenAddr,
		Handler: router.NewRouter(),
	}
	fmt.Println("Starting HTTP listener at ", listenAddr)
	return s.ListenAndServe()
}