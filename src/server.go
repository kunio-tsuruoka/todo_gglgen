package main

import (
	"log"
	"net/http"
	"os"
	"todo/graph/generated"
	"todo/graph/resolver"
	// "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {

	godotenv.Load(".env")
	log.Printf(os.Getenv("DB_CONNECTION"))
	db, err := ConnectDatabase(os.Getenv("DB_CONNECTION"))
	if err != nil {
		panic(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func ConnectDatabase(con string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(con), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}