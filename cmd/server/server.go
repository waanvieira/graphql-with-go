package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/waanvieira/graphqlgen/graph"
	"github.com/waanvieira/graphqlgen/internal/database"

	// Driver SQLite3. se não importarmos não conseguimos acessar o banco, indicamos o _ porque não vamos usar diretamente em uma variavel
	_ "github.com/mattn/go-sqlite3"
)

const defaultPort = "8080"

func main() {
	// Abrindo conexão com o banco de dados
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	// Encerramos a conexão com o banco de dados depois de utilizar, com o defer só vai encerrar quando encerrar a execução
	defer db.Close()

	// Instanciamos a struct category e injetamos a conexão com o banco de dados
	categoryDb := database.NewCategory(db)
	courseDb := database.NewCourse(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Aqui injetamos a nossa categoryDB em Resolver
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		// Aqui injetamos a nossa struct de category para que o nosso resolver possa utiliza-lo no schemma.resolvers
		CategoryDB: categoryDb,
		CourseDB:   courseDb,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
