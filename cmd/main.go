package main

import (
	"fmt"
	"github.com/emilstorgaardandersen/stockDataApi/pkg/server"
	"os"
)

var Version = "v1"

func main()  {
	srv, err := server.New("Emil Andersen")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Starting stockAPI version: '%s' is created by '%s', listening on :8080\n", Version, srv.Name)
	err = srv.ListenAndServe(":8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}