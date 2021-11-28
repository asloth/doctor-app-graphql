package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/asloth/doctor-app-graphql/database"
	"github.com/asloth/doctor-app-graphql/graph"
	"github.com/asloth/doctor-app-graphql/graph/generated"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const defaultPort = "8080"

func main() {
	//DATABASE
	//el string con los datos de conexion para conectarnos a la bd
	var dsn = "host=localhost user=postgres password=73481200 dbname=doctorapp-db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	//instaciamos la conexion con la base de datos
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	//seteamos el handler de la base de datos
	database.Handler = database.NewBaseHandler(db)
	//migramos los schemas
	database.Handler.Migrate()

	//SERVER
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
