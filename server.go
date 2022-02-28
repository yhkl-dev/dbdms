package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/fatih/color"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/yhkl-dev/dbdms/graph"
	"github.com/yhkl-dev/dbdms/graph/generated"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
	// jwt struct {
	// 	secret string
	// }
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		next.ServeHTTP(w, r)
	})
}

func routes(db *sql.DB) http.Handler {
	router := httprouter.New()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))
	router.HandlerFunc(http.MethodGet, "/", playground.Handler("GraphQL playground", "/query"))
	router.Handler(http.MethodPost, "/query", srv)
	return enableCORS(router)
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development or production)")
	flag.StringVar(&cfg.db.dsn, "dsn", "postgres://postgres:123456@82.156.172.46:5432/dbdms?sslmode=disable", "Postgres connection string ")
	// flag.StringVar(&cfg.jwt.secret, "jwt-secret", "2dce505d96a53c5768052ee90f3df2055657518dad489160df9913f66042e160", "secret")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	serve := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      routes(db),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("connect to http://localhost:%d/ for GraphQL playground", cfg.port)
	color.Green("Start")
	log.Fatal(serve.ListenAndServe())
}
