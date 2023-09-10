package main

import (
	"context"
	"database/sql"
	"log"
	"reflect"
	"sqlc-crud-go/dbsqlc"
	"sqlc-crud-go/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func run() error {
	ctx := context.Background()

	connStr := "user=pqgotest password=postgres dbname=postgres sslmode=disable host=localhost"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return err
	}

	queries := dbsqlc.New(db)

	// list all authors
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)

	// create an author
	insertedAuthor, err := queries.CreateAuthor(ctx, dbsqlc.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio:  sql.NullString{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	})
	if err != nil {
		return err
	}
	log.Println(insertedAuthor)

	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthor.ID)
	if err != nil {
		return err
	}

	log.Println(reflect.DeepEqual(insertedAuthor, fetchedAuthor))
	return nil
}

func main() {
	db, err := sql.Open("postgres", "user=pqgotest dbname=pqgotest password=postgres sslmode=disable")
	if err != nil {
		log.Fatalf("Error opening database connection: %s", err)
		return
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %s", err)
		return
	}

	defer db.Close()

	router := gin.Default()

	queries := dbsqlc.New(db)
	routes.SetupRoutes(router, queries)

	log.Fatal(router.Run(":4747"))
}
