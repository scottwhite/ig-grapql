package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
)

const defaultPort = "8080"

func main() {
	if err := SetupConfig(); err != nil {
		log.Println("nope on config", err)
	}

	conn, err := NewDB()
	if err != nil {
		log.Println("nope on db")
	}

	defer conn.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := NewResolver(conn)

	//queryHandler := handler.GraphQL(dataloader.MakeExecutableSchema(dataloader.New(conn)))

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))

	http.Handle("/query", handler.GraphQL(NewExecutableSchema(Config{Resolvers: r})))

	//http.Handle("/query", dataloader.DataloaderMiddleware(conn, queryHandler))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
